package auth

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth/data"
	"github.com/4kpros/go-api/services/user/model"
)

type AuthService struct {
	Repository *AuthRepository
}

func NewAuthService(repository *AuthRepository) *AuthService {
	return &AuthService{Repository: repository}
}

// Login with email or phone number
func (service *AuthService) SignIn(input *data.SignInRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		userFound, errFound = service.Repository.GetByEmail(input.Email)
		errMessage = "Invalid email or password! Please enter valid information."
	} else {
		errMessage = "Invalid phone number or password! Please enter valid information."
		userFound, errFound = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if errFound != nil || userFound == nil || userFound.Email != input.Email {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isPasswordMatches, errCompare = utils.CompareArgon2id(input.Password, userFound.Password)
	if errCompare != nil || !isPasswordMatches {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if account is activated
	if !userFound.IsActivated {
		var randomCode, _ = utils.GenerateRandomCode(5)
		var jwtToken, token, errEncode = utils.EncodeJWTToken(
			&types.JwtToken{
				UserId:   userFound.ID,
				RoleId:   userFound.RoleId,
				Platform: "*",
				Device:   "*",
				App:      "*",
				Code:     randomCode,
			},
			constants.JWT_ISSUER_ACTIVATE,
			utils.NewExpiresDateDefault(),
			config.Keys.JwtPrivateKey,
			config.SetRedisStr,
		)
		if errEncode != nil || jwtToken == nil {
			errCode = http.StatusInternalServerError
			err = constants.HTTP_500_ERROR_MESSAGE("encode JWT token")
			return
		}
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "Account found but not activated! Please activate your account to start using your services.")

		fmt.Printf("\nGenerated code: %d \n", randomCode)
		// Send code to email or phone number
		if utils.IsEmailValid(input.Email) {
			go config.SendMail(
				"Reset password",
				fmt.Sprintf("Your code: %d \nYour token: %s", randomCode, token),
				input.Email,
			)
		} else {
			// TODO send code to phone number
		}

		return
	}

	// Generate new token
	var jwtToken, token, errEncode = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: device.Platform,
			Device:   device.DeviceName,
			App:      device.App,
		},
		constants.JWT_ISSUER_SESSION,
		utils.NewExpiresDateSignIn(input.StayConnected),
		config.Keys.JwtPrivateKey,
		config.InsertRedisArrayStr,
	)
	if errEncode != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("encode JWT token")
	}
	accessToken = token
	accessExpires = &jwtToken.ExpiresAt.Time
	return
}

// Login with provider like Google and Facebook
func (service *AuthService) SignInWithProvider(input *data.SignInWithProviderRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Validate provider token and update user
	var providerUserId = "Test"
	if len(providerUserId) <= 0 {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
		return
	}

	// Save user if it's not in database
	var userFound, errFound = service.Repository.GetByProvider(input.Provider, providerUserId)
	if errFound != nil || userFound == nil || userFound.Provider != input.Provider {
		var user = &model.User{
			Provider:       input.Provider,
			ProviderUserId: providerUserId,
		}
		err = service.Repository.Create(user)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.HTTP_500_ERROR_MESSAGE("create user on database")
			return
		}
	}

	// Generate new token
	var jwtToken, token, errEncode = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: device.Platform,
			Device:   device.DeviceName,
			App:      device.App,
		},
		constants.JWT_ISSUER_SESSION,
		utils.NewExpiresDateSignIn(true),
		config.Keys.JwtPrivateKey,
		config.InsertRedisArrayStr,
	)
	if errEncode != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("encode JWT token")
	}
	accessToken = token
	accessExpires = &jwtToken.ExpiresAt.Time
	return
}

// Register with email or phone number
func (service *AuthService) SignUp(input *data.SignUpRequest) (errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		userFound, errFound = service.Repository.GetByEmail(input.Email)
		errMessage = "user email"
	} else {
		errMessage = "user phone number"
		userFound, errFound = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("find user on database")
		return
	}
	if userFound.Email == input.Email {
		errCode = http.StatusFound
		err = constants.HTTP_302_ERROR_MESSAGE(errMessage)
		return
	}

	// Create new user
	userFound.Email = input.Email
	userFound.PhoneNumber = input.PhoneNumber
	userFound.Password = input.Password
	err = service.Repository.Create(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user on database")
		return
	}

	// Generate new token
	var randomCode, _ = utils.GenerateRandomCode(5)
	var jwtToken, token, errEncode = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JWT_ISSUER_ACTIVATE,
		utils.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisStr,
	)
	if errEncode != nil || jwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate random code")
	}

	fmt.Printf("\nGenerated code: %d \n", randomCode)
	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go config.SendMail(
			"Reset password",
			fmt.Sprintf("Your code: %d \nYour token: %s", randomCode, token),
			input.Email,
		)
	} else {
		// TODO send code to phone number
	}
	return
}

// Activate user account
func (service *AuthService) ActivateAccount(input *data.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string = "Invalid or expired token! Please enter valid information."
	var jwtToken, errDecode = utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecode != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_ACTIVATE {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisStr)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if code is valid
	fmt.Printf("\nRequested code: %d\n", input.Code)
	if jwtToken.Code <= 0 || jwtToken.Code != input.Code {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Check if user exists
	var userFound, errFound = service.Repository.GetById(jwtToken.UserId)
	if errFound != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Check if account is activated
	if userFound.IsActivated {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User account is already activated! Please sign in and start using our services.")
		return
	}

	// Create user info
	var userInfo = &model.UserInfo{}
	err = service.Repository.CreateUserInfo(userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user")
		return
	}

	// Update account
	var tmpActivatedAt = time.Now()
	userFound.ActivatedAt = &tmpActivatedAt
	userFound.IsActivated = true
	userFound.UserInfoId = userInfo.ID
	err = service.Repository.Update(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user")
		return
	}
	activatedAt = userFound.ActivatedAt

	// Invalidate token
	config.DeleteRedisStr(utils.GetJWTCachedKey(jwtToken))

	// Send welcome message
	// TODO
	return
}

// Forgot password step 1: request forgot password
func (service *AuthService) ForgotPasswordInit(input *data.ForgotPasswordInitRequest) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		errMessage = "User with this email"
		userFound, errFound = service.Repository.GetByEmail(input.Email)
	} else {
		errMessage = "User with this phone number"
		userFound, errFound = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if errFound != nil || userFound.ID <= 0 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE(errMessage)
		return
	}

	// Generate new random code
	var randomCode, errRandomCode = utils.GenerateRandomCode(5)
	if errRandomCode != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate random code")
	}
	var expires = utils.NewExpiresDateDefault()
	var newJwtToken, newToken, errEncode = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JWT_ISSUER_FORGOT_PASSWORD_CODE,
		expires,
		config.Keys.JwtPrivateKey,
		config.SetRedisStr,
	)
	if errEncode != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate new JWT token")
		return
	}
	token = newToken

	fmt.Printf("\nGenerated code: %d \n", randomCode)
	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go config.SendMail(
			"Reset password",
			fmt.Sprintf("Your code: %d", randomCode),
			input.Email,
		)
	} else {
		// TODO send code to phone number
	}
	return
}

// Forgot password step 2: validate sended code
func (service *AuthService) ForgotPasswordCode(input *data.ForgotPasswordCodeRequest) (token string, errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string = "Invalid or expired token! Please enter valid information."
	var jwtToken, errDecode = utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecode != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_FORGOT_PASSWORD_CODE {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisStr)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if code is valid
	if jwtToken.Code <= 0 || jwtToken.Code != input.Code {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Check if user exists
	var userFound, errFound = service.Repository.GetById(jwtToken.UserId)
	if errFound != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Invalidate token
	config.DeleteRedisStr(utils.GetJWTCachedKey(jwtToken))

	// Generate new token
	var newJwtToken, newToken, errEncode = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: "*",
			Device:   "*",
			App:      "*",
		},
		constants.JWT_ISSUER_SESSION,
		utils.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisStr,
	)
	if errEncode != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate new JWT token")
	}
	token = newToken
	return
}

// Forgot password step 3: setup new password
func (service *AuthService) ForgotPasswordNewPassword(input *data.ForgotPasswordNewPasswordRequest) (errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string = "Invalid or expired token! Please enter valid information."
	var jwtToken, errDecode = utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecode != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_FORGOT_PASSWORD_NEW_PASSWORD {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisStr)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userFound, errFound = service.Repository.GetById(jwtToken.UserId)
	if errFound != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user password
	var userUpdated, errUpdate = service.Repository.UpdatePasswordById(jwtToken.UserId, input.NewPassword)
	if errUpdate != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update password")
		return
	}

	// Invalidate token
	var _, errDel = config.DeleteRedisStr(utils.GetJWTCachedKey(jwtToken))
	if errDel != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user from database")
		return
	}
	return
}

// Logout user with provided token
func (service *AuthService) SignOut(token string) (errCode int, err error) {
	// Extract token information and validate the token
	var jwtToken, errDecode = utils.DecodeJWTToken(token, config.Keys.JwtPublicKey)
	if errDecode != nil || jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_SESSION {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}
	var isTokenValid = utils.ValidateJWTToken(token, jwtToken, config.CheckRedisStrFromArrayStr(token))
	if !isTokenValid {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}

	// Invalidate token
	var sessions, errSession = config.GetRedisArrayStr(fmt.Sprintf("%d", jwtToken.UserId))
	if errSession != nil {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}
	var tokenIndex = slices.Index(sessions, token)
	if tokenIndex < 0 {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}
	var errDel = config.DeleteRedisStrFromArrayStr(fmt.Sprintf("%d", jwtToken.UserId), int64(tokenIndex))
	if errDel != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("delete cached session")
		return
	}
	return
}
