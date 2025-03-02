package domain

import "context"

type Repository interface {
	FindUnauthenticated(ctx context.Context, id string) (*UnAuthenticatedAccount, error)

	CreateUnAuthenticated(ctx context.Context, account UnAuthenticatedAccount) error
	SaveAuthenticated(ctx context.Context, account AuthenticatedAccount) error
}
