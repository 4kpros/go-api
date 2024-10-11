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

func SetupEndpoints(
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
			Description: "Create new permission by providing name and description and return created object. The name permission should be unique.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s", endpointConfig.Group),
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
			var result, errCode, err = controller.Create(&input.Body)
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
			var result, errCode, err = controller.Update(input.PermissionId.Id, &input.Body)
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
		) (*struct{ Body types.DeleteResponse }, error) {
			var result, errCode, err = controller.Delete(input.Id)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeleteResponse }{Body: types.DeleteResponse{AffectedRows: result}}, nil
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
			var result, errCode, err = controller.GetById(input.Id)
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
			OperationID: "get-permissions",
			Summary:     "Get all permissions",
			Description: "Get all permissions with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s", endpointConfig.Group),
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
			Body data.PermissionsResponse
		}, error) {
			var result, errCode, err = controller.GetAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionsResponse
			}{Body: *result}, nil
		},
	)
}