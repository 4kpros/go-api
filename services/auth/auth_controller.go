package auth

import "github.com/4kpros/go-api/common/types"

type AuthController struct {
	Service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	// --- VERY IMPORTANT ---
	var _ = &types.ErrorResponse{} // Don't remove this line. It very important for swagger docs generation. It's used to import type.
	// --- VERY IMPORTANT ---

	return &AuthController{Service: service}
}

// func (controller *AuthController) SignInWithEmail(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.SignInRequest{}
// 	var deviceName string = ctx.Header("User-Agent")
// 	ctx.Bind(reqData)
// 	isEmailValid := utils.IsEmailValid(reqData.Email)
// 	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
// 	if !isEmailValid && !isPasswordValid {
// 		message := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
// 		huma.Error400BadRequest(message, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isEmailValid {
// 		message := "Invalid email! Please enter valid information."
// 		huma.Error400BadRequest(message, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isPasswordValid {
// 		message := "Invalid password! Password missing " + missingPasswordChars
// 		huma.Error400BadRequest(message, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Execute the service
// 	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignIn(deviceName, reqData)
// 	if err != nil {
// 		if errCode == http.StatusForbidden || len(validateAccountToken) > 0 {
// 			ctx.JSON(http.StatusForbidden, response.SignUpResponse{
// 				Token:   validateAccountToken,
// 				Message: "Account is not activated! Please activate your account to start using your services.",
// 			})
// 			return
// 		}
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return response
// 	ctx.JSON(http.StatusOK, response.SignInResponse{
// 		AccessToken: accessToken,
// 		Expires:     *accessExpires,
// 	})
// }

// func (controller *AuthController) SignInWithPhoneNumber(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.SignInRequest{}
// 	var deviceName string = ctx.Header("User-Agent")
// 	ctx.Bind(reqData)
// 	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
// 	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
// 	if !isPhoneNumberValid && !isPasswordValid {
// 		message := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
// 		huma.WriteErr(nil, ctx, http.StatusBadRequest, message, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isPhoneNumberValid {
// 		message := "Invalid phone number! Please enter valid information."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isPasswordValid {
// 		message := "Invalid password! Password missing " + missingPasswordChars
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Execute the service
// 	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignIn(deviceName, reqData)
// 	if err != nil {
// 		if errCode == http.StatusForbidden || len(validateAccountToken) > 0 {
// 			// ctx.JSON(http.StatusForbidden, response.SignUpResponse{
// 			// 	Token:   validateAccountToken,
// 			// 	Message: "Account is not activated! Please activate your account to start using your services.",
// 			// })
// 			return
// 		}
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.SignInResponse{
// 		AccessToken: accessToken,
// 		Expires:     *accessExpires,
// 	})
// }

// func (controller *AuthController) SignInWithProvider(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.SignInWithProviderRequest{}
// 	var deviceName string = ctx.Header("User-Agent")
// 	ctx.Bind(reqData)

// 	// Execute the service
// 	accessToken, accessExpires, errCode, err := controller.Service.SignInWithProvider(deviceName, reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.SignInResponse{
// 		AccessToken: accessToken,
// 		Expires:     *accessExpires,
// 	})
// }

// func (controller *AuthController) SignUpWithEmail(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.SignUpRequest{}
// 	ctx.Bind(reqData)
// 	isEmailValid := utils.IsEmailValid(reqData.Email)
// 	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
// 	if !isEmailValid && !isPasswordValid {
// 		message := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isEmailValid {
// 		message := "Invalid email! Please enter valid information."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isPasswordValid {
// 		message := "Invalid password! Password missing " + missingPasswordChars
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Execute the service
// 	token, errCode, err := controller.Service.SignUp(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.SignUpResponse{
// 		Token:   token,
// 		Message: "Account created! Please activate your account to start using your services.",
// 	})
// }

// func (controller *AuthController) SignUpWithPhoneNumber(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.SignUpRequest{}
// 	ctx.Bind(reqData)
// 	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
// 	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
// 	if !isPhoneNumberValid && !isPasswordValid {
// 		message := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isPhoneNumberValid {
// 		message := "Invalid phone number! Please enter valid information."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}
// 	if !isPasswordValid {
// 		message := "Invalid password! Password missing " + missingPasswordChars
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Execute the service
// 	token, errCode, err := controller.Service.SignUp(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.SignUpResponse{
// 		Token:   token,
// 		Message: "Account created! Please activate your account to start using your services.",
// 	})
// }

// func (controller *AuthController) ActivateAccount(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.ActivateAccountRequest{}
// 	ctx.Bind(reqData)

// 	// Execute the service
// 	activatedAt, errCode, err := controller.Service.ActivateAccount(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.ActivateAccountResponse{
// 		ActivatedAt: *activatedAt,
// 	})
// }

// func (controller *AuthController) ResetPasswordEmailInit(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.ResetPasswordInitRequest{}
// 	ctx.Bind(reqData)
// 	isEmailValid := utils.IsEmailValid(reqData.Email)
// 	if !isEmailValid {
// 		message := "Invalid email! Please enter valid information."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Execute the service
// 	token, errCode, err := controller.Service.ResetPasswordInit(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}
// 	if len(token) <= 0 {
// 		errCode = http.StatusInternalServerError
// 		message := "Failed to start the process! Please try again later."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.ResetPasswordInitResponse{
// 		Token: token,
// 	})
// }

// func (controller *AuthController) ResetPasswordPhoneNumberInit(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.ResetPasswordInitRequest{}
// 	ctx.Bind(reqData)
// 	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
// 	if !isPhoneNumberValid {
// 		message := "Invalid phone number! Please enter valid information."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Execute the service
// 	token, errCode, err := controller.Service.ResetPasswordInit(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}
// 	if len(token) <= 0 {
// 		errCode = http.StatusInternalServerError
// 		message := "Failed to start the process! Please try again later."
// 		huma.WriteErr(nil, http.StatusBadRequest, fmt.Errorf("%s", message))
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.ResetPasswordInitResponse{
// 		Token: token,
// 	})
// }

// func (controller *AuthController) ResetPasswordCode(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.ResetPasswordCodeRequest{}
// 	ctx.Bind(reqData)

// 	// Execute the service
// 	token, errCode, err := controller.Service.ResetPasswordCode(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.ResetPasswordCodeResponse{
// 		Token: token,
// 	})
// }

// func (controller *AuthController) ResetPasswordNewPassword(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.ResetPasswordNewPasswordRequest{}
// 	ctx.Bind(reqData)

// 	// Execute the service
// 	errCode, err := controller.Service.ResetPasswordNewPassword(reqData)
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.ResetPasswordNewPasswordResponse{
// 		Message: "Password successful changed! Please sign in to start using our services.",
// 	})
// }

// func (controller *AuthController) SignOut(ctx huma.Context, next func(huma.Context)) {
// 	// Execute the service
// 	errCode, err := controller.Service.SignOut(utils.ExtractBearerTokenHeader(c))
// 	if err != nil {
// 		huma.WriteErr(nil, errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.SignOutResponse{
// 		Message: "Successful signed out! See you soon bye.",
// 	})
// }
