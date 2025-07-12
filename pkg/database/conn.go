package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
)

func RollbackTx(ctx context.Context, tx pgx.Tx) {
	if tx != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			if !errors.Is(rollbackErr, pgx.ErrTxClosed) {
				log.Panicf("rollback tx error: %v", rollbackErr)
			}
		}
	}
}
