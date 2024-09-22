package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/gin-gonic/gin"
)

func IsOriginKnown(host string) bool {
	hosts := strings.Split(config.AppEnv.AllowedHosts, ",")
	return slices.Contains(hosts, host)
}

func HaveSufficientPermissions(role string, allowedRoles []string) bool {
	return slices.Contains(allowedRoles, role)
}

func SecureAPIHandler(handler gin.HandlerFunc, requiredAuth bool, allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add headers for more security
		c.Header("Content-Type", "application/json")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Xss-Protection", "1; mode=block")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("X-Download-Options", "noopen")
		c.Header("Strict-Transport-Security", fmt.Sprintf("max-age=%d; %s", 31536000, "includeSubDomains"))
		c.Next()

		// Check allowed hosts
		if !IsOriginKnown(c.Request.Host) {
			message := "Our system detected your request as malicious! Please fix that before."
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
			return
		}

		// Check api key
		apiKey := utils.ExtractApiKeyHeader(c)
		if apiKey != config.AppEnv.ApiKey {
			message := "Invalid API key! Please enter valid API key and try again."
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
			return
		}

		// Check if it required to be logged
		if requiredAuth {
			JWTHandler(c, handler, allowedRoles)
		} else {
			handler(c)
		}
	}
}

func JWTHandler(c *gin.Context, handler gin.HandlerFunc, allowedRoles []string) {
	// Check if the token is valid
	bearerToken := utils.ExtractBearerTokenHeader(c)
	if len(bearerToken) <= 0 {
		message := "Missing or bad authorization header! Please enter authorization header and try again."
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", message))
		return
	}
	jwtDecrypted, jwtValid := utils.VerifyJWTToken(bearerToken, config.AppPem.JwtPublicKey)
	if !jwtValid || jwtDecrypted == nil {
		message := "Invalid or expired authorization header! Please enter valid authorization header and try again."
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", message))
		return
	}

	// Check if user have sufficient permissions(roles)
	if !HaveSufficientPermissions(jwtDecrypted.Role, allowedRoles) {
		message := "Invalid permissions! You don't have permission to access this resource."
		c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
		return
	} else {
		handler(c)
	}
}
