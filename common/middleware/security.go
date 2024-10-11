package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/danielgtaylor/huma/v2"
)

// Utility function for CORS
func isOriginKnown(host string) bool {
	hosts := strings.Split(config.Env.AllowedHosts, ",")
	return slices.Contains(hosts, host)
}

// Middleware to add more security like CORS, header XSS protection, ...
func SecureHeadersMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	var errMessage string
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
			errMessage = "Our system detected your request as malicious! Please fix that before."
			huma.WriteErr(api, ctx, http.StatusForbidden, errMessage, fmt.Errorf("%s", errMessage))
			return
		}

		next(ctx)
	}
}

// Middleware for rate limiting based on IP address
func RateLimitMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	var errMessage string
	return func(ctx huma.Context, next func(huma.Context)) {
		if 1 == 2 {
			errMessage = "Our system detected your request as malicious! Your are banned for 24hours."
			huma.WriteErr(api, ctx, http.StatusForbidden, errMessage, fmt.Errorf("%s", errMessage))
			return
		}
		next(ctx)
	}
}

// Middleware to check authorization if it's required on every endpoint
func AuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	var errMessage string
	return func(ctx huma.Context, next func(huma.Context)) {
		// Check if current endpoint require authorization
		isAuthorizationRequired := false
		for _, opScheme := range ctx.Operation().Security {
			var ok bool
			if _, ok = opScheme[constants.SECURITY_AUTH_NAME]; ok {
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
			var tempErr = constants.HTTP_401_ERROR_MESSAGE()
			huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
			return
		}

		// Validate the token by checking if it's cached
		var isTokenCached bool = false
		if jwtDecoded.Issuer == constants.JWT_ISSUER_SESSION {
			isTokenCached = utils.ValidateJWTToken(
				token,
				jwtDecoded,
				config.CheckRedisStrFromArrayStr(token),
			)
		} else if jwtDecoded.Issuer == constants.JWT_ISSUER_SESSION_GENERATED {
			isTokenCached = utils.ValidateJWTToken(
				token,
				jwtDecoded,
				config.GetRedisStr,
			)
		}
		if isTokenCached {
			next(ctx)
			return
		}

		var tempErr = constants.HTTP_401_ERROR_MESSAGE()
		huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
	}
}
