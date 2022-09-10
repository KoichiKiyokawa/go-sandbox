package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fx-di/ent"
)

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id int) (*ent.Post, error) {
	return r.postService.FindOne(ctx, id)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*ent.Post, error) {
	return r.postService.FindAll(ctx)
}
