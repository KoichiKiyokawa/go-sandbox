package dao

import (
	"context"
	"fx-di/app/domain/repository"
	"fx-di/ent"
	"fx-di/ent/post"
)

type postRepository struct {
	db *ent.Client
}

func NewPostRepository(db *ent.Client) repository.PostRepository {
	return &postRepository{db}
}

func (r *postRepository) FindOne(ctx context.Context, id int) (*ent.Post, error) {
	p, err := r.db.Post.Query().Where(post.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return convertPost(p), nil
}

func (r *postRepository) FindAll(ctx context.Context) ([]*ent.Post, error) {
	posts, err := r.db.Post.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*ent.Post, len(posts))
	for i, p := range posts {
		result[i] = convertPost(p)
	}
	return result, nil
}

func convertPost(p *ent.Post) *ent.Post {
	return &ent.Post{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	}
}
