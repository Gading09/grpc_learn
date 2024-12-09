package usecase

import (
	"context"
	"grpc/internal/constant"
	"grpc/internal/repository"
	pb "grpc/pb/v1"
	"grpc/pkg/model"
	"grpc/pkg/util"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetListUser(context.Context, *pb.GetListUserRequest) (*pb.GetListUserResponse, error)
	UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) SetUserRepository(userRepository repository.UserRepository) *userService {
	s.UserRepository = userRepository
	return s

}

func (s *userService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	now := time.Now()
	id := uuid.New().String()
	body := request.GetBody()
	password, err := util.HashPassword(body.GetPassword())
	if err != nil {
		return &pb.CreateUserResponse{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		}, nil
	}

	payload := model.User{
		Id:        id,
		Name:      body.GetName(),
		Email:     body.GetEmail(),
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.UserRepository.InsertUser(ctx, payload)
	if err != nil {
		return &pb.CreateUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return &pb.CreateUserResponse{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.CreateUserResponse_ResponseData{
			Id: id,
		},
	}, nil
}

func (s *userService) GetListUser(ctx context.Context, request *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	pagination := util.GeneratePaginationFromRequest(request.Body.Query.GetLimit(), request.Body.Query.GetPage(), request.Body.Query.GetField(), request.Body.Query.GetSort())
	users, count, err := s.UserRepository.GetListUser(ctx, pagination, request.Body.Where.GetFilter())
	if err != nil {
		return &pb.GetListUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	if count == 0 {
		return &pb.GetListUserResponse{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
			ResponseData: nil,
		}, nil
	}

	var data []*pb.GetListUserResponse_ResponseData_ListUser
	for _, item := range users {
		user := &pb.GetListUserResponse_ResponseData_ListUser{
			Id:    item.Id,
			Email: item.Email,
			Name:  item.Name,
		}
		data = append(data, user)
	}

	return &pb.GetListUserResponse{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.GetListUserResponse_ResponseData{
			Page:      request.Body.Query.GetPage(),
			Limit:     request.Body.Query.GetLimit(),
			Total:     int32(count),
			TotalPage: int32(math.Ceil(float64(count) / float64(request.Body.Query.GetLimit()))),
			ListUser:  data,
		},
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	body := request.GetBody()
	user, err := s.UserRepository.GetUserById(ctx, body.GetId())
	if err != nil {
		return &pb.UpdateUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	err = util.CheckPassword(body.GetPassword(), user.Password)
	if err != nil {
		return &pb.UpdateUserResponse{
			Code:         http.StatusForbidden,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: http.StatusText(http.StatusForbidden),
			ResponseData: nil,
		}, nil
	}
	newPassword := user.Password
	now := time.Now()
	if body.GetNewPassword() != "" {
		newPassword, err = util.HashPassword(body.GetNewPassword())
		if err != nil {
			return &pb.UpdateUserResponse{
				Code:         http.StatusBadRequest,
				ResponseCode: constant.FAILED_REQUIRED,
				ResponseDesc: err.Error(),
				ResponseData: nil,
			}, nil
		}
	}

	err = s.UserRepository.UpdateUser(ctx, model.User{
		Id:        body.GetId(),
		Email:     body.GetEmail(),
		Password:  newPassword,
		Name:      body.GetName(),
		UpdatedAt: now,
	})
	if err != nil {
		return &pb.UpdateUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return &pb.UpdateUserResponse{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.UpdateUserResponse_ResponseData{
			Id: body.GetId(),
		},
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (result *pb.DeleteUserResponse, err error) {
	body := request.GetBody()
	user, err := s.UserRepository.GetUserById(ctx, body.GetId())
	if err != nil {
		return &pb.DeleteUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	now := time.Now()
	user.DeletedAt = &now
	err = s.UserRepository.UpdateUser(ctx, user)
	if err != nil {
		return &pb.DeleteUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}
	return &pb.DeleteUserResponse{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.DeleteUserResponse_ResponseData{
			Id: body.GetId(),
		},
	}, nil
}
