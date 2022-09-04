package repository

import (
	"context"
	"fx-di/domain/model"
)

type UserRepository interface {
	FindOne(ctx context.Context, id int) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}
