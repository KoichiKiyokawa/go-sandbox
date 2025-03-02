package domain

import "github.com/google/uuid"

type AccountID struct {
	value string
}

func GenerateAccountID() AccountID {
	return AccountID{value: uuid.New().String()}
}

func (a AccountID) String() string {
	return a.value
}
