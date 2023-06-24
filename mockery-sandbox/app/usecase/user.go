package usecase

import (
	"context"
	"mockery-sandbox/app/domain/repository"
)

type UserUseCase interface {
	FetchAll(ctx context.Context) ([]*User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

type User struct {
	ID   int
	Name string
}

func (u *userUseCase) FetchAll(ctx context.Context) ([]*User, error) {
	raw, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*User, 0, len(raw))
	for i := range raw {
		res = append(res, &User{
			ID:   raw[i].ID(),
			Name: raw[i].Name(),
		})
	}

	return res, nil
}
