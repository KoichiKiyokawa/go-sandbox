package repository

import (
	"context"
	"fx-di/ent"
)

type PostRepository interface {
	FindAll(ctx context.Context) ([]*ent.Post, error)
	FindOne(ctx context.Context, id int) (*ent.Post, error)
}
