//go:build wireinject
// +build wireinject

package di

import (
	"bulletproof-go/infra/dao"
	"bulletproof-go/infra/db"
	"bulletproof-go/resolver"
	"bulletproof-go/usecase"
	"database/sql"

	"github.com/google/wire"
)

func InitializeResolver(_db *sql.DB) *resolver.Resolver {
	wire.Build(
		// infra
		db.NewTransactionManager,
		db.NewDbManager,

		// repository
		dao.NewUserDAO,

		// usecase
		usecase.NewUserUseCase,

		resolver.NewResolver,
	)

	return &resolver.Resolver{}
}
