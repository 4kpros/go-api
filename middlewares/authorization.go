package middlewares

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/utils/security"
	"api/config"
	"fmt"
	"net/http"
	"slices"

	"github.com/danielgtaylor/huma/v2"
)

// Handles authentication for API requests.
func AuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	var errMessage string
	return func(ctx huma.Context, next func(huma.Context)) {
		// Check if current endpoint require authorization
		isAuthorizationRequired := false
		for _, opScheme := range ctx.Operation().Security {
			if _, ok := opScheme[constants.SECURITY_AUTH_NAME]; ok {
				isAuthorizationRequired = true
				break
			}
		}
		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		// Parse and decode the token
		token := helpers.ExtractBearerTokenHeader(&ctx)
		if len(token) < 1 {
			errMessage = "Missing or bad authorization header! Please enter valid information."
			huma.WriteErr(api, ctx, http.StatusUnauthorized, errMessage, fmt.Errorf("%s", errMessage))
			return
		}
		jwtDecoded, errDecoded := security.DecodeJWTToken(
			token,
			config.Keys.JwtPublicKey,
		)
		if errDecoded != nil || jwtDecoded == nil {
			tempErr := constants.HTTP_401_INVALID_TOKEN_ERROR_MESSAGE()
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
			return
		}

		// Validate the token by checking if it's cached
		isTokenCached := false
		if jwtDecoded.Issuer == constants.JWT_ISSUER_SESSION {
			isTokenCached = security.ValidateJWTToken(
				token,
				jwtDecoded,
				config.CheckValueInRedisList(token),
			)
		} else if slices.Contains(constants.JWT_ISSUER_AUTH, jwtDecoded.Issuer) {
			isTokenCached = security.ValidateJWTToken(
				token,
				jwtDecoded,
				config.GetRedisString,
			)
		}
		if isTokenCached {
			next(*helpers.SetAuthContext(&ctx, token, jwtDecoded))
			return
		}

		tempErr := constants.HTTP_401_INVALID_TOKEN_ERROR_MESSAGE()
		_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
	}
}
