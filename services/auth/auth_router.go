package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/auth/data"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *AuthController,
) {
	endpointConfig := types.APIEndpointConfig{
		Group: "/auth",
		Tag:   []string{"Authentication"},
	}

	// Login with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-email",
			Summary:       "Login with email",
			Description:   "Login user with email and password. Account need to be activated to retrieve OK response.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.SignInDevice
				Body data.SignInWithEmailRequest
			},
		) (*struct{ Body data.SignInResponse }, error) {
			result, errCode, err := controller.SignInWithEmail(&input.Body, &input.SignInDevice)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignInResponse }{Body: *result}, nil
		},
	)

	// Login with phone number
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-phone",
			Summary:       "Login with phone",
			Description:   "Login user with phone number and password. Account need to be activated to retrieve OK response.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/phone", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.SignInDevice
				Body data.SignInWithPhoneNumberRequest
			},
		) (*struct{ Body data.SignInResponse }, error) {
			result, errCode, err := controller.SignInWithPhoneNumber(&input.Body, &input.SignInDevice)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignInResponse }{Body: *result}, nil
		},
	)

	// Login with provider
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-provider",
			Summary:       "Login with provider",
			Description:   "Login user with a provider(Google, Facebook, ...) and token.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/provider", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.SignInDevice
				Body data.SignInWithProviderRequest
			},
		) (*struct{ Body data.SignInResponse }, error) {
			result, errCode, err := controller.SignInWithProvider(&input.Body, &input.SignInDevice)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignInResponse }{Body: *result}, nil
		},
	)

	// Register with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "register-email",
			Summary:       "Register with email",
			Description:   "Register new user with email and password.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/register/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.SignUpWithEmailRequest
			},
		) (*struct{ Body data.SignUpResponse }, error) {
			result, errCode, err := controller.SignUpWithEmail(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignUpResponse }{Body: *result}, nil
		},
	)

	// Register with phone
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "register-phone",
			Summary:       "Register with phone",
			Description:   "Register new user with phone number and password.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/register/phone", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.SignUpWithPhoneNumberRequest
			},
		) (*struct{ Body data.SignUpResponse }, error) {
			result, errCode, err := controller.SignUpWithPhoneNumber(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignUpResponse }{Body: *result}, nil
		},
	)

	// Activate account
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "activate",
			Summary:       "Activate",
			Description:   "Activate user account.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/activate", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ActivateAccountRequest
			},
		) (*struct{ Body data.ActivateAccountResponse }, error) {
			result, errCode, err := controller.ActivateAccount(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ActivateAccountResponse }{Body: *result}, nil
		},
	)

	// Forgot password step 1 with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-init-email",
			Summary:       "Forgot step 1 - email",
			Description:   "Forgot password step 1 initialize request with email.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/init/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordWithEmailInitRequest
			},
		) (*struct {
			Body data.ForgotPasswordInitResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordEmailInit(
				&data.ForgotPasswordInitRequest{
					Email: input.Body.Email,
				},
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordInitResponse
			}{Body: *result}, nil
		},
	)

	// Forgot password step 1 with phone
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-init-phone",
			Summary:       "Forgot step 1 - phone",
			Description:   "Forgot password step 1 initialize request with phone number.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/init/phone", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordWithPhoneNumberInitRequest
			},
		) (*struct {
			Body data.ForgotPasswordInitResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordPhoneNumberInit(
				&data.ForgotPasswordInitRequest{
					PhoneNumber: input.Body.PhoneNumber,
				},
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordInitResponse
			}{Body: *result}, nil
		},
	)

	// Forgot password step 2
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-code",
			Summary:       "Forgot step 2",
			Description:   "Forgot password step 2 validate your request with your received(email/phone) code and token from step 1.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/code", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordCodeRequest
			},
		) (*struct {
			Body data.ForgotPasswordCodeResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordCode(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordCodeResponse
			}{Body: *result}, nil
		},
	)

	// Forgot password step 3
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-new-password",
			Summary:       "Forgot step 3",
			Description:   "Forgot password step 3 set your new password by providing a token received from step 2.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/password", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordNewPasswordRequest
			},
		) (*struct {
			Body data.ForgotPasswordNewPasswordResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordNewPassword(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordNewPasswordResponse
			}{Body: *result}, nil
		},
	)

	// Logout
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "logout",
			Summary:     "Logout",
			Description: "Logout user with provided token.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/logout", endpointConfig.Group),
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
				data.SignOutRequest
			},
		) (*struct{ Body data.SignOutResponse }, error) {
			result, errCode, err := controller.SignOut(input.Token)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignOutResponse }{Body: *result}, nil
		},
	)
}
