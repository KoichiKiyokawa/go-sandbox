//go:build wireinject
// +build wireinject

package di

import (
	"bulletproof-go/gen/queries"
	"bulletproof-go/infra/db"
	"bulletproof-go/resolver"
	"bulletproof-go/usecase"
	"database/sql"

	"github.com/google/wire"
)

func InitializeResolver(_db *sql.DB) *resolver.Resolver {
	wire.Build(
		// infra
		queries.New,
		wire.Bind(new(queries.DBTX), new(*sql.DB)),
		wire.Bind(new(queries.Querier), new(*queries.Queries)),
		db.NewTransactionManager,

		// usecase
		usecase.NewUserUseCase,

		resolver.NewResolver,
	)

	return &resolver.Resolver{}
}
