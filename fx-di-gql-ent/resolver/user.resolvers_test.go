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
	withDIContainer(t, func(m *mock_service.MockUserService) {
		ctx := context.Background()

		type params struct {
			id int
		}

		tests := []struct {
			name string
			mock func() *mock_service.MockUserService
			params
			want *ent.User
			err  error
		}{
			{
				name: "Normal time, can get user",
				mock: func() *mock_service.MockUserService {
					m.EXPECT().FindOne(gomock.Eq(ctx), gomock.Eq(1)).Return(&ent.User{
						ID:    1,
						Name:  "test user",
						Email: "test@example.com",
					}, nil)
					return m
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
				mock: func() *mock_service.MockUserService {
					m.EXPECT().FindOne(gomock.Eq(ctx), gomock.Eq(1)).Return(nil, errors.New("some error"))
					return m
				},
				params: params{id: 1},
				want:   nil,
				err:    errors.New("some error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				qr := &queryResolver{Resolver: &Resolver{userService: tt.mock()}}
				got, err := qr.User(ctx, tt.params.id)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.err, err)
			})
		}
	})
}

func TestUsers(t *testing.T) {
	withDIContainer(t, func(m *mock_service.MockUserService) {
		ctx := context.Background()

		tests := []struct {
			name string
			mock func() *mock_service.MockUserService
			want []*ent.User
			err  error
		}{
			{
				name: "Normal time, can get all users",
				mock: func() *mock_service.MockUserService {
					m.EXPECT().FindAll(gomock.Eq(ctx)).Return([]*ent.User{
						{ID: 1, Name: "test1", Email: "test1@example.com"},
						{ID: 2, Name: "test2", Email: "test2@example.com"},
					}, nil)
					return m
				},
				want: []*ent.User{
					{ID: 1, Name: "test1", Email: "test1@example.com"},
					{ID: 2, Name: "test2", Email: "test2@example.com"},
				},
				err: nil,
			},
			{
				name: "Error time, can't get user",
				mock: func() *mock_service.MockUserService {
					m.EXPECT().FindAll(gomock.Eq(ctx)).Return(nil, errors.New("some error"))
					return m
				},
				want: nil,
				err:  errors.New("some error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				qr := &queryResolver{Resolver: &Resolver{userService: tt.mock()}}
				got, err := qr.Users(ctx)
				assert.Equal(t, tt.want, got)
				assert.Equal(t, tt.err, err)
			})
		}
	})
}
