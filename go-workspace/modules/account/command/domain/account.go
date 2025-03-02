package domain

import (
	"time"

	"braces.dev/errtrace"
)

type UnAuthenticatedAccount struct {
	id    AccountID
	email Email
}

func NewUnAuthenticatedAccount(email string) (*UnAuthenticatedAccount, error) {
	e, err := NewEmail(email)
	if err != nil {
		return nil, errtrace.Wrap(err)
	}

	return &UnAuthenticatedAccount{
		id:    GenerateAccountID(),
		email: *e,
	}, nil
}

type AuthenticatedAccount struct {
	UnAuthenticatedAccount

	authenticatedAt time.Time
}

func (a UnAuthenticatedAccount) Authenticate() *AuthenticatedAccount {
	return &AuthenticatedAccount{
		UnAuthenticatedAccount: a,
		authenticatedAt:        time.Now(),
	}
}

// below, getters

func (a UnAuthenticatedAccount) ID() AccountID {
	return a.id
}

func (a UnAuthenticatedAccount) Email() Email {
	return a.email
}
