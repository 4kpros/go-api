package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/auth/data"
)

type AuthController struct {
	Service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (controller *AuthController) SignInWithEmail(deviceName string, input *data.SignInWithEmailRequest) (result *data.SignInResponse, errCode int, err error) {
	// Check email and password
	isEmailValid := utils.IsEmailValid(input.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isEmailValid && !isPasswordValid {
		errMessage := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isEmailValid {
		errMessage := "Invalid email! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage := "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

	var accessToken string
	var accessExpires *time.Time = nil
	accessToken, accessExpires, errCode, err = controller.Service.SignIn(
		deviceName,
		&data.SignInRequest{
			Email:         input.Email,
			Password:      input.Password,
			StayConnected: input.StayConnected,
		},
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

func (controller *AuthController) SignInWithPhoneNumber(deviceName string, input *data.SignInWithPhoneNumberRequest) (result *data.SignInResponse, errCode int, err error) {
	// Check phone number and password
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isPhoneNumberValid && !isPasswordValid {
		errMessage := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPhoneNumberValid {
		errMessage := "Invalid phone number! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage := "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

	var accessToken string
	var accessExpires *time.Time = nil
	accessToken, accessExpires, errCode, err = controller.Service.SignIn(
		deviceName,
		&data.SignInRequest{
			PhoneNumber:   input.PhoneNumber,
			Password:      input.Password,
			StayConnected: input.StayConnected,
		},
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

func (controller *AuthController) SignInWithProvider(deviceName string, input *data.SignInWithProviderRequest) (result *data.SignInResponse, errCode int, err error) {
	// Check provider
	isProviderValid := utils.IsAuthProviderValid(input.Provider)
	errCode = http.StatusBadRequest
	if !isProviderValid {
		errMessage := "Invalid or empty provider! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}

	var accessToken string
	var accessExpires *time.Time = nil
	accessToken, accessExpires, errCode, err = controller.Service.SignInWithProvider(deviceName, input)
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
	isEmailValid := utils.IsEmailValid(input.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isEmailValid && !isPasswordValid {
		errMessage := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isEmailValid {
		errMessage := "Invalid email! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage := "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

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
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
	errCode = http.StatusBadRequest
	if !isPhoneNumberValid && !isPasswordValid {
		errMessage := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPhoneNumberValid {
		errMessage := "Invalid phone number! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if !isPasswordValid {
		errMessage := "Invalid password! Password missing " + missingPasswordChars
		err = fmt.Errorf("%s", errMessage)
		return
	}

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
	activatedAt, errCode, err := controller.Service.ActivateAccount(input)
	if err != nil {
		return
	}
	result = &data.ActivateAccountResponse{
		ActivatedAt: *activatedAt,
	}

	return
}

func (controller *AuthController) ResetPasswordEmailInit(input *data.ResetPasswordInitRequest) (result *data.ResetPasswordInitResponse, errCode int, err error) {
	// Check email
	isEmailValid := utils.IsEmailValid(input.Email)
	errCode = http.StatusBadRequest
	if !isEmailValid {
		errMessage := "Invalid email! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}

	var token string
	token, errCode, err = controller.Service.ResetPasswordInit(input)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		errMessage := "Failed to start the process! Please try again later."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	result = &data.ResetPasswordInitResponse{
		Token: token,
	}

	return
}

func (controller *AuthController) ResetPasswordPhoneNumberInit(input *data.ResetPasswordInitRequest) (result *data.ResetPasswordInitResponse, errCode int, err error) {
	// Check phone number
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.PhoneNumber)
	errCode = http.StatusBadRequest
	if !isPhoneNumberValid {
		errMessage := "Invalid phone number! Please enter valid information."
		err = fmt.Errorf("%s", errMessage)
		return
	}

	var token string
	token, errCode, err = controller.Service.ResetPasswordInit(input)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		errMessage := "Failed to start the process! Please try again later."
		err = fmt.Errorf("%s", errMessage)
		return
	}
	result = &data.ResetPasswordInitResponse{
		Token: token,
	}

	return
}

func (controller *AuthController) ResetPasswordCode(input *data.ResetPasswordCodeRequest) (result *data.ResetPasswordCodeResponse, errCode int, err error) {
	var token string
	token, errCode, err = controller.Service.ResetPasswordCode(input)
	if err != nil {
		return
	}
	result = &data.ResetPasswordCodeResponse{
		Token: token,
	}

	return
}

func (controller *AuthController) ResetPasswordNewPassword(input *data.ResetPasswordNewPasswordRequest) (result *data.ResetPasswordNewPasswordResponse, errCode int, err error) {
	errCode, err = controller.Service.ResetPasswordNewPassword(input)
	if err != nil {
		return
	}
	result = &data.ResetPasswordNewPasswordResponse{
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
