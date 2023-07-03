package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"gqlgen-ent/graph"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Email is the resolver for the email field.
func (r *userResolver) Email(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *ent.User) ([]*ent.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }