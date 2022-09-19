package dao

import (
	"context"
	"fx-di/ent"
	"fx-di/ent/enttest"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// withTransaction runs the given function in a transaction and rolls it back at the end.
func withTransaction(t *testing.T, fn func(ctx context.Context, db *ent.Client)) {
	ctx := context.Background()
	client := enttest.Open(t, "postgres", os.Getenv("DB_URL"))
	defer client.Close()

	tx, err := client.Tx(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fn(ctx, tx.Client())

	_ = tx.Rollback()
}
