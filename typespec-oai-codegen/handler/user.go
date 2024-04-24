package handler

import (
	"context"

	"typespec-oai-codegen/generated"
	"typespec-oai-codegen/generated/db"

	"github.com/cockroachdb/errors"
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

// UserServiceCreate implements generated.StrictServerInterface.
func (h *handler) UserServiceCreate(ctx context.Context, request generated.UserServiceCreateRequestObject) (generated.UserServiceCreateResponseObject, error) {
	created, err := h.queries.CreateUser(ctx, db.CreateUserParams{
		Name:  request.Body.Name,
		Email: request.Body.Email,
	})
	if err != nil {
		return nil, err
	}

	return generated.UserServiceCreate200JSONResponse(convertUser(created)), nil
}

// UserServiceSendBalance implements generated.StrictServerInterface.
func (h *handler) UserServiceSendBalance(ctx context.Context, request generated.UserServiceSendBalanceRequestObject) (generated.UserServiceSendBalanceResponseObject, error) {
	var afterFromUserBalance int64
	var afterToUserBalance int64
	if err := h.transactioner.WithTx(func(qtx *db.Queries) error {
		fromUser, err := qtx.ChangeBalance(ctx, db.ChangeBalanceParams{ID: request.Params.FromUserId, Amount: -int64(request.Body.Amount)})
		if err != nil {
			return errors.WithStack(err)
		}

		if fromUser.Balance < 0 {
			return errors.New("insufficient balance")
		}

		toUser, err := qtx.ChangeBalance(ctx, db.ChangeBalanceParams{ID: request.Params.ToUserId, Amount: int64(request.Body.Amount)})
		if err != nil {
			return errors.WithStack(err)
		}

		afterFromUserBalance = fromUser.Balance
		afterToUserBalance = toUser.Balance

		return nil
	}); err != nil {
		return nil, err
	}

	return generated.UserServiceSendBalance200JSONResponse{
		FromUserBalance: int32(afterFromUserBalance),
		ToUserBalance:   int32(afterToUserBalance),
	}, nil
}
