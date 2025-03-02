package api

import (
	"context"

	"go-workspace-module-account/command/domain"

	"braces.dev/errtrace"
	"github.com/danielgtaylor/huma/v2"
	"github.com/samber/do"
)

type Handler struct {
	api               huma.API
	accountRepository domain.Repository
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		api:               do.MustInvoke[huma.API](i),
		accountRepository: do.MustInvoke[domain.Repository](i),
	}, nil
}

func (h *Handler) RegisterAll() {
	// ============================================

	type CreateAccountInput struct {
		Body struct {
			Email string
		}
	}

	type CreateAccountOutputBody struct {
		ID string
	}

	type CreateAccountOutput struct {
		Body CreateAccountOutputBody
	}

	huma.Register(h.api, huma.Operation{
		Method: "POST",
		Path:   "/accounts",
	}, func(ctx context.Context, input *CreateAccountInput) (*CreateAccountOutput, error) {
		a, err := domain.NewUnAuthenticatedAccount(input.Body.Email)
		if err != nil {
			return nil, huma.Error400BadRequest("invalid email", errtrace.Wrap(err))
		}

		if err := h.accountRepository.CreateUnAuthenticated(ctx, *a); err != nil {
			return nil, errtrace.Wrap(err)
		}

		return &CreateAccountOutput{
			Body: CreateAccountOutputBody{
				ID: a.ID().String(),
			},
		}, nil
	})
	// ============================================

	type UpdateAccountInput struct {
		ID    string
		Email string
	}

	type UpdateAccountOutput struct {
	}
}
