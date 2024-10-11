package utils

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

// Retrieves the bearer token from the current request context.
func ExtractBearerTokenHeader(ctx *huma.Context) string {
	return strings.TrimPrefix((*ctx).Header("Authorization"), "Bearer ")
}
