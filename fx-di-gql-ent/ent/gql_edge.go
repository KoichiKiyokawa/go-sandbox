// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (po *Post) Author(ctx context.Context) (*User, error) {
	result, err := po.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryAuthor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Posts(ctx context.Context) ([]*Post, error) {
	result, err := u.NamedPosts(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = u.QueryPosts().All(ctx)
	}
	return result, err
}
