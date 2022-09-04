package dao

import (
	"fx-di/domain/model"
	"fx-di/domain/repository"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindOne(id int) (*model.User, error) {
	// TODO: implement
	return &model.User{}, nil
}
