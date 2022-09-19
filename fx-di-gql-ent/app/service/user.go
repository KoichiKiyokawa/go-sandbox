package service

import (
	"context"
	"fx-di/app/domain/repository"
	"fx-di/ent"
)

//go:generate go run github.com/matryer/moq -rm -out ./mock/user.go -pkg mock . UserService
type UserService interface {
	FindOne(ctx context.Context, id int) (*ent.User, error)
	FindAll(ctx context.Context) ([]*ent.User, error)
	FindOneByPostID(ctx context.Context, postID int) (*ent.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) FindOne(ctx context.Context, id int) (*ent.User, error) {
	user, err := s.userRepo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) FindAll(ctx context.Context) ([]*ent.User, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindOneByPostID(ctx context.Context, postID int) (*ent.User, error) {
	user, err := s.userRepo.FindOneByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
