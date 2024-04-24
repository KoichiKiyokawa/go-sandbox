package db

import (
	"database/sql"
	"typespec-oai-codegen/generated/db"

	"github.com/cockroachdb/errors"
)

type Transactioner interface {
	WithTx(action func(qtx *db.Queries) error) error
}

type transactionerImpl struct {
	db      *sql.DB
	queries db.Queries
}

func NewTransactionerImpl(db *sql.DB) Transactioner {
	return &transactionerImpl{db: db}
}

func (t *transactionerImpl) WithTx(action func(qtx *db.Queries) error) error {
	tx, err := t.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	qtx := t.queries.WithTx(tx)

	if err := action(qtx); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.WithStack(err)
		}

		return err
	}

	return errors.WithStack(tx.Commit())
}
