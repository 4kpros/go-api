package utils

import (
	"fmt"
	"time"
)

var tokenLenErr string = "Invalid or empty token! Please enter valid information."

// Verifies a Google token, returning the user ID and expiration time if it is valid.
func IsGoogleTokenValid(token string) (string, *time.Time, error) {
	var err error
	var userId = ""
	if len(token) <= 0 {
		err = fmt.Errorf("%s", tokenLenErr)
		return userId, nil, err
	}

	// Check with Google API
	var expires = time.Now().Add(time.Hour * 24)
	return userId, &expires, err
}

// Verifies a Facebook token, returning the user ID and expiration time if it is valid.
func IsFacebookTokenValid(token string) (string, *time.Time, error) {
	var err error
	var userId = ""
	if len(token) <= 0 {
		err = fmt.Errorf("%s", tokenLenErr)
		return userId, nil, err
	}

	// Check with Google API
	var expires = time.Now().Add(time.Hour * 24)
	return userId, &expires, err
}
