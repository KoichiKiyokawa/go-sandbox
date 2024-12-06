package handler

import (
	"context"

	"huma-sandbox/internal/infra/storage"
)

type postHandler struct {
	storage *storage.Storage
}

func NewPostHandler(s *storage.Storage) *postHandler {
	return &postHandler{storage: s}
}

// --------------------------
// input and output structs
// --------------------------

type PostResponseBody struct {
	ID    string `json:"id" doc:"post id"`
	Title string `json:"title" doc:"post title"`
	Body  string `json:"body" doc:"post body"`
}

type PostListResponse struct {
	Body []PostResponseBody
}

// --------------------------
// handler funcs
// --------------------------

func (h *postHandler) FindPostList(ctx context.Context, input *struct{}) (*PostListResponse, error) {
	return nil, nil
}

func (h *postHandler) FindPostListByUserID(ctx context.Context, input *struct {
	UserID string `path:"id" doc:"user id"`
},
) (*PostListResponse, error) {
	return nil, nil
}
