package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/user/user/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/users",
		Tag:   []string{"Users"},
	}
	const tableName = "users"

	// Create user with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-user-email",
			Summary:     "Create user with email" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Create new user by providing email and user and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/email", endpointConfig.Group),
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
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.CreateUserWithEmailRequest
			},
		) (*struct{ Body data.UserResponse }, error) {
			result, errCode, err := controller.CreateWithEmail(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Create user with phone number
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "create-user-phone",
			Summary:     "Create user with phone" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Create new user by providing phone number and role and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/phone", endpointConfig.Group),
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
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.CreateUserWithPhoneNumberRequest
			},
		) (*struct{ Body data.UserResponse }, error) {
			result, errCode, err := controller.CreateWithPhoneNumber(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update user with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-user",
			Summary:     "Update user" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Update existing user with matching id and return the new user object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
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
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserID
				Body data.UpdateUserRequest
			},
		) (*struct{ Body data.UserResponse }, error) {
			result, errCode, err := controller.UpdateUser(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete user with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-user",
			Summary:     "Delete user" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Delete existing user with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
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
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get user by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-user-id",
			Summary:     "Get user by id" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Return one user with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
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
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UserID
			},
		) (*struct{ Body data.UserResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all users
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-user-list",
			Summary:     "Get all users" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Get all users with support for search, filter and pagination",
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
			MaxBodyBytes:  constants.DefaultBodySize,
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
			Body data.UserResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UserResponseList
			}{Body: *result}, nil
		},
	)
}
