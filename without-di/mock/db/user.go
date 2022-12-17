package db

import (
	"context"
	"log"
	"os"
	"without-di/model"

	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{}
}

type UserRepositoryFindByIdEnum string

const (
	UserRepositoryFindByIdKey = "UserRepositoryFindByIdKey"

	UserRepositoryFindByIdNormal UserRepositoryFindByIdEnum = "通常時"
	UserRepositoryFindByIdEmpty                             = "該当ユーザーがいなかったとき"
)

func (u *userRepository) FindById(ctx context.Context, id int) (*model.User, error) {
	selected := UserRepositoryFindByIdEnum(os.Getenv(UserRepositoryFindByIdKey))
	log.Printf("selected: %#+v\n", selected)
	switch selected {
	case UserRepositoryFindByIdNormal:
		return &model.User{ID: 1, Name: "user 1"}, nil
	case UserRepositoryFindByIdEmpty:
		return nil, nil
	}
	return nil, nil
}
