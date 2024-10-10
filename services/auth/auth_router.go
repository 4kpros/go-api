package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/services/auth/data"
	"github.com/danielgtaylor/huma/v2"
)

func SetupEndpoints(
	humaApi *huma.API,
	controller *AuthController,
) {
	var endpointConfig = struct {
		Group string
		Tag   []string
	}{
		Group: "/auth",
		Tag:   []string{"Authentication"},
	}
	const requireAuth = false

	// Sign in with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-email",
			Summary:       "Login with email",
			Description:   "Login user with email and password. Account need to be activated to retrieve OK response.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
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
			var result, errCode, err = controller.SignInWithEmail(&input.Body, &input.SignInDevice)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignInResponse }{Body: *result}, nil
		},
	)

	// Sign in with phone number
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-phone",
			Summary:       "Login with phone",
			Description:   "Login user with phone number and password. Account need to be activated to retrieve OK response.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/phone", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
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
			var result, errCode, err = controller.SignInWithPhoneNumber(&input.Body, &input.SignInDevice)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignInResponse }{Body: *result}, nil
		},
	)

	// Sign in with provider
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-provider",
			Summary:       "Login with provider",
			Description:   "Login user with provider and and token.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/provider", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
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
			var result, errCode, err = controller.SignInWithProvider(&input.Body, &input.SignInDevice)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignInResponse }{Body: *result}, nil
		},
	)

	// Sign up with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "register-email",
			Summary:       "Register with email",
			Description:   "Register new user with email and password.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/register/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.SignUpWithEmailRequest
			},
		) (*struct{ Body data.SignUpResponse }, error) {
			var result, errCode, err = controller.SignUpWithEmail(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignUpResponse }{Body: *result}, nil
		},
	)

	// Sign up with phone
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "register-phone",
			Summary:       "Register with phone",
			Description:   "Register new user with phone number and password.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/register/phone", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.SignUpWithPhoneNumberRequest
			},
		) (*struct{ Body data.SignUpResponse }, error) {
			var result, errCode, err = controller.SignUpWithPhoneNumber(&input.Body)
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
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ActivateAccountRequest
			},
		) (*struct{ Body data.ActivateAccountResponse }, error) {
			var result, errCode, err = controller.ActivateAccount(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ActivateAccountResponse }{Body: *result}, nil
		},
	)

	// Reset password step 1 with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "reset-init-email",
			Summary:       "Reset step 1 - email",
			Description:   "Reset user password, step 1 with email.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/reset/init/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ResetPasswordWithEmailInitRequest
			},
		) (*struct {
			Body data.ResetPasswordInitResponse
		}, error) {
			var result, errCode, err = controller.ResetPasswordEmailInit(
				&data.ResetPasswordInitRequest{
					Email: input.Body.Email,
				},
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ResetPasswordInitResponse
			}{Body: *result}, nil
		},
	)

	// Reset password step 1 with phone
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "reset-init-phone",
			Summary:       "Reset step 1 - phone",
			Description:   "Reset user password, step 1 with phone number.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/reset/init/phone", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ResetPasswordWithPhoneNumberInitRequest
			},
		) (*struct {
			Body data.ResetPasswordInitResponse
		}, error) {
			var result, errCode, err = controller.ResetPasswordPhoneNumberInit(
				&data.ResetPasswordInitRequest{
					PhoneNumber: input.Body.PhoneNumber,
				},
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ResetPasswordInitResponse
			}{Body: *result}, nil
		},
	)

	// Reset password step 2
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "reset-code",
			Summary:       "Reset step 2",
			Description:   "Reset user password, step 2 need your code received from step 1.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/reset/code", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ResetPasswordCodeRequest
			},
		) (*struct {
			Body data.ResetPasswordCodeResponse
		}, error) {
			var result, errCode, err = controller.ResetPasswordCode(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ResetPasswordCodeResponse
			}{Body: *result}, nil
		},
	)

	// Reset password step 3
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "reset-password",
			Summary:       "Reset step 3",
			Description:   "Reset user password, step 3 to set your new password by providing a token received from step 2.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/reset/password", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ResetPasswordNewPasswordRequest
			},
		) (*struct {
			Body data.ResetPasswordNewPasswordResponse
		}, error) {
			var result, errCode, err = controller.ResetPasswordNewPassword(&input.Body)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ResetPasswordNewPasswordResponse
			}{Body: *result}, nil
		},
	)

	// Logout
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "logout",
			Summary:       "Logout",
			Description:   "Logout user with provided token.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/logout", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.SignOutRequest
			},
		) (*struct{ Body data.SignOutResponse }, error) {
			var result, errCode, err = controller.SignOut(input.Token)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.SignOutResponse }{Body: *result}, nil
		},
	)
}
