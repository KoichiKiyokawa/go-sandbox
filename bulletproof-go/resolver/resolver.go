package resolver

import "bulletproof-go/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	userUseCase usecase.UserUseCase
}

func NewResolver(userUseCase usecase.UserUseCase) *Resolver {
	return &Resolver{userUseCase: userUseCase}
}
