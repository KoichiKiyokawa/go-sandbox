package usecase

import (
	"context"
	"mockery-sandbox/app/domain/model"
	"mockery-sandbox/app/domain/repository"
	"mockery-sandbox/app/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUseCase_FetchAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		makeUserRepo func() repository.UserRepository
		args         args
		want         []*User
		wantErr      error
	}{
		{
			name: "hoge",
			makeUserRepo: func() repository.UserRepository {
				m := mocks.NewUserRepository(t)
				m.On("FindAll", mock.Anything, mock.Anything).Return([]*model.User{}, nil)
				return m
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserUseCase(tt.makeUserRepo())
			got, err := u.FetchAll(tt.args.ctx)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
