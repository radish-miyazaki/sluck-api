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
	ur repository.UserRepository
	mr repository.MessageRepository
}

func NewUserUsecase(ur repository.UserRepository, mr repository.MessageRepository) UserUsecase {
	return &userUsecase{ur, mr}
}

func (u userUsecase) Delete(ctx context.Context, id string) error {
	if err := u.ur.Delete(ctx, id); err != nil {
		return err
	}

	if err := u.mr.DeleteByUserID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u userUsecase) Update(ctx context.Context, user *model.User) error {
	return u.ur.Update(ctx, user)
}

func (u userUsecase) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := u.ur.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := u.ur.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}
