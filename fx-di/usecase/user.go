package usecase

import (
	"fx-di/domain/model"
	"fx-di/domain/repository"
)

type UserUseCase interface {
	FindOne(id int) (*model.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo}
}

func (s *userUseCase) FindOne(id int) (*model.User, error) {
	user, err := s.userRepo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
