package middlewares

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/services/admin/permission"
)

// PermissionMiddleware Checks resource permissions
func PermissionMiddleware(api huma.API, repository *permission.Repository) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Retrieve jwtToken, table permission, permission type
		ctxContext := ctx.Context()
		jwtToken := helpers.GetJwtContext(&ctxContext)

		// Check if this endpoint require permissions
		if jwtToken == nil || jwtToken.UserId <= 0 || len(jwtToken.Issuer) <= 0 || jwtToken.RoleId <= 0 {
			next(ctx)
			return
		}

		// Check permission
		permissionTable, _ := ctx.Operation().Metadata[constants.PERMISSION_METADATA_TABLE_KEY].(string)
		permissionType, _ := ctx.Operation().Metadata[constants.PERMISSION_METADATA_TYPE_KEY].(string)
		if checkUserPermissions(repository, jwtToken, permissionTable, permissionType) {
			next(ctx)
			return
		}

		tempErr := constants.Http401InvalidPermissionErrorMessage()
		_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
	}
}

// Checks if the user(from JWT token) has permission to access the table name
func checkUserPermissions(repository *permission.Repository, jwtToken *types.JwtToken, table string, permissionType string) bool {
	// If there is no required permission, return true
	if len(table) <= 0 || len(permissionType) <= 0 {
		return true
	}

	if jwtToken.Issuer == constants.JwtIssuerSession {
		foundPermission, _ := repository.GetByRoleIdTable(jwtToken.RoleId, table)
		if foundPermission != nil && (foundPermission.Table == "*" || foundPermission.Table == table) && foundPermission.RoleId == jwtToken.RoleId {
			if permissionType == constants.PermissionCreate && foundPermission.Create {
				return true
			}
			if permissionType == constants.PermissionRead && foundPermission.Read {
				return true
			}
			if permissionType == constants.PermissionUpdate && foundPermission.Update {
				return true
			}
			if permissionType == constants.PermissionDelete && foundPermission.Delete {
				return true
			}
		}
	}
	return false
}
