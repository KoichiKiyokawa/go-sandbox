package repository

import (
	"context"
	"fx-di/ent"
)

type PostRepository interface {
	FindAllByUserID(ctx context.Context, userID int) ([]*ent.Post, error)
	FindAll(ctx context.Context) ([]*ent.Post, error)
	FindOne(ctx context.Context, id int) (*ent.Post, error)
}
