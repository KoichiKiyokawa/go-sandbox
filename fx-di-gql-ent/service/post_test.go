package service

import (
	"context"
	"errors"
	"fx-di/domain/repository/mock"
	"fx-di/ent"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_post_FindOne(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		postID       int
		mockPostRepo *mock.PostRepositoryMock
		want         *ent.Post
		wantErr      error
	}{
		{
			name:   "success",
			postID: 1,
			mockPostRepo: &mock.PostRepositoryMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.Post, error) {
					return &ent.Post{ID: 1, Title: "post1", Content: "post1-content"}, nil
				},
			},
			want:    &ent.Post{ID: 1, Title: "post1", Content: "post1-content"},
			wantErr: nil,
		},
		{
			name:   "error",
			postID: 1,
			mockPostRepo: &mock.PostRepositoryMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.Post, error) {
					return nil, errors.New("error")
				},
			},
			want:    nil,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ps := &postService{postRepo: tt.mockPostRepo}
			got, err := ps.FindOne(ctx, tt.postID)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestFindAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		mockPostRepo *mock.PostRepositoryMock
		want         []*ent.Post
		wantErr      error
	}{
		{
			name: "success",
			mockPostRepo: &mock.PostRepositoryMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.Post, error) {
					return []*ent.Post{
						{ID: 1, Title: "post1", Content: "post1-content"},
						{ID: 2, Title: "post2", Content: "post2-content"},
					}, nil
				},
			},
			want: []*ent.Post{
				{ID: 1, Title: "post1", Content: "post1-content"},
				{ID: 2, Title: "post2", Content: "post2-content"},
			},
			wantErr: nil,
		},
		{
			name: "error",
			mockPostRepo: &mock.PostRepositoryMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.Post, error) {
					return nil, errors.New("error")
				},
			},
			want:    nil,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ps := &postService{postRepo: tt.mockPostRepo}
			got, err := ps.FindAll(ctx)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
