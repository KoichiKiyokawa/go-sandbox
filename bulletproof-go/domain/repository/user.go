package repository

import (
	"bulletproof-go/graph/model"
	"context"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Find(ctx context.Context, id string) (*model.User, error)
}
