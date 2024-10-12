package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
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
				Body data.UserWithEmailRequest
			},
		) (*struct{ Body model.User }, error) {
			result, errCode, err := controller.CreateWithEmail(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
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
				Body data.UserWithPhoneNumberRequest
			},
		) (*struct{ Body model.User }, error) {
			result, errCode, err := controller.CreateWithPhoneNumber(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
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
				data.UserId
				Body data.UserRequest
			},
		) (*struct{ Body model.User }, error) {
			inputFormatted := &model.User{
				Email:       input.Body.Email,
				PhoneNumber: input.Body.PhoneNumber,
				Language:    input.Body.Language,
				RoleId:      input.Body.RoleId,
			}
			inputFormatted.ID = input.UserId.Id
			result, errCode, err := controller.UpdateUser(inputFormatted)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
		},
	)

	// Update user info with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-user-info",
			Summary:     "Update user info",
			Description: "Update existing user info with matching id and return the new user object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/info/:id", endpointConfig.Group),
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
				Body data.UserInfoRequest
			},
		) (*struct{ Body model.UserInfo }, error) {
			inputFormatted := &model.UserInfo{
				UserName:  input.Body.UserName,
				FirstName: input.Body.FirstName,
				LastName:  input.Body.LastName,
				Address:   input.Body.Address,
			}
			inputFormatted.ID = input.UserId.Id
			result, errCode, err := controller.UpdateUserInfo(inputFormatted)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.UserInfo }{Body: *result}, nil
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
				data.UserId
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&input.UserId)
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
				data.UserId
			},
		) (*struct{ Body model.User }, error) {
			result, errCode, err := controller.GetById(&input.UserId)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body model.User }{Body: *result}, nil
		},
	)

	// Get all users
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-users",
			Summary:     "Get all users",
			Description: "Get all users with support for search, filter and pagination",
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
			Body data.UsersResponse
		}, error) {
			result, errCode, err := controller.GetAll(&input.Filter, &input.PaginationRequest)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UsersResponse
			}{Body: *result}, nil
		},
	)
}
