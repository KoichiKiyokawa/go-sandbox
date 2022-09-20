package repository

import (
	"context"
	"fx-di/ent"
)

//go:generate go run github.com/matryer/moq -rm -out ./mock/user.go -pkg mock . UserRepository
type UserRepository interface {
	FindOne(ctx context.Context, id int) (*ent.User, error)
	FindAll(ctx context.Context) ([]*ent.User, error)
}
