package utils

import (
	"fmt"
	"net/mail"
	"slices"
	"unicode"

	"api/common/constants"
)

// IsAuthProviderValid Validates the authentication provider (e.g., Google, Facebook, ...)
// and return a boolean indicating success or failure.
func IsAuthProviderValid(provider string) bool {
	return slices.Contains(constants.AuthProviders, provider)
}

// IsMfaMethodValid Validates the multiple authentication method name (e.g., email, phone, authenticator)
// and return a boolean indicating success or failure.
func IsMfaMethodValid(mfaMethod string) bool {
	return slices.Contains(constants.AuthMfaMethods, mfaMethod)
}

// IsPhoneNumberValid Validates the phone number and return a boolean indicating success or failure.
func IsPhoneNumberValid(phoneNumber uint64) bool {
	return phoneNumber > 1000000
}

// IsEmailValid Validates the email address and return a boolean indicating success or failure.
func IsEmailValid(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}

// IsPasswordValid Validates the password and return a boolean indicating success or failure,
// along with a string listing all missing requirements.
func IsPasswordValid(password string) (bool, string) {
	var (
		hasMinLen  bool
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
		missing    string
	)

	if len(password) >= 8 {
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
		missing = fmt.Sprintf(
			"HAS_MIN_LENGTH: %t, HAS_UPPERCASE_LETTER: %t, HAS_LOWERCASE_LETTER: %t, HAS_NUMBER: %t, HAS_SPECIAL_CHARACTER: %t]",
			hasMinLen, hasUpper, hasLower, hasNumber, hasSpecial,
		)
	}

	return isValid, missing
}

// IsFacebookLoginScopesValid Validates the required scopes to allow login with Facebook
func IsFacebookLoginScopesValid(scopes []string) bool {
	counter := 0
	for _, scope := range scopes {
		if slices.Contains(constants.AuthLoginWithFacebookRequiredScopes, scope) {
			counter++
		}
	}
	return counter == len(constants.AuthLoginWithFacebookRequiredScopes)
}
