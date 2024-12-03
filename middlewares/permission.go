package middlewares

import (
	"api/common/helpers"
	"api/services/user/permission"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
)

// PermissionMiddleware Checks resource permissions
func PermissionMiddleware(api huma.API, repository *permission.Repository) func(huma.Context, func(huma.Context)) {
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
			// Retrieve permission feature for role
			permissionFeature, errFeature := repository.GetPermissionFeature(jwtToken.RoleID, featureScope)
			tempErr := constants.Http403InvalidPermissionErrorMessage()
			if errFeature != nil || permissionFeature == nil || permissionFeature.Feature != featureScope {
				_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}

			if len(tableName) >= 1 {
				// Retrieve permission table for role
				permissionTable, errTable := repository.GetPermissionTable(jwtToken.RoleID, tableName)
				if errTable != nil || permissionTable == nil || (permissionTable.TableName != tableName && permissionTable.TableName != "*") {
					_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
					return
				}

				if len(tableOperation) >= 1 {
					if tableOperation == constants.PermissionCreate && !permissionTable.Create {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionRead && !permissionTable.Read {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionUpdate && !permissionTable.Update {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionDelete && !permissionTable.Delete {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					}
				}
			}
		}

		next(ctx)

	}
}
