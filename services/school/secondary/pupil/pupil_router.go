package pupil

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/secondary/pupil/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/pupils",
		Tag:   []string{"Pupils"},
	}
	const tableName = "pupils"

	// Create pupil
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-pupil",
			Summary:     "Create pupil" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Create new pupil by providing name and description and return created object. The name pupil should be unique.",
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
				Body data.CreatePupilRequest
			},
		) (*struct{ Body data.PupilResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.PupilResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update pupil with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-pupil",
			Summary:     "Update pupil" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Update existing pupil with matching id and return the new pupil object.",
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
				data.PupilID
				Body data.UpdatePupilRequest
			},
		) (*struct{ Body data.PupilResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.PupilResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete pupil with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-pupil",
			Summary:     "Delete pupil" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Delete existing pupil with matching id and return affected rows in database.",
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
				data.PupilID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get pupil by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-pupil-id",
			Summary:     "Get pupil by id" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Return one pupil with matching id",
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
				data.PupilID
			},
		) (*struct{ Body data.PupilResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.PupilResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all pupils
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-pupil-list",
			Summary:     "Get all pupils" + " (" + constants.FeatureDirectorLabel + ")",
			Description: "Get all pupils with support for search, filter and pagination",
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
			Body data.PupilResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PupilResponseList
			}{Body: *result}, nil
		},
	)
}
