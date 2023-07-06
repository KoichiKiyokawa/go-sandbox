package usecase

import (
	"bulletproof-go/domain/repository"
	"bulletproof-go/graph/model"
	"context"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Find(ctx context.Context, id string) (*model.User, error)
	Register(ctx context.Context, input RegisterInput) error
}

type userUseCase struct {
	userRepo           repository.UserRepository
	transactionManager TransactionManager
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

type RegisterInput struct {
	Name  string
	Email string
}

func (u *userUseCase) Register(ctx context.Context, input RegisterInput) error {
	return u.transactionManager.Transaction(ctx, func(ctx context.Context) error {
		return nil
	})
}
