package repository

import (
	"context"
	"mockery-sandbox/app/domain/model"
)

//go:generate go run github.com/vektra/mockery/v2 --name UserRepository
type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
}
