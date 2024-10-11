package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/auth/data"
)

type AuthController struct {
	Service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (controller *AuthController) SignInWithEmail(input *data.SignInWithEmailRequest, device *data.SignInDevice) (result *data.SignInResponse, errCode int, err error) {
	// Check email and password
	var errMessage string
	var isEmailValid = utils.IsEmailValid(input.Email)
	var isPasswordValid, missingPasswordChars = utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isEmailValid && !isPasswordValid {
		errMessage = "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isEmailValid {
		errMessage = "Invalid email! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage = "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	var accessToken string
	var accessExpires *time.Time = nil
	accessToken, accessExpires, errCode, err = controller.Service.SignIn(
		&data.SignInRequest{
			Email:         input.Email,
			Password:      input.Password,
			StayConnected: input.StayConnected,
		},
		device,
	)
	if err != nil {
		return
	}
	result = &data.SignInResponse{
		AccessToken: accessToken,
		Expires:     *accessExpires,
	}
	return
}

func (controller *AuthController) SignInWithPhoneNumber(input *data.SignInWithPhoneNumberRequest, device *data.SignInDevice) (result *data.SignInResponse, errCode int, err error) {
	// Check phone number and password
	var errMessage string
	var isPhoneNumberValid = utils.IsPhoneNumberValid(input.PhoneNumber)
	var isPasswordValid, missingPasswordChars = utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isPhoneNumberValid && !isPasswordValid {
		errMessage = "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPhoneNumberValid {
		errMessage = "Invalid phone number! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage = "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	var accessToken string
	var accessExpires *time.Time = nil
	accessToken, accessExpires, errCode, err = controller.Service.SignIn(
		&data.SignInRequest{
			PhoneNumber:   input.PhoneNumber,
			Password:      input.Password,
			StayConnected: input.StayConnected,
		},
		device,
	)
	if err != nil {
		return
	}
	result = &data.SignInResponse{
		AccessToken: accessToken,
		Expires:     *accessExpires,
	}
	return
}

func (controller *AuthController) SignInWithProvider(input *data.SignInWithProviderRequest, device *data.SignInDevice) (result *data.SignInResponse, errCode int, err error) {
	// Check provider
	var errMessage string
	var isProviderValid = utils.IsAuthProviderValid(input.Provider)
	errCode = http.StatusBadRequest
	if !isProviderValid {
		errMessage = "Invalid or empty provider! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	var accessToken string
	var accessExpires *time.Time = nil
	accessToken, accessExpires, errCode, err = controller.Service.SignInWithProvider(input, device)
	if err != nil {
		return
	}
	result = &data.SignInResponse{
		AccessToken: accessToken,
		Expires:     *accessExpires,
	}
	return
}

func (controller *AuthController) SignUpWithEmail(input *data.SignUpWithEmailRequest) (result *data.SignUpResponse, errCode int, err error) {
	// Check email and password
	var errMessage string
	var isEmailValid = utils.IsEmailValid(input.Email)
	var isPasswordValid, missingPasswordChars = utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isEmailValid && !isPasswordValid {
		errMessage = "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isEmailValid {
		errMessage = "Invalid email! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage = "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	errCode, err = controller.Service.SignUp(
		&data.SignUpRequest{
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data.SignUpResponse{
		Message: "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *AuthController) SignUpWithPhoneNumber(input *data.SignUpWithPhoneNumberRequest) (result *data.SignUpResponse, errCode int, err error) {
	// Check phone number and password
	var errMessage string
	var isPhoneNumberValid = utils.IsPhoneNumberValid(input.PhoneNumber)
	var isPasswordValid, missingPasswordChars = utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isPhoneNumberValid && !isPasswordValid {
		errMessage = "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPhoneNumberValid {
		errMessage = "Invalid phone number! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage = "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	errCode, err = controller.Service.SignUp(
		&data.SignUpRequest{
			PhoneNumber: input.PhoneNumber,
			Password:    input.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data.SignUpResponse{
		Message: "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *AuthController) ActivateAccount(input *data.ActivateAccountRequest) (result *data.ActivateAccountResponse, errCode int, err error) {
	var activatedAt *time.Time
	activatedAt, errCode, err = controller.Service.ActivateAccount(input)
	if err != nil {
		return
	}
	result = &data.ActivateAccountResponse{
		ActivatedAt: *activatedAt,
	}
	return
}

func (controller *AuthController) ForgotPasswordEmailInit(input *data.ForgotPasswordInitRequest) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	// Check email
	var errMessage string
	var isEmailValid = utils.IsEmailValid(input.Email)
	errCode = http.StatusBadRequest
	if !isEmailValid {
		errMessage = "Invalid email! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	var token string
	token, errCode, err = controller.Service.ForgotPasswordInit(input)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		errMessage = "Failed to start the process! Please try again later."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	result = &data.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordPhoneNumberInit(input *data.ForgotPasswordInitRequest) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	// Check phone number
	var errMessage string
	var isPhoneNumberValid = utils.IsPhoneNumberValid(input.PhoneNumber)
	errCode = http.StatusBadRequest
	if !isPhoneNumberValid {
		errMessage = "Invalid phone number! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Execute the service
	var token string
	token, errCode, err = controller.Service.ForgotPasswordInit(input)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		errMessage = "Failed to start the process! Please try again later."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	result = &data.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordCode(input *data.ForgotPasswordCodeRequest) (result *data.ForgotPasswordCodeResponse, errCode int, err error) {
	var token string
	token, errCode, err = controller.Service.ForgotPasswordCode(input)
	if err != nil {
		return
	}
	result = &data.ForgotPasswordCodeResponse{
		Token: token,
	}
	return
}

func (controller *AuthController) ForgotPasswordNewPassword(input *data.ForgotPasswordNewPasswordRequest) (result *data.ForgotPasswordNewPasswordResponse, errCode int, err error) {
	errCode, err = controller.Service.ForgotPasswordNewPassword(input)
	if err != nil {
		return
	}
	result = &data.ForgotPasswordNewPasswordResponse{
		Message: "Password successful changed! Please sign in to start using our services.",
	}
	return
}

func (controller *AuthController) SignOut(token string) (result *data.SignOutResponse, errCode int, err error) {
	errCode, err = controller.Service.SignOut(token)
	if err != nil {
		return
	}
	result = &data.SignOutResponse{
		Message: "Successful signed out! See you soon bye.",
	}
	return
}
