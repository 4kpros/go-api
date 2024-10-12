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
	// Check input
	isEmailValid := utils.IsEmailValid(input.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
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
	accessToken, accessExpires, errCode, err := controller.Service.SignIn(
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
		Expires:     accessExpires,
	}
	return
}

func (controller *AuthController) SignInWithPhoneNumber(input *data.SignInWithPhoneNumberRequest, device *data.SignInDevice) (result *data.SignInResponse, errCode int, err error) {
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
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
	accessToken, accessExpires, errCode, err := controller.Service.SignIn(
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
		Expires:     accessExpires,
	}
	return
}

func (controller *AuthController) SignInWithProvider(input *data.SignInWithProviderRequest, device *data.SignInDevice) (result *data.SignInResponse, errCode int, err error) {
	// Check input
	isProviderValid := utils.IsAuthProviderValid(input.Provider)
	if !isProviderValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid or empty provider! Please enter valid information.")
		return
	}

	// Execute the service
	accessToken := ""
	var accessExpires *time.Time
	accessToken, accessExpires, errCode, err = controller.Service.SignInWithProvider(input, device)
	if err != nil {
		return
	}
	result = &data.SignInResponse{
		AccessToken: accessToken,
		Expires:     accessExpires,
	}
	return
}

func (controller *AuthController) SignUpWithEmail(input *data.SignUpWithEmailRequest) (result *data.SignUpResponse, errCode int, err error) {
	// Check input
	isEmailValid := utils.IsEmailValid(input.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
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
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.Password)
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
		ActivatedAt: activatedAt,
	}
	return
}

func (controller *AuthController) ForgotPasswordEmailInit(input *data.ForgotPasswordInitRequest) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	// Check input
	isEmailValid := utils.IsEmailValid(input.Email)
	if !isEmailValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid email! Please enter valid information.")
		return
	}

	// Execute the service
	token := ""
	token, errCode, err = controller.Service.ForgotPasswordInit(input)
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

func (controller *AuthController) ForgotPasswordPhoneNumberInit(input *data.ForgotPasswordInitRequest) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(input.PhoneNumber)
	if !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid phone number! Please enter valid information.")
		return
	}

	// Execute the service
	token := ""
	token, errCode, err = controller.Service.ForgotPasswordInit(input)
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

func (controller *AuthController) ForgotPasswordCode(input *data.ForgotPasswordCodeRequest) (result *data.ForgotPasswordCodeResponse, errCode int, err error) {
	// Check input
	if len(input.Token) <= 0 && input.Code < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token and code! Please enter valid information.")
		return
	}
	if len(input.Token) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if input.Code < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Execute the service
	token := ""
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
	// Check input
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.NewPassword)
	if len(input.Token) <= 0 && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid token and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if len(input.Token) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
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
