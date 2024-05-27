package repository

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) Create(ctx context.Context) error {
	return nil
}
