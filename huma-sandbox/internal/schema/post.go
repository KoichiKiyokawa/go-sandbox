package schema

import (
	"database/sql"
	"net/http"

	"huma-sandbox/internal/handler"
	"huma-sandbox/internal/infra/storage"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterPostHandlers(api huma.API, db *sql.DB) {
	postHandler := handler.NewPostHandler(storage.NewStorage(db))

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/posts",
		Description: "List all posts (optionally with filtering)",
		Responses: map[string]*huma.Response{
			"500": {Description: "Internal server error"},
		},
		Tags: []string{"posts"},
	}, postHandler.FindPostList)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/users/{id}/posts",
		Description: "List all posts of a user",
		Responses: map[string]*huma.Response{
			"500": {Description: "Internal server error"},
		},
		Tags: []string{"posts"},
	}, postHandler.FindPostListByUserID)
}
