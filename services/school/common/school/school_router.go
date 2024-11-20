package school

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/school/data"
	"api/services/school/common/school/model"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools",
		Tag:   []string{"Schools"},
	}
	const tableName = "schools"

	// Create school
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-school",
			Summary:     "Create school" + constants.FeatureAdmin,
			Description: "Create new school by providing name and description and return created object. The name school should be unique.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
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
				Body data.SchoolRequest
			},
		) (*struct{ Body data.SchoolResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SchoolResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Add school director
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-school-director",
			Summary:     "Add school director" + constants.FeatureAdmin,
			Description: "Add new user as director in order to manage school.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/director", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
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
				Body data.DirectorRequest
			},
		) (*struct{ Body model.SchoolDirector }, error) {
			result, errCode, err := controller.AddDirector(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.SchoolDirector }{Body: *result}, nil
		},
	)

	// Update school with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-school",
			Summary:     "Update school" + constants.FeatureAdmin,
			Description: "Update existing school with matching id and return the new school object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
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
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.SchoolId
				Body data.SchoolRequest
			},
		) (*struct{ Body data.SchoolResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SchoolResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete school with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-school",
			Summary:     "Delete school" + constants.FeatureAdmin,
			Description: "Delete existing school with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
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
				data.SchoolId
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete school director
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-school-director",
			Summary:     "Delete director" + constants.FeatureAdmin,
			Description: "Delete director with matching school id and user id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/director/{url}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
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
				data.DirectorRequest
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteDirector(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get school by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-school-id",
			Summary:     "Get school by id" + constants.FeatureAdmin,
			Description: "Return one school with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
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
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.SchoolId
			},
		) (*struct{ Body data.SchoolResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SchoolResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all schools
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-school-list",
			Summary:     "Get all schools" + constants.FeatureAdmin,
			Description: "Get all schools with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
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
			Body data.SchoolResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.SchoolResponseList
			}{Body: *result}, nil
		},
	)
}
