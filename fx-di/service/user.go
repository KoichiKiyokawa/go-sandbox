package service

import (
	"context"
	"fx-di/domain/model"
	"fx-di/domain/repository"
)

type UserService interface {
	FindOne(ctx context.Context, id int) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) FindOne(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) FindAll(ctx context.Context) ([]*model.User, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
