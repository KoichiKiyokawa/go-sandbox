package model

import (
	"fmt"
	"typespec-oai-codegen/generated"
)

func NewInsufficientBalanceError(currentBalance int, changeAmount int) *generated.UserServiceInsufficientBalanceError {
	return &generated.UserServiceInsufficientBalanceError{
		Code:    generated.N4001,
		Message: fmt.Sprintf("Insufficient balance(current: %d, change: %d)", currentBalance, changeAmount),
	}
}
