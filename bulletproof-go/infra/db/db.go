package db

import (
	"bulletproof-go/usecase"
	"context"
	"database/sql"
)

const dbKey = iota

type transactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) usecase.TransactionManager {
	return &transactionManager{db: db}
}

func (t *transactionManager) Transaction(ctx context.Context, action func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, dbKey, tx)
	if err := action(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}

	return tx.Commit()
}

type DbManager struct {
	db *sql.DB
}

func NewDbManager(db *sql.DB) *DbManager {
	return &DbManager{db: db}
}

func (d DbManager) GetDB(ctx context.Context) *sql.DB {
	if db := ctx.Value(dbKey).(*sql.DB); db != nil {
		return db
	}

	return d.db
}
