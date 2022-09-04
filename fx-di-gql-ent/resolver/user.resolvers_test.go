package resolver

import (
	"context"
	"errors"
	"fx-di/ent"
	mock_service "fx-di/service/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	withDIContainer(t, func(mockUserService *mock_service.MockUserService) {
		ctx := context.Background()

		type params struct {
			id int
		}

		tests := []struct {
			name    string
			preExec func()
			params
			want *ent.User
			err  error
		}{
			{
				name: "Normal time, can get user",
				preExec: func() {
					mockUserService.EXPECT().FindOne(gomock.Eq(ctx), gomock.Eq(1)).Return(&ent.User{
						ID:    1,
						Name:  "test user",
						Email: "test@example.com",
					}, nil)
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
				preExec: func() {
					mockUserService.EXPECT().FindOne(gomock.Eq(ctx), gomock.Eq(1)).Return(nil, errors.New("some error"))
				},
				params: params{id: 1},
				want:   nil,
				err:    errors.New("some error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.preExec()
				qr := &queryResolver{Resolver: &Resolver{userService: mockUserService}}
				got, err := qr.User(ctx, tt.params.id)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.err, err)
			})
		}
	})
}

func TestUsers(t *testing.T) {
	withDIContainer(t, func(mockUserService *mock_service.MockUserService) {
		ctx := context.Background()

		tests := []struct {
			name    string
			preExec func()
			want    []*ent.User
			err     error
		}{
			{
				name: "Normal time, can get all users",
				preExec: func() {
					mockUserService.EXPECT().FindAll(gomock.Eq(ctx)).Return(
						[]*ent.User{
							{ID: 1, Name: "test1", Email: "test1@example.com"},
							{ID: 2, Name: "test2", Email: "test2@example.com"},
						},
						nil,
					)
				},
				want: []*ent.User{
					{ID: 1, Name: "test1", Email: "test1@example.com"},
					{ID: 2, Name: "test2", Email: "test2@example.com"},
				},
				err: nil,
			},
			{
				name: "Error time, can't get user",
				preExec: func() {
					mockUserService.EXPECT().FindAll(gomock.Eq(ctx)).Return(nil, errors.New("some error"))
				},
				want: nil,
				err:  errors.New("some error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.preExec()
				qr := &queryResolver{Resolver: &Resolver{userService: mockUserService}}
				got, err := qr.Users(ctx)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.err, err)
			})
		}
	})
}
