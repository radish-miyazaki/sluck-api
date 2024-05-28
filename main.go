package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/radish-miyazaki/sluck/controller"
	"github.com/radish-miyazaki/sluck/repository"
	"github.com/radish-miyazaki/sluck/usecase"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator.New()}

	ur := repository.UserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	e.POST("/users", uc.Create)

	e.Start(":8080")
}
