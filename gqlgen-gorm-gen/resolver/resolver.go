package resolver

import "gqlgen-prisma/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	db *db.PrismaClient
}

func NewResolver(db *db.PrismaClient) *Resolver {
	return &Resolver{db: db}
}
