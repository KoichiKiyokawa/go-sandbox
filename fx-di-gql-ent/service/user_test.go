package service

import (
	"context"
	"errors"
	"fx-di/domain/repository/mock"
	"fx-di/ent"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_user_FindOne(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		userID       int
		mockUserRepo *mock.UserRepositoryMock
		want         *ent.User
		wantErr      error
	}{
		{
			name:   "success",
			userID: 1,
			mockUserRepo: &mock.UserRepositoryMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.User, error) {
					return &ent.User{ID: 1, Name: "user1", Email: "user1@example.com"}, nil
				},
			},
			want:    &ent.User{ID: 1, Name: "user1", Email: "user1@example.com"},
			wantErr: nil,
		},
		{
			name:   "error",
			userID: 1,
			mockUserRepo: &mock.UserRepositoryMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.User, error) {
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
			us := &userService{userRepo: tt.mockUserRepo}
			got, err := us.FindOne(ctx, tt.userID)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_user_FindAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		mockUserRepo *mock.UserRepositoryMock
		want         []*ent.User
		wantErr      error
	}{
		{
			name: "success",
			mockUserRepo: &mock.UserRepositoryMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.User, error) {
					return []*ent.User{
						{ID: 1, Name: "user1", Email: "user1@example.com"},
						{ID: 2, Name: "user2", Email: "user2@example.com"},
					}, nil
				},
			},
			want: []*ent.User{
				{ID: 1, Name: "user1", Email: "user1@example.com"},
				{ID: 2, Name: "user2", Email: "user2@example.com"},
			},
			wantErr: nil,
		},
		{
			name: "error",
			mockUserRepo: &mock.UserRepositoryMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.User, error) {
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
			us := &userService{userRepo: tt.mockUserRepo}
			got, err := us.FindAll(ctx)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
