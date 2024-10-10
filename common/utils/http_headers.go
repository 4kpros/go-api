package utils

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

func ExtractBearerTokenHeader(ctx *huma.Context) string {
	return strings.TrimPrefix((*ctx).Header("Authorization"), "Bearer ")
}
