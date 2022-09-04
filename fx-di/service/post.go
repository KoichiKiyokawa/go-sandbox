package service

import (
	"context"
	"fx-di/domain/model"
	"fx-di/domain/repository"
)

type PostService interface {
	FindAllByUserID(ctx context.Context, userID int) ([]*model.Post, error)
	FindOne(ctx context.Context, id int) (*model.Post, error)
	FindAll(ctx context.Context) ([]*model.Post, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}

func (s *postService) FindAllByUserID(ctx context.Context, userID int) ([]*model.Post, error) {
	posts, err := s.postRepo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// FindAll implements PostService
func (s *postService) FindAll(ctx context.Context) ([]*model.Post, error) {
	posts, err := s.postRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// FindOne implements PostService
func (s *postService) FindOne(ctx context.Context, id int) (*model.Post, error) {
	post, err := s.postRepo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
