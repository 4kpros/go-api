package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/auth/data"
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
		data.SignInDevice
		Body data.SignInWithEmailRequest
	},
) (result *data.SignInResponse, errCode int, err error) {
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
		&data.SignInRequest{
			Email:         input.Body.Email,
			Password:      input.Body.Password,
			StayConnected: input.Body.StayConnected,
		},
		&input.SignInDevice,
	)
	if err != nil {
		return
	}
	result = &data.SignInResponse{
		AccessToken:          accessToken,
		Expires:              accessExpires,
		ActivateAccountToken: activateAccountToken,
	}
	return
}

func (controller *AuthController) SignInWithPhoneNumber(
	ctx *context.Context,
	input *struct {
		data.SignInDevice
		Body data.SignInWithPhoneNumberRequest
	},
) (result *data.SignInResponse, errCode int, err error) {
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
		&data.SignInRequest{
			PhoneNumber:   input.Body.PhoneNumber,
			Password:      input.Body.Password,
			StayConnected: input.Body.StayConnected,
		},
		&input.SignInDevice,
	)
	if err != nil {
		return
	}
	result = &data.SignInResponse{
		AccessToken:          accessToken,
		Expires:              accessExpires,
		ActivateAccountToken: activateAccountToken,
	}
	return
}

func (controller *AuthController) SignInWithProvider(
	ctx *context.Context,
	input *struct {
		data.SignInDevice
		Body data.SignInWithProviderRequest
	},
) (result *data.SignInResponse, errCode int, err error) {
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
	result = &data.SignInResponse{
		AccessToken: accessToken,
		Expires:     accessExpires,
	}
	return
}

func (controller *AuthController) SignUpWithEmail(
	ctx *context.Context,
	input *struct {
		Body data.SignUpWithEmailRequest
	},
) (result *data.SignUpResponse, errCode int, err error) {
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
		&data.SignUpRequest{
			Email:    input.Body.Email,
			Password: input.Body.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data.SignUpResponse{
		ActivateAccountToken: activateAccountToken,
		Message:              "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *AuthController) SignUpWithPhoneNumber(
	ctx *context.Context,
	input *struct {
		Body data.SignUpWithPhoneNumberRequest
	},
) (result *data.SignUpResponse, errCode int, err error) {
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
		&data.SignUpRequest{
			PhoneNumber: input.Body.PhoneNumber,
			Password:    input.Body.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data.SignUpResponse{
		ActivateAccountToken: activateAccountToken,
		Message:              "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *AuthController) ActivateAccount(
	ctx *context.Context,
	input *struct {
		Body data.ActivateAccountRequest
	},
) (result *data.ActivateAccountResponse, errCode int, err error) {
	activatedAt, errCode, err := controller.Service.ActivateAccount(&input.Body)
	if err != nil {
		return
	}
	result = &data.ActivateAccountResponse{
		ActivatedAt: activatedAt,
	}
	return
}

func (controller *AuthController) ForgotPasswordEmailInit(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordWithEmailInitRequest
	},
) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	token, errCode, err := controller.Service.ForgotPasswordInit(
		&data.ForgotPasswordInitRequest{
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
	result = &data.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordPhoneNumberInit(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordWithPhoneNumberInitRequest
	},
) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.Body.PhoneNumber)
	if !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid phone number! Please enter valid information.")
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.ForgotPasswordInit(
		&data.ForgotPasswordInitRequest{
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
	result = &data.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordCode(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordCodeRequest
	},
) (result *data.ForgotPasswordCodeResponse, errCode int, err error) {
	token, errCode, err := controller.Service.ForgotPasswordCode(&input.Body)
	if err != nil {
		return
	}
	result = &data.ForgotPasswordCodeResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordNewPassword(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordNewPasswordRequest
	},
) (result *data.ForgotPasswordNewPasswordResponse, errCode int, err error) {
	errCode, err = controller.Service.ForgotPasswordNewPassword(&input.Body)
	if err != nil {
		return
	}
	result = &data.ForgotPasswordNewPasswordResponse{
		Message: "Password successful changed! Please sign in to start using our services.",
	}
	return
}

func (controller *AuthController) SignOut(
	ctx *context.Context,
	input *struct {
		data.SignOutRequest
	},
) (result *data.SignOutResponse, errCode int, err error) {
	errCode, err = controller.Service.SignOut(helpers.GetJwtContext(ctx), helpers.GetBearerContext(ctx))
	if err != nil {
		return
	}
	result = &data.SignOutResponse{
		Message: "Successful signed out! See you soon bye.",
	}
	return
}
