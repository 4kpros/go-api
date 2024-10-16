package helpers

import (
	"context"

	"api/common/types"

	"github.com/danielgtaylor/huma/v2"
)

const (
	tokenKey  = "bearer"
	userIdKey = "userId"
	issuerKey = "issuer"
	roleIdKey = "roleId"
)

// Add information such as JWT token and bearer token to context in order
// to pass information to middleware, operation and handler func
func SetAuthContext(ctx *huma.Context, token string, jwtToken *types.JwtToken) *huma.Context {
	ctxToken := huma.WithValue(*ctx, tokenKey, token)
	ctxUserId := huma.WithValue(ctxToken, userIdKey, jwtToken.UserId)
	ctxIssuer := huma.WithValue(ctxUserId, issuerKey, jwtToken.Issuer)
	ctxRoleId := huma.WithValue(ctxIssuer, roleIdKey, jwtToken.RoleId)
	return &ctxRoleId
}

// Get JWT token from context
func GetJwtContext(ctx *context.Context) *types.JwtToken {
	result := &types.JwtToken{}
	if id, okId := (*ctx).Value(userIdKey).(int64); okId {
		result.UserId = id
	}
	if iss, okIss := (*ctx).Value(issuerKey).(string); okIss {
		result.Issuer = iss
	}
	if role, okRole := (*ctx).Value(roleIdKey).(int64); okRole {
		result.RoleId = role
	}
	return result
}

// Get Bearer token from context
func GetBearerContext(ctx *context.Context) string {
	if token, ok := (*ctx).Value(tokenKey).(string); ok {
		return token
	}
	return ""
}
