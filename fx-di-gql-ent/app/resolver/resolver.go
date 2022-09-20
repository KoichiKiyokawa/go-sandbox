//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"fx-di/app/service"
	"fx-di/ent"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	userService service.UserService
	postService service.PostService
	db          *ent.Client
}

func NewResolver(userService service.UserService, postService service.PostService, db *ent.Client) *Resolver {
	return &Resolver{userService: userService, postService: postService, db: db}
}
