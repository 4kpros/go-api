package role

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/role/data"
	"github.com/4kpros/go-api/services/role/model"
	"github.com/danielgtaylor/huma/v2"
)

func SetupEndpoints(
	humaApi *huma.API,
	controller *RoleController,
) {
	var endpointConfig = struct {
		Group string
		Tag   []string
	}{
		Group: "/roles",
		Tag:   []string{"Roles"},
	}
	const requireAuth = true

	// Create role
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "post-role",
			Summary:       "Create role",
			Description:   "Create new role by providing name and description and return created object. The name role should be unique.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s ", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.RoleRequest
			},
		) (*struct{ Body model.Role }, error) {
			result, errCode, err := controller.Create(&input.Body)
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
			OperationID:   "update-role",
			Summary:       "Update role",
			Description:   "Update existing role with matching id and return the new role object.",
			Method:        http.MethodPut,
			Path:          fmt.Sprintf("%s/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.Id
				Body data.RoleRequest
			},
		) (*struct{ Body model.Role }, error) {
			var inputFormatted = &model.Role{
				Name:        input.Body.Name,
				Description: input.Body.Name,
			}
			inputFormatted.ID = uint(input.Id.Id)
			result, errCode, err := controller.Update(inputFormatted)
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
			OperationID:   "delete-role",
			Summary:       "Delete role",
			Description:   "Delete existing role with matching id and return affected rows in database.",
			Method:        http.MethodDelete,
			Path:          fmt.Sprintf("%s/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.Id
			},
		) (*struct{ Body data.DeleteResponse }, error) {
			result, errCode, err := controller.Delete(&input.Id)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.DeleteResponse }{Body: data.DeleteResponse{AffectedRows: result}}, nil
		},
	)

	// Find role by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "find-role-id",
			Summary:       "Find role by id",
			Description:   "Return one role with matching id",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/:id", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.Id
			},
		) (*struct{ Body model.Role }, error) {
			result, errCode, err := controller.FindById(&input.Id)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.Role }{Body: *result}, nil
		},
	)

	// Find all roles
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "find-roles",
			Summary:       "Find all roles",
			Description:   "Find all roles with support for search, filter and pagination",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s ", endpointConfig.Group),
			Middlewares:   *middleware.GenerateMiddlewares(requireAuth),
			Tags:          endpointConfig.Tag,
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
			Body data.GetAllResponse
		}, error) {
			result, errCode, err := controller.FindAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.GetAllResponse
			}{Body: *result}, nil
		},
	)
}
