package api

import (
	"context"
	"grpc/internal/constant"
	pb "grpc/pb/v1"
	"net/http"
)

func (s Handler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.UserService.CreateUser(ctx, request)
	if err != nil {
		return &pb.CreateUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return user, nil
}

func (s Handler) GetListUser(ctx context.Context, request *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	user, err := s.UserService.GetListUser(ctx, request)
	if err != nil {
		return &pb.GetListUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return user, nil
}

func (s Handler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, err := s.UserService.UpdateUser(ctx, request)
	if err != nil {
		return &pb.UpdateUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return user, nil
}

func (s Handler) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	user, err := s.UserService.DeleteUser(ctx, request)
	if err != nil {
		return &pb.DeleteUserResponse{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return user, nil
}
