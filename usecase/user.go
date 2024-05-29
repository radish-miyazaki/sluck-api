package usecase

import (
	"context"

	"github.com/radish-miyazaki/sluck/model"
	"github.com/radish-miyazaki/sluck/repository"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user *model.User) error
}

type userUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

func (u userUsecase) Delete(ctx context.Context, id string) error {
	return u.r.Delete(ctx, id)
}

func (u userUsecase) Update(ctx context.Context, user *model.User) error {
	return u.r.Update(ctx, user)
}

func (u userUsecase) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := u.r.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := u.r.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}
