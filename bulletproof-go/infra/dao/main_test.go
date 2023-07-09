package dao

import (
	"bulletproof-go/gen/queries"
	"database/sql"
	"os"
	"testing"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db          *sql.DB
	testQueries *queries.Queries
)

const dbFile = "test.db"

func TestMain(m *testing.M) {
	var err error
	_ = os.Remove(dbFile)
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	// create tables
	b, err := os.ReadFile("../../db/schema.sql")
	if err != nil {
		panic(err)
	}

	ddl := string(b)
	if _, err := db.Exec(ddl); err != nil {
		panic(err)
	}

	testQueries = queries.New(db)
	m.Run()
}

func WithTx(action func(q *queries.Queries)) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	action(testQueries.WithTx(tx))
	_ = tx.Rollback()
}
