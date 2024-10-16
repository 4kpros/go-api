package profile

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
	controller *ProfileController,
) {
	var endpointConfig = types.APIEndpointConfig{
		Group: "/profile",
		Tag:   []string{"Profile"},
	}

	// Update profile
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile",
			Summary:     "Update profile",
			Description: "Update user profile such as email, phone number and password",
			Method:      http.MethodPut,
			Path:        endpointConfig.Group,
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
				Body data.UpdateProfileRequest
			},
		) (*struct{ Body data.UserResponse }, error) {
			result, errCode, err := controller.UpdateProfile(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update profile info
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-info",
			Summary:     "Update profile info",
			Description: "Update profile information such as username, first name, last name, address, language",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/info", endpointConfig.Group),
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
				Body data.UpdateProfileInfoRequest
			},
		) (*struct{ Body data.UserInfoResponse }, error) {
			result, errCode, err := controller.UpdateProfileInfo(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserInfoResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update profile MFA
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-mfa",
			Summary:     "Update profile MFA",
			Description: "Update profile MFA",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/mfa", endpointConfig.Group),
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
				Body data.UpdateProfileMfaRequest
			},
		) (*struct{ Body data.UserMfaResponse }, error) {
			result, errCode, err := controller.UpdateProfileMfa(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserMfaResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete account
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-profile",
			Summary:     "Delete user account",
			Description: "Delete current user account with provided bearer token",
			Method:      http.MethodDelete,
			Path:        endpointConfig.Group,
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
			input *struct{},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteProfile(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get profile
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-profile",
			Summary:     "Get profile info",
			Description: "Retrieve profile information for the current user with provided bearer token",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
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
			input *struct{},
		) (*struct{ Body data.UserResponse }, error) {
			result, errCode, err := controller.GetProfile(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)
}
