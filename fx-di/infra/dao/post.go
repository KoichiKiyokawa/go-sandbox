package dao

import (
	"context"
	"fx-di/domain/model"
	"fx-di/domain/repository"
	"fx-di/ent"
	"fx-di/ent/post"
	"fx-di/ent/user"
)

type postRepository struct {
	db *ent.Client
}

func NewPostRepository(db *ent.Client) repository.PostRepository {
	return &postRepository{db}
}

func (r *postRepository) FindOne(ctx context.Context, id int) (*model.Post, error) {
	p, err := r.db.Post.Query().Where(post.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return convertPost(p), nil
}

func (r *postRepository) FindAll(ctx context.Context) ([]*model.Post, error) {
	posts, err := r.db.Post.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Post, len(posts))
	for i, p := range posts {
		result[i] = convertPost(p)
	}
	return result, nil
}

func (r *postRepository) FindAllByUserID(ctx context.Context, userID int) ([]*model.Post, error) {
	posts, err := r.db.User.Query().Where(user.IDEQ(userID)).QueryPosts().All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Post, len(posts))
	for i, p := range posts {
		result[i] = convertPost(p)
	}
	return result, nil
}

func convertPost(p *ent.Post) *model.Post {
	return &model.Post{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	}
}
