package cmd

import (
	"context"
	"gin/api"
	"gin/internal/repository"
	"gin/internal/usecase"
	"gin/pkg/database"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	ctx := context.Background()
	_, err := database.InitDB(ctx)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// REPOSITORY
	userRepository := repository.NewUserRepository()

	// USECASE
	userUseCase := usecase.NewUserService().SetUserRepository(userRepository)

	// Handler
	handler := api.CreateHandler(userUseCase)

	api.RegisterPath(r, handler)

	// Start server
	r.Run(":" + os.Getenv("HTTP_PORT"))
}
