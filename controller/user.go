package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/radish-miyazaki/sluck/usecase"
)

type UserController interface {
	Create(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (u userController) Create(ctx echo.Context) error {
	fmt.Println("creating ...")
	return nil
}
