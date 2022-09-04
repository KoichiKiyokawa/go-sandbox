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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_service.NewMockUserService(ctrl)
	ctx := context.Background()

	type params struct {
		id int
	}

	type want struct {
		user *ent.User
		err  error
	}

	tests := []struct {
		name string
		mock func() *mock_service.MockUserService
		params
		want want
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
			want: want{user: &ent.User{
				ID:    1,
				Name:  "test user",
				Email: "test@example.com",
			}, err: nil},
		},
		{
			name: "Error time, can't get user",
			mock: func() *mock_service.MockUserService {
				m.EXPECT().FindOne(gomock.Eq(ctx), gomock.Eq(1)).Return(nil, errors.New("some error"))
				return m
			},
			params: params{id: 1},
			want:   want{user: nil, err: errors.New("some error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := &queryResolver{Resolver: &Resolver{userService: tt.mock()}}
			got, err := qr.User(ctx, tt.params.id)
			assert.Equal(t, tt.want.user, got)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
