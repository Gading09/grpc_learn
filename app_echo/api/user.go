package api

import (

	// "grpc/internal/constant"
	// pb "grpc/pb/v1"

	"net/http"

	"github.com/labstack/echo"

	"echo/pkg/model"
)

func (h *Handler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	body := new(model.CreateUserRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := h.UserService.CreateUser(ctx, body)
	if err != nil {
		return c.JSON(resp.Code, err)
	}
	return c.JSON(resp.Code, resp)
}

func (h *Handler) GetListUser(c echo.Context) error {
	ctx := c.Request().Context()
	body := new(model.GetRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := h.UserService.GetListUser(ctx, body)
	if err != nil {
		return c.JSON(resp.Code, err)
	}
	return c.JSON(resp.Code, resp)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	body := new(model.UpdateUserRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := h.UserService.UpdateUser(ctx, body)
	if err != nil {
		return c.JSON(resp.Code, err)
	}
	return c.JSON(resp.Code, resp)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	body := new(model.DeleteUserRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := h.UserService.DeleteUser(ctx, body)
	if err != nil {
		return c.JSON(resp.Code, err)
	}
	return c.JSON(resp.Code, resp)
}
