package repository

import (
	"context"
	"fx-di/ent"
)

type UserRepository interface {
	FindOne(ctx context.Context, id int) (*ent.User, error)
	FindAll(ctx context.Context) ([]*ent.User, error)
	FindOneByPostID(ctx context.Context, postID int) (*ent.User, error)
}
