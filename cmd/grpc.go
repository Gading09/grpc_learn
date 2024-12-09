package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"strings"
	"syscall"

	domain "grpc/api"
	"grpc/internal/constant"
	repository "grpc/internal/repository"
	usecase "grpc/internal/usecase"
	middleware "grpc/middleware"

	pb "grpc/pb/v1"
	pkg "grpc/pkg"
	"grpc/pkg/database"
	"grpc/pkg/logger"
	"grpc/pkg/model"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.elastic.co/apm/module/apmgrpc/v2"
	"go.elastic.co/apm/module/apmhttp/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func Init() {
	ctx := context.Background()
	_, err := database.InitGorm(ctx)
	if err != nil {
		panic(err)
	}
	g, _ := errgroup.WithContext(ctx)
	var servers []*http.Server
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			log.Printf("received signal: %s\n", sig)
			for i, s := range servers {
				if err := s.Shutdown(ctx); err != nil {
					if err == nil {
						log.Printf("error shutting down server %d: %v", i, err)
						panic(err)
					}
				}
			}
			os.Exit(1)
		}
		return nil
	})

	g.Go(func() error { return NewGrpcServer() })
	g.Go(func() error { return NewHttpServer() })
	err = g.Wait()
	if err != nil {
		panic(err)
	}
	return
}

func NewGrpcServer() error {
	logger.Configure()
	port, err := net.Listen("tcp", ":"+pkg.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// REPOSITORY
	userRepository := repository.NewUserRepository()

	// USECASE
	userUseCase := usecase.NewUserService().SetUserRepository(userRepository)

	s := grpc.NewServer(grpc.MaxMsgSize(constant.MAX_SIZE_GRPC), grpc.MaxRecvMsgSize(constant.MAX_SIZE_GRPC), grpc.MaxSendMsgSize(constant.MAX_SIZE_GRPC),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		),
		grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery())),
		grpc.StreamInterceptor(apmgrpc.NewStreamServerInterceptor()),
	)

	handler := domain.CreateHandler(s, userUseCase)

	pb.RegisterUserServer(s, handler)

	logger.Infof("Serving gRPC on 0.0.0.0: %v", pkg.Getenv("GRPC_PORT"))
	s.Serve(port)

	return nil
}

func NewHttpServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	addr := "0.0.0.0:" + pkg.Getenv("GRPC_PORT")
	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(constant.MAX_SIZE_GRPC), grpc.MaxCallSendMsgSize(constant.MAX_SIZE_GRPC)),
		grpc.WithStreamInterceptor(apmgrpc.NewStreamClientInterceptor()),
		grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		panic(err)
	}
	defer conn.Close()

	// Create new grpc-gateway
	rmux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(middleware.CustomMatcher))

	for _, f := range []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		pb.RegisterUserHandler,
	} {
		if err = f(ctx, rmux, conn); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}

	// create http server mux
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   pkg.GetEnvCors("CORS_ORIGIN_ALLOWED"),
		AllowedMethods:   pkg.GetEnvCors("CORS_METHOD_ALLOWED"),
		AllowedHeaders:   pkg.GetEnvCors("CORS_HEADER_ALLOWED"),
		AllowCredentials: false,
	})

	// running rest http server
	logger.Infof("Serving Rest Http on 0.0.0.0: %v", pkg.Getenv("HTTP_PORT"))
	err = http.ListenAndServe("0.0.0.0:"+pkg.Getenv("HTTP_PORT"), corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		wHead := w.Header()
		rHead := request.Header
		wHead.Set(constant.ACAM, strings.Join(pkg.GetEnvCors("CORS_METHOD_ALLOWED"), ", "))
		wHead.Set(constant.ACAH, constant.ACAH_VALUE)
		wHead.Set(constant.ACAC, constant.ACAC_VALUE)
		wHead.Set(constant.HSTS, constant.HSTS_VALUE)
		wHead.Set(constant.CC, constant.CC_VALUE)
		wHead.Set(constant.XCTO, constant.XCTO_VALUE)
		wHead.Set("Content-Security-Policy", "default-src 'self'")

		rHead.Set(constant.GRPC_METHOD, strings.Join(pkg.GetEnvCors("CORS_METHOD_ALLOWED"), ", "))
		var bodyString string
		if request.Body != nil {
			buf := new(strings.Builder)
			_, err := io.Copy(buf, request.Body)
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}

			bodyString, err = middleware.ValidateRequestBody(buf.String())
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			// fmt.Println("request : ", bodyString)

			request.Body = io.NopCloser(strings.NewReader(bodyString))
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, request)
		apmhttp.Wrap(mux)

		// var result interface{}
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		// grab the captured response body
		data := rec.Body.Bytes()

		response := model.Response{}
		json.Unmarshal(data, &response)
		var bodyInterface interface{}
		json.Unmarshal([]byte(bodyString), &bodyInterface)
		bodyMarshal, _ := json.Marshal(bodyInterface)

		// REPOSITORY LOGGER
		logger.Info(fmt.Sprintf("[ms-sme-crow:log] [RequestURL] : %s, [RequestMethod] : %s, [RequestBody] : %s, [ResponseData] : %s", request.RequestURI, request.Method, string(bodyMarshal), string(data)))

		if request.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			code := response.ErrorCode
			if code < 100 {
				code = response.Code
			}

			if code < 100 {
				response = model.Response{
					Code:         http.StatusInternalServerError,
					ResponseCode: constant.FAILED_INTERNAL,
					ResponseDesc: http.StatusText(http.StatusInternalServerError),
					ResponseData: nil,
				}
				data, _ = json.Marshal(response)
				code = http.StatusInternalServerError
			}
			w.WriteHeader(code)
		}

		i, err := w.Write(data)
		if err != nil {
			logger.Info(fmt.Sprintf("Error Write Data %d : %v", i, err.Error()))
		}
	})))
	if err != nil {
		log.Fatal(err)
		panic(err)
		return err
	}
	return nil
}
