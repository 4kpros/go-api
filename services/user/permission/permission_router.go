package permission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/data"
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

	// Update permission feature
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-permission-feature",
			Summary:     "Update permission feature" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Update permission feature with matching role id and feature name",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/role/{roleID}/feature", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
						tableName,                  // Table name
						constants.PermissionUpdate, // Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.PermissionPathRequest
				Body data.UpdatePermissionFeatureRequest
			},
		) (*struct {
			Body data.PermissionFeatureResponse
		}, error) {
			result, errCode, err := controller.UpdatePermissionFeature(
				&ctx, input,
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionFeatureResponse
			}{Body: *result}, nil
		},
	)

	// Update permission table
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-permission-table",
			Summary:     "Update permission table" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Update permission table with matching role id and table name with actions(CRUD)",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/role/{roleID}/table", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
						tableName,                  // Table name
						constants.PermissionUpdate, // Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.PermissionPathRequest
				Body data.UpdatePermissionTableRequest
			},
		) (*struct {
			Body data.PermissionTableResponse
		}, error) {
			result, errCode, err := controller.UpdatePermissionTable(
				&ctx, input,
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionTableResponse
			}{Body: *result}, nil
		},
	)

	// Get all permissions for role
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-role-permission-list",
			Summary:     "Get all role permissions" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Get all permissions with matching role id and support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/role/{roleID}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
						tableName,                // Table name
						constants.PermissionRead, // Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden},
		},
		func(
			ctx context.Context,
			input *struct {
				data.PermissionPathRequest
				types.Filter
				types.PaginationRequest
			},
		) (*struct {
			Body data.PermissionListResponse
		}, error) {
			result, errCode, err := controller.GetAllByRoleID(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionListResponse
			}{Body: *result}, nil
		},
	)
}
