package repository

import (
	"context"
	"without-di/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) FindById(ctx context.Context, id int) (*model.User, error) {
	var user *model.User
	res := u.db.First(user, id)
	return user, res.Error
}
