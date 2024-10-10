package utils

import (
	"fmt"
	"net/mail"
	"slices"
	"unicode"

	"github.com/4kpros/go-api/common/constants"
)

func IsAuthProviderValid(provider string) bool {
	return slices.Contains(constants.AUTH_PROVIDERS, provider)
}

func IsPhoneNumberValid(phoneNumber int) bool {
	return phoneNumber > 100000
}

func IsEmailValid(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}

func IsPasswordValid(password string) (bool, string) {
	var (
		hasMinLen  bool
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
		missing    string
	)

	if len(password) >= 7 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	isValid := hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial

	if !isValid {
		missing = fmt.Sprintf("[hasMinLen: %t, hasUpper: %t, hasLower: %t, hasNumber: %t, hasSpecial: %t]", hasMinLen, hasUpper, hasLower, hasMinLen, hasSpecial)
	}

	return isValid, missing
}
