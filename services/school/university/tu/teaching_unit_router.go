package tu

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/tu/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/teaching-units",
		Tag:   []string{"Teaching Units"},
	}
	const tableName = "teaching_units"

	// Create teaching unit
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-teaching-unit",
			Summary:     "Create teaching unit" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Create new teaching unit and return the created object.",
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
				Body data.CreateTeachingUnitRequest
			},
		) (*struct{ Body data.TeachingUnitResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.TeachingUnitResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// add professor
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-teaching-unit-professor",
			Summary:     "Add new professor" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Add new professor and return the created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/{id}/professor", endpointConfig.Group),
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
				data.TeachingUnitID
				Body data.TeachingUnitProfessorRequest
			},
		) (*struct {
			Body data.TeachingUnitProfessorResponse
		}, error) {
			result, errCode, err := controller.AddProfessor(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TeachingUnitProfessorResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update teaching unit with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-teaching-unit",
			Summary:     "Update teaching unit" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Update existing teaching unit with matching id and return the updated object.",
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
				data.TeachingUnitID
				Body data.UpdateTeachingUnitRequest
			},
		) (*struct{ Body data.TeachingUnitResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.TeachingUnitResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete teaching unit with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-teaching-unit",
			Summary:     "Delete teaching unit" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Delete existing teaching unit and return affected rows in database.",
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
				data.TeachingUnitID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete teaching unit professor with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-teaching-unit-professor",
			Summary:     "Delete teaching unit professor" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Delete existing professor and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}/professor", endpointConfig.Group),
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
				data.TeachingUnitID
				Body data.TeachingUnitProfessorRequest
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteProfessor(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get teaching unit by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-teaching-unit-id",
			Summary:     "Get teaching unit by id" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Return one teaching unit with matching id",
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
				data.TeachingUnitID
			},
		) (*struct{ Body data.TeachingUnitResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.TeachingUnitResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all teaching units
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-teaching-unit-list",
			Summary:     "Get all teaching units" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Get all teaching units with support for search, filter and pagination",
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
			Body data.TeachingUnitResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TeachingUnitResponseList
			}{Body: *result}, nil
		},
	)
}
