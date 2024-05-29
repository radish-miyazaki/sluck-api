package repository

import (
	"context"
	"database/sql"
)

type MessageRepository interface {
	DeleteByUserID(ctx context.Context, userID string) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db}
}

func (m messageRepository) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := m.db.ExecContext(ctx, "DELETE FROM messages WHERE user_id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}
