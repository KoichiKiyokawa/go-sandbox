package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fx-di/ent"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.userService.FindOne(ctx, id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.userService.FindAll(ctx)
}
