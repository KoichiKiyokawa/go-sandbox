package schema

import (
	"database/sql"
	"net/http"

	"huma-sandbox/internal/handler"
	"huma-sandbox/internal/handler/middleware"
	"huma-sandbox/internal/infra/storage"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterUserHandlers(api huma.API, db *sql.DB) {
	userHandler := handler.NewUserHandler(storage.NewStorage(db))

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/users",
		Description: "List all users",
		Responses: map[string]*huma.Response{
			// "200": {Description: "Success"},
			"500": {Description: "Internal server error"},
		},
		Tags: []string{"users"},
	}, userHandler.FindUserList)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/users/{id}",
		Description: "Get a user by ID",
		Tags:        []string{"users"},
	}, userHandler.FindUser)

	huma.Register(api, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/users",
		Description: "Create a new user",
		Tags:        []string{"users"},
	}, userHandler.CreateUser)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/me",
		Description: "Get the current user",
		Tags:        []string{"users"},
		Middlewares: huma.Middlewares{middleware.AuthMiddleware},
	}, userHandler.FindCurrentUser)
}
