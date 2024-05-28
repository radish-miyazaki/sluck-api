package usecase

import (
	"context"
	"fmt"

	"github.com/radish-miyazaki/sluck/model"
	"github.com/radish-miyazaki/sluck/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, user *model.User) error
}

type userUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

func (u userUsecase) Create(ctx context.Context, user *model.User) error {
	fmt.Println("call usecase creating")
	return nil
}
