package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/radish-miyazaki/sluck/controller"
	"github.com/radish-miyazaki/sluck/repository"
	"github.com/radish-miyazaki/sluck/usecase"
)

func main() {
	e := echo.New()

	fmt.Println("Hello, World!")

	ur := repository.UserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)

	e.Start(":8080")
}
