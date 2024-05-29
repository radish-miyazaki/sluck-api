package transaction

import (
	"context"
	"database/sql"

	"github.com/radish-miyazaki/sluck/repository"
)

type Transaction interface {
	DoInTx(context.Context, func(context.Context) (any, error)) (any, error)
}

type transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &transaction{db}
}

func (t transaction) DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, repository.TxKey, tx)
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}
