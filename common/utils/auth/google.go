package auth

import (
	"context"
	"fmt"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/config"

	googleIdToken "google.golang.org/api/idtoken"
)

var contextGoogle = context.Background()

// Verifies a Google ID token and returns the associated user information.
//
// Refer to the official Google documentation for more details on token validation
// https://developers.google.com/identity/openid-connect/openid-connect#discovery
func VerifyGoogleIDToken(token string) (*types.GoogleUserProfileResponse, error) {
	// Validate the token
	if len(token) <= 0 {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	tokenValidator, err := googleIdToken.NewValidator(contextGoogle)
	if err != nil {
		return nil, constants.HTTP_500_ERROR_MESSAGE("validate Google token")
	}
	payload, err := tokenValidator.Validate(contextGoogle, token, config.Env.GooglePlusClientId)
	if err != nil {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	if payload.Expires <= time.Now().Unix() {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}

	// Retrieve info from claims
	user := &types.GoogleUserProfileResponse{}
	user.ID, _ = payload.Claims["sub"].(string)
	user.Email, _ = payload.Claims["email"].(string)
	user.EmailVerified, _ = payload.Claims["email_verified"].(bool)
	user.LastName, _ = payload.Claims["family_name"].(string)
	user.FirstName, _ = payload.Claims["given_name"].(string)
	user.FullName, _ = payload.Claims["name"].(string)
	user.Language, _ = payload.Claims["locale"].(string)
	user.Picture, _ = payload.Claims["picture"].(string)
	user.Expires = payload.Expires
	user.IssuedAt = payload.IssuedAt
	return user, nil
}
