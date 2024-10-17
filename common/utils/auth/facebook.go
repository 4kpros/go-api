package auth

import (
	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/common/utils/http"
	"api/common/utils/security"
	"api/config"
	"fmt"
)

// Verifies a Facebook token and returns the associated user information.
//
// Refer to the official documentation for more details
// Debug: https://graph.facebook.com/debug_token?input_token=YOUR_AUTH_TOKEN&access_token=YOUR_APP_ID|YOUR_CLIENT_SECRET
// Profile: https://graph.facebook.com/me?fields=id,name,last_name,first_name,email,languages,picture.width(100).height(100).as(picture_small),picture.width(720).height(720).as(picture_large)&access_token=YOUR_AUT_TOKEN
func VerifyFacebookToken(token string) (*types.FacebookUserProfileResponse, error) {
	// Retrieve token info
	if len(token) <= 0 {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	debugResp := &types.FacebookDebugAccessTokenResponse{}
	errDebug := http.HTTPGet(
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

	// Validate AppId, Application, DataAccessExpires, Expires, IsValid
	if !debugResp.Data.IsValid || debugResp.Data.AppId != config.Env.FacebookAppId || debugResp.Data.Application != config.Env.FacebookAppName {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	// Validate scopes
	if !utils.IsFacebookLoginScopesValid(debugResp.Data.Scopes) {
		return nil, fmt.Errorf("%s%v",
			"Invalid scopes! You need to enable these scopes: ",
			constants.AUTH_LOGIN_WITH_FACEBOOK_REQUIRED_SCOPES,
		)
	}

	// Retrieve user info
	userResp := &types.FacebookUserProfileResponse{}
	secretProof, errSecretProof := security.EncodeHMAC_SHA256(token, config.Env.FacebookClientSecret)
	if errSecretProof != nil {
		return nil, constants.HTTP_500_ERROR_MESSAGE("encode Facebook HMAC HS256 secret proof")
	}
	errUser := http.HTTPGet(
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
	userResp.Expires = debugResp.Data.Expires
	userResp.DataAccessExpires = debugResp.Data.DataAccessExpires
	return userResp, nil
}
