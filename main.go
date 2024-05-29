package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/radish-miyazaki/sluck/infra"
	"github.com/radish-miyazaki/sluck/transaction"

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator.New()}

	db, err := infra.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)
	mr := repository.NewMessageRepository(db)
	tx := transaction.NewTransaction(db)
	uu := usecase.NewUserUsecase(ur, mr, tx)
	uc := controller.NewUserController(uu)
	e.GET("/users/:id", uc.Get)
	e.POST("/users", uc.Create)
	e.PUT("/users", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	e.Start(":8080")
}
