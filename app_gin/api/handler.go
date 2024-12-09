package api

import (
	"gin/internal/usecase"
)

type Handler struct {
	UserService usecase.UserService
}

func CreateHandler(
	user usecase.UserService,
) *Handler {
	return &Handler{
		user,
	}
}
