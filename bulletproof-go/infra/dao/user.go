package dao

import (
	"bulletproof-go/domain/repository"
	"bulletproof-go/graph/model"
	"bulletproof-go/infra/db"
	"context"
)

type userDAO struct {
	db *db.DbManager
}

func NewUserDAO(db *db.DbManager) repository.UserRepository {
	return &userDAO{db: db}
}

// Find implements repository.UserRepository.
func (u *userDAO) Find(ctx context.Context, id string) (*model.User, error) {
	// TODO:
	u.db.GetDB(ctx).QueryContext(ctx, "SELECT id, name, email FROM users WHERE id = ?", id)
	return nil, nil
}

// FindAll implements repository.UserRepository.
func (*userDAO) FindAll(ctx context.Context) ([]*model.User, error) {
	panic("unimplemented")
}
