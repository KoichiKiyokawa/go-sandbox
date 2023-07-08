package db

import (
	"bulletproof-go/gen/queries"
	"bulletproof-go/usecase"
	"context"
	"database/sql"
)

const sqlcKey = iota

type transactionManager struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewTransactionManager(queries *queries.Queries) usecase.TransactionManager {
	return &transactionManager{queries: queries}
}

func (t *transactionManager) Transaction(ctx context.Context, action func(queries *queries.Queries) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := action(t.queries.WithTx(tx)); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}

	return tx.Commit()
}
