package role

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/role/data"
	"github.com/4kpros/go-api/services/role/model"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *RoleController,
) {
	var endpointConfig = types.APIEndpointConfig{
		Group: "/roles",
		Tag:   []string{"Roles"},
	}

	// Create role
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-role",
			Summary:     "Create role",
			Description: "Create new role by providing name and description and return created object. The name role should be unique.",
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
				Body data.RoleRequest
			},
		) (*struct{ Body model.Role }, error) {
			var result, errCode, err = controller.Create(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Role }{Body: *result}, nil
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
				data.RoleId
				Body data.RoleRequest
			},
		) (*struct{ Body model.Role }, error) {
			var inputFormatted = &model.Role{
				Name:        input.Body.Name,
				Description: input.Body.Name,
			}
			inputFormatted.ID = input.RoleId.Id
			var result, errCode, err = controller.Update(inputFormatted)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Role }{Body: *result}, nil
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
				data.RoleId
			},
		) (*struct{ Body types.DeleteResponse }, error) {
			var result, errCode, err = controller.Delete(&input.RoleId)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeleteResponse }{Body: types.DeleteResponse{AffectedRows: result}}, nil
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
				data.RoleId
			},
		) (*struct{ Body model.Role }, error) {
			var result, errCode, err = controller.GetById(&input.RoleId)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Role }{Body: *result}, nil
		},
	)

	// Get all roles
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-roles",
			Summary:     "Get all roles",
			Description: "Get all roles with support for search, filter and pagination",
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
			Body data.RolesResponse
		}, error) {
			var result, errCode, err = controller.GetAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.RolesResponse
			}{Body: *result}, nil
		},
	)
}
