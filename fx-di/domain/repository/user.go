package repository

import "fx-di/domain/model"

type UserRepository interface {
	FindOne(id int) (*model.User, error)
}
