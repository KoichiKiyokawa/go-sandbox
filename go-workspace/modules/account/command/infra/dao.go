package infra

import (
	"context"
	"database/sql"

	"go-workspace-module-account/command/domain"

	"github.com/samber/do"
)

type dao struct {
	db *sql.DB
}

func NewDao(i *do.Injector) (domain.Repository, error) {
	return &dao{
		db: do.MustInvoke[*sql.DB](i),
	}, nil
}

func (d dao) FindUnauthenticated(ctx context.Context, id string) (*domain.UnAuthenticatedAccount, error) {
	return nil, nil
}

func (d dao) CreateUnAuthenticated(ctx context.Context, account domain.UnAuthenticatedAccount) error {
	return nil
}

func (d dao) SaveAuthenticated(ctx context.Context, account domain.AuthenticatedAccount) error {
	panic("unimplemented")
}
