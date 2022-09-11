package repository

import (
	"context"
	"fx-di/ent"
)

//go:generate go run github.com/matryer/moq -rm -out ./mock/post.go -pkg mock . PostRepository
type PostRepository interface {
	FindAll(ctx context.Context) ([]*ent.Post, error)
	FindOne(ctx context.Context, id int) (*ent.Post, error)
}
