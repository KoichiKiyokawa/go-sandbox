package usecase

import (
	"bulletproof-go/gen/queries"
	"bulletproof-go/graph/model"
	"context"

	"github.com/jinzhu/copier"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Find(ctx context.Context, id string) (*model.User, error)
	Register(ctx context.Context, input RegisterInput) error
}

type userUseCase struct {
	queries            queries.Querier
	transactionManager TransactionManager
}

func NewUserUseCase(queries queries.Querier, transactionManager TransactionManager) UserUseCase {
	return &userUseCase{queries: queries, transactionManager: transactionManager}
}

func (u *userUseCase) FindAll(ctx context.Context) ([]*model.User, error) {
	users, err := u.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*model.User, len(users))
	if err := copier.Copy(&res, &users); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) Find(ctx context.Context, id string) (*model.User, error) {
	user, err := u.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	var res *model.User
	if err := copier.Copy(res, user); err != nil {
		return nil, err
	}

	return res, nil
}

type RegisterInput struct {
	Name  string
	Email string
}

func (u *userUseCase) Register(ctx context.Context, input RegisterInput) error {
	return u.transactionManager.Transaction(ctx, func(queries *queries.Queries) error {
		return nil
	})
}
