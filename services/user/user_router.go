package user

import (
	"context"
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/data"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *UserController,
) {
	var endpointConfig = types.APIEndpointConfig{
		Group: "/users",
		Tag:   []string{"Users"},
	}

	// Create user with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-user-email",
			Summary:     "Create user with email",
			Description: "Create new user by providing email and user and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/email", endpointConfig.Group),
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
			Summary:     "Create user with phone",
			Description: "Create new user by providing phone number and role and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/phone", endpointConfig.Group),
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
			Summary:     "Update user",
			Description: "Update existing user with matching id and return the new user object.",
			Method:      http.MethodPut,
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
				data.UserId
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
			Summary:     "Delete user",
			Description: "Delete existing user with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
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
				data.UserId
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
			Summary:     "Get user by id",
			Description: "Return one user with matching id",
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
				data.UserId
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
			Summary:     "Get all users",
			Description: "Get all users with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SECURITY_AUTH_NAME: { // Used for authentication
						constants.PERMISSION_TABLE_NAME_USER, // Used for permission table name
						constants.PERMISSION_READ,            // Used for permission type
					},
				},
			},
			// Metadata: map[string]any{
			// 	"permissionTable": constants.PERMISSION_TABLE_NAME_USER,
			// 	"permissionName":  constants.PERMISSION_READ,
			// },
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
