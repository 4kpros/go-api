package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/config"
	googleIdToken "google.golang.org/api/idtoken"
)

const invalidTokenErrMessage = "Invalid or expired token! Please enter valid information."

var contextGoogle = context.Background()

// Verifies a Google ID token and returns the associated user information.
//
// Refer to the official Google documentation for more details on token validation
// https://developers.google.com/identity/openid-connect/openid-connect#discovery
func VerifyGoogleIDToken(token string) (*types.GoogleUserProfileResponse, error) {
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
	// Validate token expires time
	if payload.Expires <= time.Now().Unix() {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}

	// Extract info
	user := &types.GoogleUserProfileResponse{}
	user.Id, _ = payload.Claims["sub"].(string)
	user.Email, _ = payload.Claims["email"].(string)
	user.EmailVerified, _ = payload.Claims["email_verified"].(bool)
	user.LastName, _ = payload.Claims["family_name"].(string)
	user.FirstName, _ = payload.Claims["given_name"].(string)
	user.FullName, _ = payload.Claims["name"].(string)
	user.Language, _ = payload.Claims["locale"].(string)
	user.Picture, _ = payload.Claims["picture"].(string)
	return user, nil
}

// Verifies a Facebook token and returns the associated user information.
//
// Refer to the official documentation for more details
// Debug: https://graph.facebook.com/debug_token?input_token=YOUR_AUTH_TOKEN&access_token=YOUR_APP_ID|YOUR_CLIENT_SECRET
// Profile: https://graph.facebook.com/me?fields=id,name,last_name,first_name,email,languages,picture.width(100).height(100).as(picture_small),picture.width(720).height(720).as(picture_large)&access_token=YOUR_AUT_TOKEN
func VerifyFacebookToken(token string) (*types.FacebookUserProfileResponse, error) {
	if len(token) <= 0 {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}

	// Retrieve token info
	debugResp := &types.FacebookDebugAccessTokenResponse{}
	errDebug := HTTPGet(
		fmt.Sprintf(
			"%s%s&access_token=%s|%s",
			config.Env.FacebookDebugTokenUrl,
			token,
			config.Env.FacebookAppId,
			config.Env.FacebookClientSecret,
		),
		debugResp,
	)
	if errDebug != nil {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	// Validate claims
	if !debugResp.Data.IsValid || debugResp.Data.AppId != config.Env.FacebookAppId || debugResp.Data.Application != config.Env.FacebookAppName {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	// Retrieve user info
	userResp := &types.FacebookUserProfileResponse{}
	secretProof, errSecretProof := EncodeHMAC_SHA256(token, config.Env.FacebookClientSecret)
	if errSecretProof != nil {
		return nil, constants.HTTP_500_ERROR_MESSAGE("encode Facebook HMAC HS256 secret proof")
	}
	errUser := HTTPGet(
		fmt.Sprintf(
			"%s%s&appsecret_proof=%s",
			config.Env.FacebookProfileUrl,
			token,
			secretProof,
		),
		userResp,
	)
	if errUser != nil {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	return userResp, nil
}
