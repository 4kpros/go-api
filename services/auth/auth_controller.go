package auth

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/auth/data/request"
	"github.com/4kpros/go-api/services/auth/data/response"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	// --- VERY IMPORTANT ---
	var _ = &types.ErrorResponse{} // Don't remove this line. It very important for swagger docs generation. It's used to import type.
	// --- VERY IMPORTANT ---

	return &AuthController{Service: service}
}

// @Tags Sign in
// @Summary Sign in user with email
// @Accept  json
// @Produce  json
// @Param   payload body request.SignInEmailRequest true "Enter your information"
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 403 {object} types.ErrorResponse "Account is not activated!"
// @Failure 404 {object} types.ErrorResponse "Invalid email or password!"
// @Security ApiKey
// @Router /auth/signin-email [post]
func (controller *AuthController) SignInWithEmail(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignInRequest{}
	var deviceName string = c.GetHeader("User-Agent")
	c.Bind(reqData)
	isEmailValid := utils.IsEmailValid(reqData.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
	if !isEmailValid && !isPasswordValid {
		message := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isEmailValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignIn(deviceName, reqData)
	if err != nil {
		if errCode == http.StatusForbidden || len(validateAccountToken) > 0 {
			c.JSON(http.StatusForbidden, response.SignUpResponse{
				Token:   validateAccountToken,
				Message: "Account is not activated! Please activate your account to start using your services.",
			})
			return
		}
		c.AbortWithError(errCode, err)
		return
	}

	// Return response
	c.JSON(http.StatusOK, response.SignInResponse{
		AccessToken: accessToken,
		Expires:     *accessExpires,
	})
}

// @Tags Sign in
// @Summary Sign in user with phone number
// @Accept  json
// @Produce  json
// @Param   payload body request.SignInPhoneNumberRequest true "Enter your information"
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 403 {object} types.ErrorResponse "Account is not activated!"
// @Failure 404 {object} types.ErrorResponse "Invalid phone number or password!"
// @Security ApiKey
// @Router /auth/signin-phone [post]
func (controller *AuthController) SignInWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignInRequest{}
	var deviceName string = c.GetHeader("User-Agent")
	c.Bind(reqData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
	if !isPhoneNumberValid && !isPasswordValid {
		message := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPhoneNumberValid {
		message := "Invalid phone number! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignIn(deviceName, reqData)
	if err != nil {
		if errCode == http.StatusForbidden || len(validateAccountToken) > 0 {
			c.JSON(http.StatusForbidden, response.SignUpResponse{
				Token:   validateAccountToken,
				Message: "Account is not activated! Please activate your account to start using your services.",
			})
			return
		}
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.SignInResponse{
		AccessToken: accessToken,
		Expires:     *accessExpires,
	})
}

// @Tags Sign in
// @Summary Sign in user with provider
// @Accept  json
// @Produce  json
// @Param   payload body request.SignInWithProviderRequest true "Enter your information"
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid inputs!"
// @Failure 404 {object} types.ErrorResponse "Invalid token!"
// @Security ApiKey
// @Router /auth/signin-provider [post]
func (controller *AuthController) SignInWithProvider(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignInWithProviderRequest{}
	var deviceName string = c.GetHeader("User-Agent")
	c.Bind(reqData)

	// Execute the service
	accessToken, accessExpires, errCode, err := controller.Service.SignInWithProvider(deviceName, reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.SignInResponse{
		AccessToken: accessToken,
		Expires:     *accessExpires,
	})
}

// @Tags Sign up
// @Summary Sign up user with email
// @Accept  json
// @Produce  json
// @Param   payload body request.SignUpEmailRequest true "Enter your information"
// @Success 200 {object} response.SignUpResponse "OK"
// @Failure 302 {object} types.ErrorResponse "User with this email already exists!"
// @Security ApiKey
// @Router /auth/signup-email [post]
func (controller *AuthController) SignUpWithEmail(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignUpRequest{}
	c.Bind(reqData)
	isEmailValid := utils.IsEmailValid(reqData.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
	if !isEmailValid && !isPasswordValid {
		message := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isEmailValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.SignUp(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.SignUpResponse{
		Token:   token,
		Message: "Account created! Please activate your account to start using your services.",
	})
}

// @Tags Sign up
// @Summary Sign up user with phone number
// @Accept  json
// @Produce  json
// @Param   payload body request.SignUpPhoneNumberRequest true "Enter your information"
// @Success 200 {object} response.SignUpResponse "OK"
// @Failure 302 {object} types.ErrorResponse "User with this phone number already exists!"
// @Security ApiKey
// @Router /auth/signup-phone [post]
func (controller *AuthController) SignUpWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignUpRequest{}
	c.Bind(reqData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
	if !isPhoneNumberValid && !isPasswordValid {
		message := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPhoneNumberValid {
		message := "Invalid phone number! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.SignUp(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.SignUpResponse{
		Token:   token,
		Message: "Account created! Please activate your account to start using your services.",
	})
}

// @Tags Activate account
// @Summary Activate new user account
// @Accept  json
// @Produce  json
// @Param   payload body request.ActivateAccountRequest true "Enter your information"
// @Success 200 {object} response.ActivateAccountResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid token!"
// @Failure 403 {object} types.ErrorResponse "User account is already activated!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey
// @Router /auth/activate [post]
func (controller *AuthController) ActivateAccount(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ActivateAccountRequest{}
	c.Bind(reqData)

	// Execute the service
	activatedAt, errCode, err := controller.Service.ActivateAccount(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.ActivateAccountResponse{
		ActivatedAt: *activatedAt,
	})
}

// @Tags Reset password
// @Summary Reset password  with email init
// @Accept  json
// @Produce  json
// @Param   payload body request.ResetPasswordEmailInitRequest true "Enter your information"
// @Success 200 {object} response.ResetPasswordInitResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid email input!"
// @Failure 404 {object} types.ErrorResponse "User with email not found!"
// @Security ApiKey
// @Router /auth/reset/init-email [post]
func (controller *AuthController) ResetPasswordEmailInit(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordInitRequest{}
	c.Bind(reqData)
	isEmailValid := utils.IsEmailValid(reqData.Email)
	if !isEmailValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.ResetPasswordInit(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		message := "Failed to start the process! Please try again later."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.ResetPasswordInitResponse{
		Token: token,
	})
}

// @Tags Reset password
// @Summary Reset password with phone number init
// @Accept  json
// @Produce  json
// @Param   payload body request.ResetPasswordPhoneNumberInitRequest true "Enter your information"
// @Success 200 {object} response.ResetPasswordInitResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid phone number input!"
// @Failure 404 {object} types.ErrorResponse "User with phone number not found!"
// @Security ApiKey
// @Router /auth/reset/init-phone [post]
func (controller *AuthController) ResetPasswordPhoneNumberInit(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordInitRequest{}
	c.Bind(reqData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
	if !isPhoneNumberValid {
		message := "Invalid phone number! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.ResetPasswordInit(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		message := "Failed to start the process! Please try again later."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.ResetPasswordInitResponse{
		Token: token,
	})
}

// @Tags Reset password
// @Summary Reset password set code
// @Accept  json
// @Produce  json
// @Param   payload body request.ResetPasswordCodeRequest true "Enter your information"
// @Success 200 {object} response.ResetPasswordInitResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid token!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey
// @Router /auth/reset/code [post]
func (controller *AuthController) ResetPasswordCode(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordCodeRequest{}
	c.Bind(reqData)

	// Execute the service
	token, errCode, err := controller.Service.ResetPasswordCode(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.ResetPasswordCodeResponse{
		Token: token,
	})
}

// @Tags Reset password
// @Summary Reset password set new password
// @Accept  json
// @Produce  json
// @Param   payload body request.ResetPasswordNewPasswordRequest true "Enter your information"
// @Success 200 {object} response.ResetPasswordNewPasswordResponse "OK"
// @Failure 400 {object} types.ErrorResponse "Invalid token or password input!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey
// @Router /auth/reset/new-password [post]
func (controller *AuthController) ResetPasswordNewPassword(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordNewPasswordRequest{}
	c.Bind(reqData)

	// Execute the service
	errCode, err := controller.Service.ResetPasswordNewPassword(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.ResetPasswordNewPasswordResponse{
		Message: "Password successful changed! Please sign in to start using our services.",
	})
}

// @Tags SignOut
// @Summary Sign out user
// @Accept  json
// @Produce  json
// @Success 200 {object} response.SignOutResponse "OK"
// @Failure 403 {object} types.ErrorResponse "Invalid or expired token!"
// @Failure 404 {object} types.ErrorResponse "User not found!"
// @Security ApiKey && Bearer
// @Router /auth/signout [post]
func (controller *AuthController) SignOut(c *gin.Context) {
	// Execute the service
	errCode, err := controller.Service.SignOut(utils.ExtractBearerTokenHeader(c))
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, response.SignOutResponse{
		Message: "Successful signed out! See you soon bye.",
	})
}
