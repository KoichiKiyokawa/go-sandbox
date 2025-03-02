package domain

import (
	"fmt"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(value string) (*Email, error) {
	if !emailRegex.MatchString(value) {
		return nil, fmt.Errorf("invalid email: %s", value)
	}
	return &Email{value: value}, nil
}
