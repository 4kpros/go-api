package role

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/user/role/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/roles",
		Tag:   []string{"Roles"},
	}
	const tableName = "roles"

	// Create role
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-role",
			Summary:     "Create role",
			Description: "Create new role by providing name and description and return created object. The name role should be unique.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        append(endpointConfig.Tag, constants.FeatureAdmin),
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
						tableName,                  // Table name
						constants.PermissionCreate, // Operation
					},
				},
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.RoleRequest
			},
		) (*struct{ Body data.RoleResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.RoleResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update role with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-role",
			Summary:     "Update role",
			Description: "Update existing role with matching id and return the new role object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        append(endpointConfig.Tag, constants.FeatureAdmin),
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				data.RoleId
				Body data.RoleRequest
			},
		) (*struct{ Body data.RoleResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.RoleResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete role with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-role",
			Summary:     "Delete role",
			Description: "Delete existing role with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        append(endpointConfig.Tag, constants.FeatureAdmin),
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
						tableName,                  // Table name
						constants.PermissionDelete, // Operation
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
				data.RoleId
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get role by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-role-id",
			Summary:     "Get role by id",
			Description: "Return one role with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        append(endpointConfig.Tag, constants.FeatureAdmin),
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
						tableName,                // Table name
						constants.PermissionRead, // Operation
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
				data.RoleId
			},
		) (*struct{ Body data.RoleResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.RoleResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all roles
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-role-list",
			Summary:     "Get all roles",
			Description: "Get all roles with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        append(endpointConfig.Tag, constants.FeatureAdmin),
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
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
				types.Filter
				types.PaginationRequest
			},
		) (*struct {
			Body data.RoleResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.RoleResponseList
			}{Body: *result}, nil
		},
	)
}
