package permission

import (
	"api/services/admin"
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/admin/permission/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/permissions",
		Tag:   []string{"Permissions"},
	}
	const tableName = "permissions"

	// Update permission
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-role-feature-permission",
			Summary:     "Update permission",
			Description: "Update permission with matching role id and feature name",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/role/{roleId}/{featureName}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						admin.FeaturePermission,    // Feature scope
						tableName,                  // Table name
						constants.PermissionUpdate, // Operation
					},
				},
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UpdateRoleFeaturePermissionPathRequest
				Body data.UpdateRoleFeaturePermissionBodyRequest
			},
		) (*struct {
			Body data.PermissionFeatureTableResponse
		}, error) {
			result, errCode, err := controller.UpdateByRoleIdFeatureName(
				&ctx, input.RoleId, input.FeatureName, input.Body,
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionFeatureTableResponse
			}{Body: *result}, nil
		},
	)

	// Get all permissions for role
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-role-permission-list",
			Summary:     "Get all role permissions",
			Description: "Get all permissions with matching role id and support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/role/{roleId}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						admin.FeaturePermission,  // Feature scope
						tableName,                // Table name
						constants.PermissionRead, // Operation
					},
				},
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden},
		},
		func(
			ctx context.Context,
			input *struct {
				data.GetRolePermissionListRequest
				types.Filter
				types.PaginationRequest
			},
		) (*struct {
			Body data.PermissionListResponse
		}, error) {
			result, errCode, err := controller.GetAllByRoleId(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionListResponse
			}{Body: *result}, nil
		},
	)
}
