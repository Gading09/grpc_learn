package api

import (
	"github.com/labstack/echo"
)

func RegisterPath(e *echo.Echo, h *Handler) {
	if h == nil {
		panic("item controller cannot be nil")
	}

	e.POST("v1/create-user", h.CreateUser)
	e.POST("v1/get-list-user", h.GetListUser)
	e.POST("v1/update-user", h.UpdateUser)
	e.POST("v1/delete-user", h.DeleteUser)
}
