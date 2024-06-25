package utils

import (
	"net/mail"
	"strconv"
	"unicode"
)

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
		missing = "[hasMinLen: " + strconv.FormatBool(hasMinLen) + ", hasUpper: " + strconv.FormatBool(hasUpper) + ", hasLower: " + strconv.FormatBool(hasLower) + ", hasNumber: " + strconv.FormatBool(hasNumber) + ", hasSpecial: " + strconv.FormatBool(hasSpecial) + "]"
	}

	return isValid, missing
}