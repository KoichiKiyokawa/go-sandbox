package main

import "fmt"

func main() {
	anyEmail := findEmailByUserID("123")
	switch anyEmail.(type) {
	case UnverifiedEmail:
		fmt.Println("UnverifiedEmail")
	case VerifiedEmail:
		fmt.Println("VerifiedEmail")
	}

	saveAnyEmail(UnverifiedEmail{})
	saveAnyEmail(VerifiedEmail{})
}

type (
	AnyEmail        interface{ _AnyEmail() }
	UnverifiedEmail struct {
		AnyEmail
		value string
	}
	VerifiedEmail struct {
		AnyEmail
		value string
	}
	InvalidEmail struct {
		AnyEmail
		value string
	}
)

func (e UnverifiedEmail) Verify() (VerifiedEmail, error) {
	return VerifiedEmail{value: e.value}, nil
}

func findEmailByUserID(userID string) AnyEmail {
	if userID == "123" {
		return UnverifiedEmail{value: ""}
	} else {
		return VerifiedEmail{value: ""}
	}
}

func saveAnyEmail(email AnyEmail) {
	switch email.(type) {
	case UnverifiedEmail:
		fmt.Println("UnverifiedEmail")
	case VerifiedEmail:
		fmt.Println("VerifiedEmail")
	}
}
