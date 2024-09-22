package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerTokenHeader(c *gin.Context) string {
	return strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
}

func ExtractApiKeyHeader(c *gin.Context) string {
	return c.GetHeader("X-API-Key")
}
