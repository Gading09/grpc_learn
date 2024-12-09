package api

import (

	// "grpc/internal/constant"
	// pb "grpc/pb/v1"

	"net/http"

	"gin/pkg/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	body := new(model.CreateUserRequest)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserService.CreateUser(ctx, body)
	if err != nil {
		c.JSON(resp.Code, err)
		return
	}
	c.JSON(resp.Code, resp)
	return
}

func (h *Handler) GetListUser(c *gin.Context) {
	ctx := c.Request.Context()
	body := new(model.GetRequest)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserService.GetListUser(ctx, body)
	if err != nil {
		c.JSON(resp.Code, err)
		return
	}
	c.JSON(resp.Code, resp)
	return
}

func (h *Handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	body := new(model.UpdateUserRequest)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserService.UpdateUser(ctx, body)
	if err != nil {
		c.JSON(resp.Code, err)
		return
	}
	c.JSON(resp.Code, resp)
	return
}

func (h *Handler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	body := new(model.DeleteUserRequest)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := h.UserService.DeleteUser(ctx, body)
	if err != nil {
		c.JSON(resp.Code, err)
		return
	}
	c.JSON(resp.Code, resp)
	return
}
