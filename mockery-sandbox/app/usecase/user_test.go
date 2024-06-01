package usecase

import (
	"context"
	"mockery-sandbox/app/domain/model"
	"mockery-sandbox/app/domain/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_userUseCase_FetchAll(t *testing.T) {
	t.Run("hoge", func(t *testing.T) {
		ctx := context.Background()
		mockUserRepo := mocks.NewUserRepository(t)
		mockUserRepo.
			EXPECT().
			FindAll(ctx).
			Return([]*model.User{
				model.NewUser(1, "hoge"),
			}, nil)

		sut := NewUserUseCase(mockUserRepo)
		got, err := sut.FetchAll(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []*User{
			{ID: 1, Name: "hoge"},
		}, got)
	})
}
