package handler

import (
	"context"
	"net/http"
	"testing"

	"grpc/api"
	"grpc/internal/constant"
	"grpc/internal/mocks"
	"grpc/internal/usecase"
	"grpc/pb/v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_CreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()
	email := "hedy@gmail.com"
	password := "123"
	name := "hedy"

	request := &pb.CreateUserRequest{
		Body: &pb.CreateUserRequest_DataUser{
			Email:    email,
			Password: password,
			Name:     name,
		},
	}

	expectedResponse := &pb.CreateUserResponse{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.CreateUserResponse_ResponseData{
			Id: mock.Anything,
		},
	}

	mockRepo.On("InsertUser", ctx, mock.AnythingOfType("model.User")).Return(nil).Once()

	service := usecase.NewUserService().SetUserRepository(mockRepo)
	handler := &api.Handler{
		UserService: service,
	}

	resp, err := handler.CreateUser(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Code, resp.Code)
	mockRepo.AssertExpectations(t)
}
