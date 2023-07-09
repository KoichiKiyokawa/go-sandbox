package dao

import (
	"bulletproof-go/gen/queries"
	"context"
	"database/sql"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDAO_GetUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	WithTx(func(q *queries.Queries) {
		created, err := q.CreateUser(ctx, queries.CreateUserParams{
			ID:    uuid.New().String(),
			Name:  sql.NullString{String: "test1", Valid: true},
			Email: sql.NullString{String: "test1@example.com", Valid: true},
		})
		if err != nil {
			t.Error(err)
		}

		spew.Dump(created)

		got, err := q.GetUser(ctx, created.ID)

		assert.NoError(t, err)
		assert.Equal(t, queries.User{
			ID:    created.ID,
			Name:  sql.NullString{String: "test1", Valid: true},
			Email: sql.NullString{String: "test1@example.com", Valid: true},
		},
			got)
	})
}
