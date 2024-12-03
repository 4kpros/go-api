package profile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/user/profile/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/profile",
		Tag:   []string{"Profile"},
	}

	// Update profile info
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-info",
			Summary:     "Update profile info",
			Description: "Update profile information such as username, first name, last name, address, ...",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/info", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						// Feature scope
						// Operation
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
				Body data.UpdateProfileInfoRequest
			},
		) (*struct{ Body data.UserProfileInfoResponse }, error) {
			result, errCode, err := controller.UpdateProfileInfo(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserProfileInfoResponse }{Body: *data.FromUserInfo(result)}, nil
		},
	)

	// Update profile MFA
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-mfa",
			Summary:     "Update profile MFA",
			Description: "Update profile multiple factor authentication",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/mfa", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						// Feature scope
						// Operation
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
				Body data.UpdateProfileMfaRequest
			},
		) (*struct{ Body data.UserProfileMfaResponse }, error) {
			result, errCode, err := controller.UpdateProfileMfa(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserProfileMfaResponse }{Body: *data.FromUserMfa(result)}, nil
		},
	)

	// Delete account
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-profile",
			Summary:     "Delete user account",
			Description: "Delete current user account",
			Method:      http.MethodDelete,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						// Feature scope
						// Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
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
			Description: "Retrieve profile information for the current user",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						// Feature scope
						// Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct{},
		) (*struct{ Body data.UserProfileResponse }, error) {
			result, errCode, err := controller.GetProfile(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserProfileResponse }{Body: *data.FromUser(result)}, nil
		},
	)

	// Get profile from login
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-profile-logged",
			Summary:     "Get profile logged",
			Description: "Retrieve profile information after successful logged for the current user",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/logged", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						// Feature scope
						// Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct{},
		) (*struct{ Body data.UserLoginResponse }, error) {
			result, errCode, err := controller.GetProfileLogged(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UserLoginResponse }{Body: *result}, nil
		},
	)
}
