package usecase

import (
	"bulletproof-go/domain/repository"
	"bulletproof-go/graph/model"
	"context"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Find(ctx context.Context, id string) (*model.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) FindAll(ctx context.Context) ([]*model.User, error) {
	return u.userRepo.FindAll(ctx)
}

func (u *userUseCase) Find(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.Find(ctx, id)
}
