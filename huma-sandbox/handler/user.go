package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type UserResponseBody struct {
	ID       string `json:"id" doc:"user id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname,omitempty"`
}

func RegisterUserHandlers(api huma.API) {
	type UserListInput struct {
		PaginationInput
		Name string `query:"name" doc:"Filter users by name"`
	}
	type UserListResponse struct {
		Body struct {
			Users []UserResponseBody `json:"users"`
		}
	}
	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/users",
		Description: "List all users",
		Responses: map[string]*huma.Response{
			// "200": {Description: "Success"},
			"500": {Description: "Internal server error"},
		},
		Tags: []string{"users"},
	}, func(ctx context.Context, input *UserListInput) (*UserListResponse, error) {
		var resp UserListResponse

		list, err := FindUserList(ctx)
		if err != nil {
			return nil, huma.Error500InternalServerError("", err)
		}

		if input.Name == "" {
			resp.Body.Users = list
		} else {
			for _, u := range list {
				if strings.Contains(strings.ToLower(u.Name), strings.ToLower(input.Name)) {
					resp.Body.Users = append(resp.Body.Users, u)
				}
			}
		}

		return &resp, nil
	})

	type UserResponse struct {
		Body UserResponseBody
	}
	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/users/{id}",
		Description: "Get a user by ID",
		Tags:        []string{"users"},
	}, func(ctx context.Context, i *struct{}) (*UserResponse, error) {
		var resp UserResponse
		return &resp, nil
	})

	type UserCreateInput struct {
		Body struct {
			Name     string `json:"name"`
			Nickname string `json:"nickname,omitempty"`
		}
	}
	huma.Register(api, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/users",
		Description: "Create a new user",
		Tags:        []string{"users"},
	}, func(ctx context.Context, i *UserCreateInput) (*UserResponse, error) {
		var resp UserResponse
		return &resp, nil
	})
}

func FindUserList(ctx context.Context) ([]UserResponseBody, error) {
	return []UserResponseBody{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
	}, nil
}
