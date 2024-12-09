package cmd

import (
	"context"
	"echo/api"
	"echo/internal/repository"
	"echo/internal/usecase"
	"echo/pkg/database"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	ctx := context.Background()
	_, err := database.InitDB(ctx)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// REPOSITORY
	userRepository := repository.NewUserRepository()

	// USECASE
	userUseCase := usecase.NewUserService().SetUserRepository(userRepository)

	// Handler
	handler := api.CreateHandler(userUseCase)

	api.RegisterPath(e, handler)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("HTTP_PORT")))
}
