package main

import (
	"os"
	"testing"

	mockDB "without-di/mock/db"

	"github.com/stretchr/testify/assert"
)

func TestHoge(t *testing.T) {
	tests := []struct {
		name                               string
		selectedUserRepositoryFindByIdMock mockDB.UserRepositoryFindByIdEnum
		want                               string
	}{
		{
			name:                               "first case",
			selectedUserRepositoryFindByIdMock: mockDB.UserRepositoryFindByIdNormal,
			want:                               `{"id":1,"name":"user 1"}`,
		},
		{
			name:                               "second case",
			selectedUserRepositoryFindByIdMock: mockDB.UserRepositoryFindByIdEmpty,
			want:                               "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(mockDB.UserRepositoryFindByIdKey, string(tt.selectedUserRepositoryFindByIdMock))
			got := run()
			assert.Equal(t, tt.want, got)
		})
	}
}
