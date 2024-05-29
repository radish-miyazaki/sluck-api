package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/radish-miyazaki/sluck/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	Read(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) Create(ctx context.Context, user *model.User) (string, error) {
	result, err := u.db.ExecContext(ctx, "INSERT INTO users (id, name, email, age) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Age)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

func (u userRepository) Update(ctx context.Context, user *model.User) error {
	result, err := u.db.ExecContext(ctx, "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?", user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err

	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", user.ID)
	}

	return nil
}

func (u userRepository) Delete(ctx context.Context, id string) error {
	result, err := u.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", id)
	}

	return nil
}

func (u userRepository) Read(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRowContext(ctx, "SELECT id, name, email, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
