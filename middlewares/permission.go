package middlewares

import (
	"api/common/helpers"
	"api/services/user/permission"
	"api/services/user/role"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
)

// PermissionMiddleware Checks resource permissions
func PermissionMiddleware(api huma.API, roleRepo *role.Repository, permissionRepo *permission.Repository) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Retrieve jwtToken
		ctxContext := ctx.Context()
		jwtToken := helpers.GetJwtContext(&ctxContext)
		if jwtToken == nil || jwtToken.UserID <= 0 {
			next(ctx)
			return
		}

		// Retrieve feature permissions
		var featureScope, tableName, tableOperation string
		for _, opScheme := range ctx.Operation().Security {
			if securityScheme, ok := opScheme[constants.SecurityAuthName]; ok {
				if len(securityScheme) > 0 {
					featureScope = securityScheme[0]
					if len(securityScheme) > 1 {
						tableName = securityScheme[1]
						if len(securityScheme) > 2 {
							tableOperation = securityScheme[2]
						}
					}
				}
				break
			}
		}

		// Check for required permissions
		if len(featureScope) >= 1 {
			// Retrieve role
			role, errFeature := roleRepo.GetByID(jwtToken.RoleID)
			if errFeature != nil {
				tempErr := constants.Http500ErrorMessage("get role from database")
				_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, tempErr.Error(), tempErr)
				return
			}
			tempErr := constants.Http403InvalidPermissionErrorMessage()
			if !(role != nil && role.Feature == featureScope) {
				_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}

			if len(tableName) >= 1 {
				// Retrieve permissions for role
				userPermission, errPerm := permissionRepo.GetPermissionOR(jwtToken.RoleID, tableName, "*")
				if errPerm != nil {
					tempErr := constants.Http500ErrorMessage("get permission from database")
					_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, tempErr.Error(), tempErr)
					return
				}
				if !(userPermission != nil && (userPermission.TableName == tableName || userPermission.TableName == "*")) {
					_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
					return
				}

				if len(tableOperation) >= 1 {
					if tableOperation == constants.PermissionCreate && !userPermission.Create {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionRead && !userPermission.Read {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionUpdate && !userPermission.Update {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionDelete && !userPermission.Delete {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					}
				}
			}
		}

		next(ctx)

	}
}
