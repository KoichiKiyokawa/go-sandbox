package service

import (
	"context"
	"fx-di/domain/repository"
	"fx-di/ent"
)

type PostService interface {
	FindOne(ctx context.Context, id int) (*ent.Post, error)
	FindAll(ctx context.Context) ([]*ent.Post, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}

// FindAll implements PostService
func (s *postService) FindAll(ctx context.Context) ([]*ent.Post, error) {
	posts, err := s.postRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// FindOne implements PostService
func (s *postService) FindOne(ctx context.Context, id int) (*ent.Post, error) {
	post, err := s.postRepo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
