package dao

import (
	"context"
	"fmt"
	"fx-di/ent"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOne(t *testing.T) {
	withTransaction(t, func(ctx context.Context, db *ent.Client) {
		// set dummy data
		for id := 1; id <= 3; id++ {
			db.User.Create().SetID(id).SetName(fmt.Sprintf("user%d", id)).SetEmail(fmt.Sprintf("user%d@example.com", id)).Save(ctx)
		}

		repo := NewUserRepository(db)
		got, err := repo.FindOne(ctx, 1)
		want := &ent.User{
			ID:    1,
			Name:  "user1",
			Email: "user1@example.com",
		}
		assert.Equal(t, want, got)
		assert.Equal(t, nil, err)
	})
}

func TestFindAll(t *testing.T) {
	withTransaction(t, func(ctx context.Context, db *ent.Client) {
		// set dummy data
		for id := 1; id <= 3; id++ {
			db.User.Create().SetID(id).SetName(fmt.Sprintf("user%d", id)).SetEmail(fmt.Sprintf("user%d@example.com", id)).Save(ctx)
		}

		repo := NewUserRepository(db)
		got, err := repo.FindAll(ctx)
		want := []*ent.User{
			{ID: 1, Name: "user1", Email: "user1@example.com"},
			{ID: 2, Name: "user2", Email: "user2@example.com"},
			{ID: 3, Name: "user3", Email: "user3@example.com"},
		}
		assert.Equal(t, want, got)
		assert.Equal(t, nil, err)
	})
}
