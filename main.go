package main

import (
	"fmt"

	"github.com/radish-miyazaki/sluck/controller"
	"github.com/radish-miyazaki/sluck/repository"
	"github.com/radish-miyazaki/sluck/usecase"
)

func main() {
	fmt.Println("Hello, World!")

	ur := repository.UserRepository(nil)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	uc.Create(nil)
}
