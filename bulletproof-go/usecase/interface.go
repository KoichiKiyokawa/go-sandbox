package usecase

import "context"

type TransactionManager interface {
	Transaction(ctx context.Context, action func(ctx context.Context) error) error
}
