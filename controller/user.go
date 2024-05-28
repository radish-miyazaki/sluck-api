package controller

import (
	"net/http"

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

func (uc userController) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	u := toModel(req)
	uc.u.Create(ctx.Request().Context(), u)
	return nil
}
