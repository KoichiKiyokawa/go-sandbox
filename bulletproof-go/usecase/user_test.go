package usecase

import (
	"bulletproof-go/gen/queries"
	"bulletproof-go/graph/model"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_userUseCase_FindAll(t *testing.T) {
	tests := []struct {
		name    string
		q       *QuerierMock
		want    []*model.User
		wantErr error
	}{
		{
			name: "Should return all users",
			q: &QuerierMock{
				GetUsersFunc: func(ctx context.Context) ([]queries.User, error) {
					return []queries.User{
						{ID: "1", Name: sql.NullString{String: "test1", Valid: true}, Email: sql.NullString{String: "test1@example.com", Valid: true}},
						{ID: "2", Name: sql.NullString{String: "test2", Valid: true}, Email: sql.NullString{String: "test2@example.com", Valid: true}},
					}, nil
				},
			},
			want: []*model.User{
				{ID: "1", Name: "test1", Email: "test1@example.com"},
				{ID: "2", Name: "test2", Email: "test2@example.com"},
			},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserUseCase(tt.q, nil).FindAll(ctx)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}
