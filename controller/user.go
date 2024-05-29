package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/radish-miyazaki/sluck/usecase"
)

type UserController interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func (uc userController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	u, err := uc.u.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc userController) Update(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	u := toModel(req)
	if err := uc.u.Update(ctx.Request().Context(), u); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc userController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := uc.u.Delete(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (uc userController) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	u := toModel(req)
	id, err := uc.u.Create(ctx.Request().Context(), u)
	if err != nil {
		return err
	}

	u.ID = id
	return ctx.JSON(http.StatusCreated, u)
}
