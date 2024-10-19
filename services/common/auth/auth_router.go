package auth

import (
	data2 "api/services/common/auth/data"
	"context"
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
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
				data2.SignInDevice
				Body data2.SignInWithEmailRequest
			},
		) (*struct{ Body data2.SignInResponse }, error) {
			result, errCode, err := controller.SignInWithEmail(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.SignInResponse }{Body: *result}, nil
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
				data2.SignInDevice
				Body data2.SignInWithPhoneNumberRequest
			},
		) (*struct{ Body data2.SignInResponse }, error) {
			result, errCode, err := controller.SignInWithPhoneNumber(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.SignInResponse }{Body: *result}, nil
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
				data2.SignInDevice
				Body data2.SignInWithProviderRequest
			},
		) (*struct{ Body data2.SignInResponse }, error) {
			result, errCode, err := controller.SignInWithProvider(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.SignInResponse }{Body: *result}, nil
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
				Body data2.SignUpWithEmailRequest
			},
		) (*struct{ Body data2.SignUpResponse }, error) {
			result, errCode, err := controller.SignUpWithEmail(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.SignUpResponse }{Body: *result}, nil
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
				Body data2.SignUpWithPhoneNumberRequest
			},
		) (*struct{ Body data2.SignUpResponse }, error) {
			result, errCode, err := controller.SignUpWithPhoneNumber(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.SignUpResponse }{Body: *result}, nil
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
				Body data2.ActivateAccountRequest
			},
		) (*struct{ Body data2.ActivateAccountResponse }, error) {
			result, errCode, err := controller.ActivateAccount(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.ActivateAccountResponse }{Body: *result}, nil
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
				Body data2.ForgotPasswordWithEmailInitRequest
			},
		) (*struct {
			Body data2.ForgotPasswordInitResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordEmailInit(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data2.ForgotPasswordInitResponse
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
				Body data2.ForgotPasswordWithPhoneNumberInitRequest
			},
		) (*struct {
			Body data2.ForgotPasswordInitResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordPhoneNumberInit(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data2.ForgotPasswordInitResponse
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
				Body data2.ForgotPasswordCodeRequest
			},
		) (*struct {
			Body data2.ForgotPasswordCodeResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordCode(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data2.ForgotPasswordCodeResponse
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
				Body data2.ForgotPasswordNewPasswordRequest
			},
		) (*struct {
			Body data2.ForgotPasswordNewPasswordResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordNewPassword(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data2.ForgotPasswordNewPasswordResponse
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
				data2.SignOutRequest
			},
		) (*struct{ Body data2.SignOutResponse }, error) {
			result, errCode, err := controller.SignOut(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data2.SignOutResponse }{Body: *result}, nil
		},
	)
}
