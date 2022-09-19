package resolver

import (
	"context"
	"errors"
	mock_service "fx-di/app/service/mock"
	"fx-di/ent"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type params struct {
		id int
	}

	tests := []struct {
		name string
		params
		userServiceMock *mock_service.UserServiceMock
		want            *ent.User
		err             error
	}{
		{
			name: "Normal time, can get user",
			userServiceMock: &mock_service.UserServiceMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.User, error) {
					return &ent.User{
						ID:    1,
						Name:  "test user",
						Email: "test@example.com",
					}, nil
				},
			},
			params: params{id: 1},
			want: &ent.User{
				ID:    1,
				Name:  "test user",
				Email: "test@example.com",
			},
			err: nil,
		},
		{
			name: "Error time, can't get user",
			userServiceMock: &mock_service.UserServiceMock{
				FindOneFunc: func(ctx context.Context, id int) (*ent.User, error) {
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
			qr := &queryResolver{Resolver: &Resolver{userService: tt.userServiceMock}}
			got, err := qr.User(ctx, tt.params.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestUsers(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name            string
		preExec         func()
		userServiceMock *mock_service.UserServiceMock
		want            []*ent.User
		err             error
	}{
		{
			name: "Normal time, can get all users",
			userServiceMock: &mock_service.UserServiceMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.User, error) {
					return []*ent.User{
						{ID: 1, Name: "test1", Email: "test1@example.com"},
						{ID: 2, Name: "test2", Email: "test2@example.com"},
					}, nil
				},
			},
			want: []*ent.User{
				{ID: 1, Name: "test1", Email: "test1@example.com"},
				{ID: 2, Name: "test2", Email: "test2@example.com"},
			},
			err: nil,
		},
		{
			name: "Error time, can't get user",
			userServiceMock: &mock_service.UserServiceMock{
				FindAllFunc: func(ctx context.Context) ([]*ent.User, error) {
					return nil, errors.New("some error")
				},
			},
			want: nil,
			err:  errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := &queryResolver{Resolver: &Resolver{userService: tt.userServiceMock}}
			got, err := qr.Users(ctx)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
