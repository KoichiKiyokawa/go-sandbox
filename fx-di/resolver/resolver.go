//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"fx-di/service"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	userService service.UserService
	postService service.PostService
}

func NewResolver(userService service.UserService, postService service.PostService) *Resolver {
	return &Resolver{userService: userService, postService: postService}
}
