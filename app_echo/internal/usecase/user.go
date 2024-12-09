package usecase

import (
	"context"
	"echo/internal/constant"
	"echo/internal/repository"
	"echo/pkg/model"
	"echo/pkg/util"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, request *model.CreateUserRequest) (model.Response, error)
	GetListUser(ctx context.Context, request *model.GetRequest) (model.Response, error)
	UpdateUser(ctx context.Context, request *model.UpdateUserRequest) (model.Response, error)
	DeleteUser(ctx context.Context, request *model.DeleteUserRequest) (model.Response, error)
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

func (s *userService) CreateUser(ctx context.Context, request *model.CreateUserRequest) (model.Response, error) {
	now := time.Now()
	id := uuid.New().String()
	password, err := util.HashPassword(request.Password)
	if err != nil {
		return model.Response{
			Code:         http.StatusBadRequest,
			ResponseCode: constant.FAILED_REQUIRED,
			ResponseDesc: err.Error(),
			ResponseData: nil,
		}, nil
	}

	payload := model.User{
		Id:        id,
		Name:      request.Name,
		Email:     request.Email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.UserRepository.InsertUser(ctx, payload)
	if err != nil {
		return model.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return model.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: model.CreateUserResponse{
			Id: id,
		},
	}, nil
}

func (s *userService) GetListUser(ctx context.Context, request *model.GetRequest) (model.Response, error) {
	pagination := util.GeneratePaginationFromRequest(int32(request.Query.Limit), int32(request.Query.Page), request.Query.Field, request.Query.Sort)
	users, count, err := s.UserRepository.GetListUser(ctx, pagination, request.Where.Filter)
	if err != nil {
		return model.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	if count == 0 {
		return model.Response{
			Code:         http.StatusNotFound,
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
			ResponseData: nil,
		}, nil
	}

	var data []model.ListUser
	for _, item := range users {
		user := model.ListUser{
			Id:    item.Id,
			Email: item.Email,
			Name:  item.Name,
		}
		data = append(data, user)
	}

	return model.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: model.GetListUserResponse{
			Page:      request.Query.Page,
			Limit:     request.Query.Limit,
			Total:     int32(count),
			TotalPage: int32(math.Ceil(float64(count) / float64(request.Query.Limit))),
			ListUser:  data,
		},
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, request *model.UpdateUserRequest) (model.Response, error) {
	user, err := s.UserRepository.GetUserById(ctx, request.Id)
	if err != nil {
		return model.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	err = util.CheckPassword(request.Password, user.Password)
	if err != nil {
		return model.Response{
			Code:         http.StatusForbidden,
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: http.StatusText(http.StatusForbidden),
			ResponseData: nil,
		}, nil
	}
	newPassword := user.Password
	now := time.Now()
	if request.NewPassword != "" {
		newPassword, err = util.HashPassword(request.NewPassword)
		if err != nil {
			return model.Response{
				Code:         http.StatusBadRequest,
				ResponseCode: constant.FAILED_REQUIRED,
				ResponseDesc: err.Error(),
				ResponseData: nil,
			}, nil
		}
	}

	err = s.UserRepository.UpdateUser(ctx, model.User{
		Id:        request.Id,
		Email:     request.Email,
		Password:  newPassword,
		Name:      request.Name,
		UpdatedAt: now,
	})
	if err != nil {
		return model.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}

	return model.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: model.UpdateUserResponse{
			Id: request.Id,
		},
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, request *model.DeleteUserRequest) (model.Response, error) {
	user, err := s.UserRepository.GetUserById(ctx, request.Id)
	if err != nil {
		return model.Response{
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
		return model.Response{
			Code:         http.StatusInternalServerError,
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
			ResponseData: nil,
		}, nil
	}
	return model.Response{
		Code:         http.StatusOK,
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: model.DeleteUserResponse{
			Id: request.Id,
		},
	}, nil
}
