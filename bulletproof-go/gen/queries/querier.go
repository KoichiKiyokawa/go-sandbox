// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package queries

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, id string) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)
