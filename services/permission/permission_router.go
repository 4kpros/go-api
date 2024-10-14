package permission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/permission/data"
	"github.com/4kpros/go-api/services/permission/model"
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

	// Create permission
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-permission",
			Summary:     "Create permission",
			Description: "Create new permission for role. You will need the table name and set read, write, update and delete values",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.CreatePermissionRequest
			},
		) (*struct{ Body model.Permission }, error) {
			result, errCode, err := controller.Create(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Permission }{Body: *result}, nil
		},
	)

	// Update permission with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-permission",
			Summary:     "Update permission",
			Description: "Update existing permission with matching id and return the new permission object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/:id", endpointConfig.Group),
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
				data.PermissionId
				Body data.UpdatePermissionRequest
			},
		) (*struct{ Body model.Permission }, error) {
			result, errCode, err := controller.Update(input.PermissionId.ID, &input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Permission }{Body: *result}, nil
		},
	)

	// Delete permission with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-permission",
			Summary:     "Delete permission",
			Description: "Delete existing permission with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/:id", endpointConfig.Group),
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
				data.PermissionId
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(input.ID)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get permission by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-permission-id",
			Summary:     "Get permission by id",
			Description: "Return one permission with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/:id", endpointConfig.Group),
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
				data.PermissionId
			},
		) (*struct{ Body model.Permission }, error) {
			result, errCode, err := controller.GetById(input.ID)
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
			Body data.PermissionList
		}, error) {
			result, errCode, err := controller.GetAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionList
			}{Body: *result}, nil
		},
	)
}
