package handler

import (
	"context"

	"huma-sandbox/internal/infra/storage"

	"braces.dev/errtrace"
	"github.com/danielgtaylor/huma/v2"
)

type userHandler struct {
	storage *storage.Storage
}

func NewUserHandler(s *storage.Storage) *userHandler {
	return &userHandler{storage: s}
}

type UserResponseBody struct {
	ID       string  `json:"id"                 doc:"user id"`
	Name     string  `json:"name"`
	Nickname *string `json:"nickname,omitempty"`
}

type UserListInput struct {
	PaginationInput
	Name string `json:"name,omitempty" query:"name" doc:"Filter users by name"`
}
type UserListResponse struct {
	Body struct {
		Users []UserResponseBody `json:"users"`
	}
}

func (h *userHandler) FindUserList(ctx context.Context, input *UserListInput) (*UserListResponse, error) {
	var resp UserListResponse

	filter := storage.FindUserListFilter{}
	if input.Name != "" {
		filter.Name = &input.Name
	}

	userList, err := h.storage.FindUserList(ctx, filter)
	if err != nil {
		return nil, huma.Error500InternalServerError("", errtrace.Wrap(err))
	}

	resp.Body.Users = make([]UserResponseBody, 0, len(userList))

	for _, user := range userList {
		resp.Body.Users = append(resp.Body.Users, UserResponseBody{
			ID:       user.GetID().String(),
			Name:     user.GetName(),
			Nickname: user.GetNickname(),
		})
	}

	return &resp, nil
}

type UserResponse struct {
	Body UserResponseBody
}

func (*userHandler) FindUser(_ context.Context, _ *struct {
	ID string `doc:"user id" path:"id"`
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

func (*userHandler) CreateUser(_ context.Context, _ *UserCreateInput) (*UserResponse, error) {
	var resp UserResponse

	return &resp, nil
}
