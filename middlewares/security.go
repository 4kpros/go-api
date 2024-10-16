package middlewares

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"api/common/constants"
	"api/common/helpers"
	"api/common/utils"
	"api/config"

	"github.com/danielgtaylor/huma/v2"
)

// Utility function for CORS
func isOriginKnown(host string) bool {
	hosts := strings.Split(config.Env.AllowedHosts, ",")
	return slices.Contains(hosts, host)
}

// Sets security-related HTTP headers for responses.
func SecureHeadersMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Add headers for more security
		ctx.SetHeader("Content-Type", "application/json")
		ctx.SetHeader("X-Frame-Options", "DENY")
		ctx.SetHeader("X-Content-Type-Options", "nosniff")
		ctx.SetHeader("X-Xss-Protection", "1; mode=block")
		ctx.SetHeader("Content-Security-Policy", "default-src 'self'")
		ctx.SetHeader("Referrer-Policy", "strict-origin-when-cross-origin")
		ctx.SetHeader("X-Download-Options", "noopen")
		ctx.SetHeader("Strict-Transport-Security", fmt.Sprintf("max-age=%d; %s", 31536000, "includeSubDomains"))

		// Check for allowed hosts
		if !isOriginKnown(ctx.Host()) {
			errMessage := " Our system detected your request as malicious! Please fix that before."
			huma.WriteErr(api, ctx, http.StatusForbidden, errMessage, fmt.Errorf("%s", errMessage))
			return
		}

		next(ctx)
	}
}

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
		token := utils.ExtractBearerTokenHeader(&ctx)
		if len(token) < 1 {
			errMessage = "Missing or bad authorization header! Please enter valid information."
			huma.WriteErr(api, ctx, http.StatusUnauthorized, errMessage, fmt.Errorf("%s", errMessage))
			return
		}
		jwtDecoded, errDecoded := utils.DecodeJWTToken(
			token,
			config.Keys.JwtPublicKey,
		)
		if errDecoded != nil || jwtDecoded == nil {
			tempErr := constants.HTTP_401_INVALID_TOKEN_ERROR_MESSAGE()
			huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
			return
		}

		// Validate the token by checking if it's cached
		isTokenCached := false
		if jwtDecoded.Issuer == constants.JWT_ISSUER_SESSION {
			isTokenCached = utils.ValidateJWTToken(
				token,
				jwtDecoded,
				config.CheckValueInRedisList(token),
			)
		} else if slices.Contains(constants.JWT_ISSUER_AUTH, jwtDecoded.Issuer) {
			isTokenCached = utils.ValidateJWTToken(
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
		huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
	}
}
