package resolver

import (
	"context"
	"errors"
	"fx-di/ent"
	mock_service "fx-di/service/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type params struct {
		id int
	}

	tests := []struct {
		name string
		params
		postServiceMock *mock_service.PostServiceMock
		want            *ent.Post
		err             error
	}{
		{
			name: "Normal time, can get user",
			postServiceMock: &mock_service.PostServiceMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.Post, error) {
					return &ent.Post{
						ID:      1,
						Title:   "title",
						Content: "content",
					}, nil
				},
			},
			params: params{id: 1},
			want: &ent.Post{
				ID:      1,
				Title:   "title",
				Content: "content",
			},
			err: nil,
		},
		{
			name: "Error time, can't get user",
			postServiceMock: &mock_service.PostServiceMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.Post, error) {
					return nil, errors.New("some error")
				},
			},
			params: params{id: 1},
			want:   nil,
			err:    errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := &queryResolver{Resolver: &Resolver{postService: tt.postServiceMock}}
			got, err := qr.Post(ctx, tt.params.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestPosts(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name            string
		preExec         func()
		postServiceMock *mock_service.PostServiceMock
		want            []*ent.Post
		err             error
	}{
		{
			name: "Normal time, can get all users",
			postServiceMock: &mock_service.PostServiceMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.Post, error) {
					return []*ent.Post{
						{ID: 1, Title: "title1", Content: "content1"},
						{ID: 2, Title: "title2", Content: "content2"},
					}, nil
				},
			},
			want: []*ent.Post{
				{ID: 1, Title: "title1", Content: "content1"},
				{ID: 2, Title: "title2", Content: "content2"},
			},
			err: nil,
		},
		{
			name: "Error time, can't get user",
			postServiceMock: &mock_service.PostServiceMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.Post, error) {
					return nil, errors.New("some error")
				},
			},
			want: nil,
			err:  errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := &queryResolver{Resolver: &Resolver{postService: tt.postServiceMock}}
			got, err := qr.Posts(ctx)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
