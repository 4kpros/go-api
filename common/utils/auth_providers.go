package utils

import (
	"fmt"
	"time"
)

var tokenLenErr string = "Invalid or empty token! Please enter valid information."

func IsGoogleTokenValid(token string) (userId string, expires time.Time, err error) {
	if len(token) <= 0 {
		err = fmt.Errorf("%s", tokenLenErr)
		return
	}

	// Check with Google API
	userId = "1234567890"
	expires = time.Now().Add(time.Hour * 24)
	return
}

func IsFacebookTokenValid(token string) (userId string, expires time.Time, err error) {
	if len(token) <= 0 {
		err = fmt.Errorf("%s", tokenLenErr)
		return
	}

	// Check with Facebook API
	userId = "1234567890"
	expires = time.Now().Add(time.Hour * 24)
	return
}
