package api

import (
	"grpc/internal/usecase"

	"google.golang.org/grpc"
)

type Handler struct {
	Grpc        *grpc.Server
	UserService usecase.UserService
}

func CreateHandler(
	grpc *grpc.Server,
	user usecase.UserService,
) *Handler {
	return &Handler{
		grpc,
		user,
	}
}
