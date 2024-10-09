package middleware

import (
	"fmt"
	"slices"
	"strings"

	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/danielgtaylor/huma/v2"
)

var defaultMiddleware *huma.Middlewares = &huma.Middlewares{}
var authRequiredMiddleware *huma.Middlewares = &huma.Middlewares{JWTMiddleware}

func GenerateMiddlewares(requireAuth bool) *huma.Middlewares {
	if requireAuth {
		return authRequiredMiddleware
	}
	return defaultMiddleware
}

func IsOriginKnown(host string) bool {
	hosts := strings.Split(config.Env.AllowedHosts, ",")
	return slices.Contains(hosts, host)
}

func SecureApiMiddleware(ctx huma.Context, next func(huma.Context)) {
	// Add headers for more security
	ctx.SetHeader("Content-Type", "application/json")
	ctx.SetHeader("X-Frame-Options", "DENY")
	ctx.SetHeader("X-Content-Type-Options", "nosniff")
	ctx.SetHeader("X-Xss-Protection", "1; mode=block")
	ctx.SetHeader("Content-Security-Policy", "default-src 'self'")
	ctx.SetHeader("Referrer-Policy", "strict-origin-when-cross-origin")
	ctx.SetHeader("X-Download-Options", "noopen")
	ctx.SetHeader("Strict-Transport-Security", fmt.Sprintf("max-age=%d; %s", 31536000, "includeSubDomains"))

	// Check allowed hosts
	if !IsOriginKnown(ctx.Host()) {
		message := "Our system detected your request as malicious! Please fix that before."
		huma.Error403Forbidden(message, fmt.Errorf("%s", message))
		return
	}

	// Check api key
	apiKey := utils.ExtractApiKeyHeader(&ctx)
	if apiKey != config.Env.ApiKey {
		message := "Invalid API key! Please enter valid API key and try again."
		huma.Error403Forbidden(message, fmt.Errorf("%s", message))
		return
	}

	next(ctx)
}

func JWTMiddleware(ctx huma.Context, next func(huma.Context)) {
	// Check if the token is valid
	bearerToken := utils.ExtractBearerTokenHeader(&ctx)
	if len(bearerToken) <= 0 {
		message := "Missing or bad authorization header! Please enter authorization header and try again."
		huma.Error401Unauthorized(message, fmt.Errorf("%s", message))
		return
	}
	jwtDecrypted, jwtValid := utils.VerifyJWTToken(bearerToken, config.Keys.JwtPublicKey)
	if !jwtValid || jwtDecrypted == nil {
		message := "Invalid or expired authorization header! Please enter valid authorization header and try again."
		huma.Error401Unauthorized(message, fmt.Errorf("%s", message))
		return
	}

	next(ctx)
}
