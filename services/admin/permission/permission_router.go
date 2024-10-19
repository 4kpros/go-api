package permission

import (
	data2 "api/services/admin/permission/data"
	"api/services/admin/permission/model"
	"context"
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *PermissionController,
) {
	var endpointConfig = types.APIEndpointConfig{
		Group: "/permissions",
		Tag:   []string{"Permissions"},
	}

	// Update permission
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-permission",
			Summary:     "Update permission",
			Description: "Update existing role permission.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data2.UpdatePermissionRequest
			},
		) (*struct{ Body model.Permission }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Permission }{Body: *result}, nil
		},
	)

	// Get permission by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-permission",
			Summary:     "Get permission by id",
			Description: "Return one permission with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data2.PermissionId
			},
		) (*struct{ Body model.Permission }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Permission }{Body: *result}, nil
		},
	)

	// Get all permissions
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-permission-list",
			Summary:     "Get all permissions",
			Description: "Get all permissions with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
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
			Body data2.PermissionList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data2.PermissionList
			}{Body: *result}, nil
		},
	)
}
