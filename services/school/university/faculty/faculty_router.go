package faculty

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/faculty/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/faculties",
		Tag:   []string{"Faculties"},
	}
	const tableName = "faculties"

	// Create faculty
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-faculty",
			Summary:     "Create faculty" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Create new faculty by providing name and description and return created object. The name faculty should be unique.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureDirector,  // Feature scope
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
				Body data.CreateFacultyRequest
			},
		) (*struct{ Body data.FacultyResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.FacultyResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update faculty with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-faculty",
			Summary:     "Update faculty" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Update existing faculty with matching id and return the new faculty object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureDirector,  // Feature scope
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
				data.FacultyID
				Body data.UpdateFacultyRequest
			},
		) (*struct{ Body data.FacultyResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.FacultyResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete faculty with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-faculty",
			Summary:     "Delete faculty" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Delete existing faculty with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureDirector,  // Feature scope
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
				data.FacultyID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get faculty by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-faculty-id",
			Summary:     "Get faculty by id" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Return one faculty with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureDirector, // Feature scope
						tableName,                 // Table name
						constants.PermissionRead,  // Operation
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
				data.FacultyID
			},
		) (*struct{ Body data.FacultyResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.FacultyResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all faculties
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-faculty-list",
			Summary:     "Get all faculties" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Get all faculties with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureDirector, // Feature scope
						tableName,                 // Table name
						constants.PermissionRead,  // Operation
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
			Body data.FacultyResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.FacultyResponseList
			}{Body: *result}, nil
		},
	)
}
