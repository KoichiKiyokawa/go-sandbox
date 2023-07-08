package usecase

import (
	"bulletproof-go/gen/queries"
	"context"
)

type TransactionManager interface {
	Transaction(ctx context.Context, action func(queries *queries.Queries) error) error
}
