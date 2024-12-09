package api

import "github.com/gin-gonic/gin"

func RegisterPath(r *gin.Engine, h *Handler) {
	if h == nil {
		panic("item controller cannot be nil")
	}

	r.POST("v1/create-user", h.CreateUser)
	r.POST("v1/get-list-user", h.GetListUser)
	r.POST("v1/update-user", h.UpdateUser)
	r.POST("v1/delete-user", h.DeleteUser)
}
