package handler

import (
	"context"
	"huma-sandbox/internal/storage"

	"braces.dev/errtrace"

	"github.com/danielgtaylor/huma/v2"
)

type userHandler struct {
	storage *storage.Storage
}

func NewUserHandler(storage *storage.Storage) *userHandler {
	return &userHandler{storage: storage}
}

type UserResponseBody struct {
	ID       string  `doc:"user id"             json:"id"`
	Name     string  `json:"name"`
	Nickname *string `json:"nickname,omitempty"`
}

type UserListInput struct {
	PaginationInput
	Name string `doc:"Filter users by name" query:"name"`
}
type UserListResponse struct {
	Body struct {
		Users []UserResponseBody `json:"users"`
	}
}

func (h *userHandler) FindUserList(ctx context.Context, input *UserListInput) (*UserListResponse, error) {
	var resp UserListResponse

	userList, err := h.storage.FindUserList(ctx)
	if err != nil {
		return nil, huma.Error500InternalServerError("", errtrace.Wrap(err))
	}

	resp.Body.Users = make([]UserResponseBody, 0, len(userList))
	for _, user := range userList {
		resp.Body.Users = append(resp.Body.Users, UserResponseBody{
			ID:       user.ID.Value().String(),
			Name:     user.Name.Value(),
			Nickname: user.Nickname.Value(),
		})
	}

	return &resp, nil
}

type UserResponse struct {
	Body UserResponseBody
}

func (h *userHandler) FindUser(ctx context.Context, input *struct {
	id string `doc:"user id" path:"id"`
},
) (*UserResponse, error) {
	var resp UserResponse

	return &resp, nil
}

type UserCreateInput struct {
	Body struct {
		Name     string `json:"name"`
		Nickname string `json:"nickname,omitempty"`
	}
}

func (h *userHandler) CreateUser(ctx context.Context, input *UserCreateInput) (*UserResponse, error) {
	var resp UserResponse

	return &resp, nil
}
