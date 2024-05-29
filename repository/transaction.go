package repository

import (
	"context"
	"database/sql"
)

var TxKey = struct{}{}

func GetTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(&TxKey).(*sql.Tx)
	return tx, ok
}
