package repository

import (
	"context"
	"fx-di/domain/model"
)

type PostRepository interface {
	FindAllByUserID(ctx context.Context, userID int) ([]*model.Post, error)
	FindAll(ctx context.Context) ([]*model.Post, error)
	FindOne(ctx context.Context, id int) (*model.Post, error)
}
