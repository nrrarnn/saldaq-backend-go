package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nrrarnn/saldaq-backend-go/internal/config"
	"github.com/nrrarnn/saldaq-backend-go/internal/user"
)

func main() {
	db := config.InitDB()
	e := echo.New()

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	user.NewUserHandler(e, userService)

	db.AutoMigrate(&user.User{})

	e.Logger.Fatal(e.Start(":8080"))
}
