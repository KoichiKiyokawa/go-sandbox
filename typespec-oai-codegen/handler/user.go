package handler

import (
	"context"

	"typespec-oai-codegen/generated"
	"typespec-oai-codegen/generated/db"
)

// UserServiceList implements generated.StrictServerInterface.
func (h *handler) UserServiceList(ctx context.Context, request generated.UserServiceListRequestObject) (generated.UserServiceListResponseObject, error) {
	users, err := h.queries.GetUserList(ctx, db.GetUserListParams{Limit: 100, Offset: 0})
	if err != nil {
		return nil, err
	}

	res := make(generated.UserServiceList200JSONResponse, len(users))
	for i, user := range users {
		res[i] = convertUser(user)
	}

	return res, nil
}

// UserServiceRead implements generated.StrictServerInterface.
func (h *handler) UserServiceRead(ctx context.Context, request generated.UserServiceReadRequestObject) (generated.UserServiceReadResponseObject, error) {
	user, err := h.queries.GetUserByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return generated.UserServiceRead200JSONResponse(convertUser(user)), nil
}

func convertUser(user db.User) generated.User {
	return generated.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
