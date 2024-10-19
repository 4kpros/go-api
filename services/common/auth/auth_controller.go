package auth

import (
	data2 "api/services/common/auth/data"
	"context"
	"fmt"
	"net/http"
	"time"

	"api/common/helpers"
	"api/common/utils"
)

type AuthController struct {
	Service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (controller *AuthController) SignInWithEmail(
	ctx *context.Context,
	input *struct {
		data2.SignInDevice
		Body data2.SignInWithEmailRequest
	},
) (result *data2.SignInResponse, errCode int, err error) {
	// Check input
	isEmailValid := utils.IsEmailValid(input.Body.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Body.Password)
	if !isEmailValid && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid email and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if !isEmailValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid email! Please enter valid information.")
		return
	}
	if !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid password! Password missing",
			missingPasswordChars,
		)
		return
	}

	// Execute the service
	accessToken, accessExpires, activateAccountToken, errCode, err := controller.Service.SignIn(
		&data2.SignInRequest{
			Email:         input.Body.Email,
			Password:      input.Body.Password,
			StayConnected: input.Body.StayConnected,
		},
		&input.SignInDevice,
	)
	if err != nil {
		return
	}
	result = &data2.SignInResponse{
		AccessToken:          accessToken,
		Expires:              accessExpires,
		ActivateAccountToken: activateAccountToken,
	}
	return
}

func (controller *AuthController) SignInWithPhoneNumber(
	ctx *context.Context,
	input *struct {
		data2.SignInDevice
		Body data2.SignInWithPhoneNumberRequest
	},
) (result *data2.SignInResponse, errCode int, err error) {
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.Body.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Body.Password)
	if !isPhoneNumberValid && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid phone number and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid phone number! Please enter valid information.")
		return
	}
	if !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid password! Password missing",
			missingPasswordChars,
		)
		return
	}

	// Execute the service
	accessToken, accessExpires, activateAccountToken, errCode, err := controller.Service.SignIn(
		&data2.SignInRequest{
			PhoneNumber:   input.Body.PhoneNumber,
			Password:      input.Body.Password,
			StayConnected: input.Body.StayConnected,
		},
		&input.SignInDevice,
	)
	if err != nil {
		return
	}
	result = &data2.SignInResponse{
		AccessToken:          accessToken,
		Expires:              accessExpires,
		ActivateAccountToken: activateAccountToken,
	}
	return
}

func (controller *AuthController) SignInWithProvider(
	ctx *context.Context,
	input *struct {
		data2.SignInDevice
		Body data2.SignInWithProviderRequest
	},
) (result *data2.SignInResponse, errCode int, err error) {
	// Check input
	isProviderValid := utils.IsAuthProviderValid(input.Body.Provider)
	if !isProviderValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid or empty provider! Please enter valid information.")
		return
	}

	// Execute the service
	accessToken := ""
	var accessExpires *time.Time
	accessToken, accessExpires, errCode, err = controller.Service.SignInWithProvider(&input.Body, &input.SignInDevice)
	if err != nil {
		return
	}
	result = &data2.SignInResponse{
		AccessToken: accessToken,
		Expires:     accessExpires,
	}
	return
}

func (controller *AuthController) SignUpWithEmail(
	ctx *context.Context,
	input *struct {
		Body data2.SignUpWithEmailRequest
	},
) (result *data2.SignUpResponse, errCode int, err error) {
	// Check input
	isEmailValid := utils.IsEmailValid(input.Body.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Body.Password)
	if !isEmailValid && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid email and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if !isEmailValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid email! Please enter valid information.")
		return
	}
	if !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid password! Password missing",
			missingPasswordChars,
		)
		return
	}

	// Execute the service
	var activateAccountToken string
	activateAccountToken, errCode, err = controller.Service.SignUp(
		&data2.SignUpRequest{
			Email:    input.Body.Email,
			Password: input.Body.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data2.SignUpResponse{
		ActivateAccountToken: activateAccountToken,
		Message:              "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *AuthController) SignUpWithPhoneNumber(
	ctx *context.Context,
	input *struct {
		Body data2.SignUpWithPhoneNumberRequest
	},
) (result *data2.SignUpResponse, errCode int, err error) {
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.Body.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Body.Password)
	if !isPhoneNumberValid && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid phone number and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid phone number! Please enter valid information.")
		return
	}
	if !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid password! Password missing",
			missingPasswordChars,
		)
		return
	}

	// Execute the service
	var activateAccountToken string
	activateAccountToken, errCode, err = controller.Service.SignUp(
		&data2.SignUpRequest{
			PhoneNumber: input.Body.PhoneNumber,
			Password:    input.Body.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data2.SignUpResponse{
		ActivateAccountToken: activateAccountToken,
		Message:              "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *AuthController) ActivateAccount(
	ctx *context.Context,
	input *struct {
		Body data2.ActivateAccountRequest
	},
) (result *data2.ActivateAccountResponse, errCode int, err error) {
	activatedAt, errCode, err := controller.Service.ActivateAccount(&input.Body)
	if err != nil {
		return
	}
	result = &data2.ActivateAccountResponse{
		ActivatedAt: activatedAt,
	}
	return
}

func (controller *AuthController) ForgotPasswordEmailInit(
	ctx *context.Context,
	input *struct {
		Body data2.ForgotPasswordWithEmailInitRequest
	},
) (result *data2.ForgotPasswordInitResponse, errCode int, err error) {
	token, errCode, err := controller.Service.ForgotPasswordInit(
		&data2.ForgotPasswordInitRequest{
			Email: input.Body.Email,
		},
	)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		err = fmt.Errorf("%s", "Failed to start the process! Please try again later.")
		return
	}
	result = &data2.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordPhoneNumberInit(
	ctx *context.Context,
	input *struct {
		Body data2.ForgotPasswordWithPhoneNumberInitRequest
	},
) (result *data2.ForgotPasswordInitResponse, errCode int, err error) {
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.Body.PhoneNumber)
	if !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid phone number! Please enter valid information.")
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.ForgotPasswordInit(
		&data2.ForgotPasswordInitRequest{
			PhoneNumber: input.Body.PhoneNumber,
		},
	)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		err = fmt.Errorf("%s", "Failed to start the process! Please try again later.")
		return
	}
	result = &data2.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordCode(
	ctx *context.Context,
	input *struct {
		Body data2.ForgotPasswordCodeRequest
	},
) (result *data2.ForgotPasswordCodeResponse, errCode int, err error) {
	token, errCode, err := controller.Service.ForgotPasswordCode(&input.Body)
	if err != nil {
		return
	}
	result = &data2.ForgotPasswordCodeResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordNewPassword(
	ctx *context.Context,
	input *struct {
		Body data2.ForgotPasswordNewPasswordRequest
	},
) (result *data2.ForgotPasswordNewPasswordResponse, errCode int, err error) {
	errCode, err = controller.Service.ForgotPasswordNewPassword(&input.Body)
	if err != nil {
		return
	}
	result = &data2.ForgotPasswordNewPasswordResponse{
		Message: "Password successful changed! Please sign in to start using our services.",
	}
	return
}

func (controller *AuthController) SignOut(
	ctx *context.Context,
	input *struct {
		data2.SignOutRequest
	},
) (result *data2.SignOutResponse, errCode int, err error) {
	errCode, err = controller.Service.SignOut(helpers.GetJwtContext(ctx), helpers.GetBearerContext(ctx))
	if err != nil {
		return
	}
	result = &data2.SignOutResponse{
		Message: "Successful signed out! See you soon bye.",
	}
	return
}
