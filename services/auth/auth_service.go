package auth

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth/data"
	"github.com/4kpros/go-api/services/user/model"
)

type AuthService interface {
	SignIn(input *data.SignInRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error)
	SignInWithProvider(input *data.SignInWithProviderRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error)

	SignUp(input *data.SignUpRequest) (errCode int, err error)

	ActivateAccount(input *data.ActivateAccountRequest) (date *time.Time, errCode int, err error)

	ResetPasswordInit(input *data.ResetPasswordInitRequest) (token string, errCode int, err error)
	ResetPasswordCode(input *data.ResetPasswordCodeRequest) (token string, errCode int, err error)
	ResetPasswordNewPassword(input *data.ResetPasswordNewPasswordRequest) (errCode int, err error)

	SignOut(token string) (errCode int, err error)
}

type AuthServiceImpl struct {
	Repository AuthRepository
}

func NewAuthServiceImpl(repository AuthRepository) AuthService {
	return &AuthServiceImpl{Repository: repository}
}

func (service *AuthServiceImpl) SignIn(input *data.SignInRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error) {
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
			utils.JWT_ISSUER_ACTIVATE,
			utils.NewExpiresDateDefault(),
			config.Keys.JwtPrivateKey,
			config.SetRedisStr,
		)
		if errEncode != nil || jwtToken == nil {
			errCode = http.StatusInternalServerError
			err = errEncode
			return
		}
		errMessage = "Account found but not activated! Please activate your account to start using your services."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)

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
		utils.JWT_ISSUER_SESSION,
		utils.NewExpiresDateSignIn(input.StayConnected),
		config.Keys.JwtPrivateKey,
		config.InsertRedisArrayStr,
	)
	if errEncode != nil {
		errCode = http.StatusInternalServerError
		err = errEncode
	}
	accessToken = token
	accessExpires = &jwtToken.ExpiresAt.Time
	return
}

func (service *AuthServiceImpl) SignInWithProvider(input *data.SignInWithProviderRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Validate provider token and update user
	var errMessage string
	var providerUserId = "Test"
	if len(providerUserId) <= 0 {
		errMessage = "Invalid provider or token! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
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
		utils.JWT_ISSUER_SESSION,
		utils.NewExpiresDateSignIn(true),
		config.Keys.JwtPrivateKey,
		config.InsertRedisArrayStr,
	)
	if errEncode != nil {
		errCode = http.StatusInternalServerError
		err = errEncode
	}
	accessToken = token
	accessExpires = &jwtToken.ExpiresAt.Time
	return
}

func (service *AuthServiceImpl) SignUp(input *data.SignUpRequest) (errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		userFound, errFound = service.Repository.GetByEmail(input.Email)
		errMessage = "User with this email already exists! Please enter valid information."
	} else {
		errMessage = "User with this phone number already exists! Please enter valid information."
		userFound, errFound = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if userFound != nil && userFound.Email == input.Email {
		errCode = http.StatusFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Create new user
	userFound.Email = input.Email
	userFound.PhoneNumber = input.PhoneNumber
	userFound.Password = input.Password
	err = service.Repository.Create(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
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
		utils.JWT_ISSUER_ACTIVATE,
		utils.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisStr,
	)
	if errEncode != nil || jwtToken == nil {
		errCode = http.StatusInternalServerError
		err = errEncode
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

func (service *AuthServiceImpl) ActivateAccount(input *data.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string
	var jwtToken, errDecode = utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecode != nil {
		errCode = http.StatusNotFound
		err = errDecode
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != utils.JWT_ISSUER_ACTIVATE {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisStr)
	if !isTokenValid {
		errMessage = "Token already expired! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if code is valid
	fmt.Printf("\nRequested code: %d\n", input.Code)
	if jwtToken.Code <= 0 || jwtToken.Code != input.Code {
		errMessage = "Invalid code! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwtToken.UserId)
	var userFound, errFound = service.Repository.GetById(userId)
	if errFound != nil || userFound == nil {
		errMessage = "User not found! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if account is activated
	if userFound.IsActivated {
		errMessage = "User account is already activated! Please sign in and start using our services."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Update account
	var tmpActivatedAt = time.Now()
	userFound.ActivatedAt = &tmpActivatedAt
	userFound.IsActivated = true
	err = service.Repository.Update(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Create user info
	var userInfo = &model.UserInfo{}
	err = service.Repository.CreateUserInfo(userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Update user with user info id
	userFound.UserInfoId = userInfo.ID
	err = service.Repository.Update(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	activatedAt = userFound.ActivatedAt

	// Invalidate token
	config.DeleteRedisStr(utils.GetJWTCachedKey(jwtToken))

	// Send welcome message
	// TODO
	return
}

func (service *AuthServiceImpl) ResetPasswordInit(input *data.ResetPasswordInitRequest) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		userFound, errFound = service.Repository.GetByEmail(input.Email)
		errMessage = "User not found! Please enter valid information."
	} else {
		errMessage = "User not found! Please enter valid information."
		userFound, errFound = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if userFound == nil || userFound.ID <= 0 {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Generate new random code
	var randomCode, errRandomCode = utils.GenerateRandomCode(5)
	if errRandomCode != nil {
		errCode = http.StatusInternalServerError
		err = errRandomCode
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
		utils.JWT_ISSUER_RESET_CODE,
		expires,
		config.Keys.JwtPrivateKey,
		config.SetRedisStr,
	)
	if errEncode != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = errEncode
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

func (service *AuthServiceImpl) ResetPasswordCode(input *data.ResetPasswordCodeRequest) (token string, errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string
	var jwtToken, errDecode = utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecode != nil {
		errCode = http.StatusNotFound
		err = errDecode
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != utils.JWT_ISSUER_RESET_CODE {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisStr)
	if !isTokenValid {
		errMessage = "Token already expired! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if code is valid
	if jwtToken.Code <= 0 || jwtToken.Code != input.Code {
		errMessage = "Invalid code! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwtToken.UserId)
	var userFound, errFound = service.Repository.GetById(userId)
	if errFound != nil || userFound == nil {
		errMessage = "User not found! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
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
		utils.JWT_ISSUER_SESSION,
		utils.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisStr,
	)
	if errEncode != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = errEncode
	}
	token = newToken
	return
}

func (service *AuthServiceImpl) ResetPasswordNewPassword(input *data.ResetPasswordNewPasswordRequest) (errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string
	var jwtToken, errDecode = utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecode != nil {
		errCode = http.StatusForbidden
		err = errDecode
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != utils.JWT_ISSUER_RESET_PASSWORD {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisStr)
	if !isTokenValid {
		errMessage = "Token already expired! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwtToken.UserId)
	var userFound, errFound = service.Repository.GetById(userId)
	if errFound != nil || userFound == nil {
		errMessage = "User not found! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Update user password
	var userUpdated, errUpdate = service.Repository.UpdatePasswordById(userId, input.NewPassword)
	if errUpdate != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = errUpdate
		return
	}

	// Invalidate token
	config.DeleteRedisStr(utils.GetJWTCachedKey(jwtToken))
	return
}

func (service *AuthServiceImpl) SignOut(token string) (errCode int, err error) {
	// Extract token information and validate the token
	var errMessage string
	var jwtToken, errDecode = utils.DecodeJWTToken(token, config.Keys.JwtPublicKey)
	if errDecode != nil || jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != utils.JWT_ISSUER_SESSION {
		errMessage = "Invalid token! Please enter valid information."
		errCode = http.StatusUnauthorized
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var isTokenValid = utils.ValidateJWTToken(token, jwtToken, config.CheckRedisStrFromArrayStr(token))
	if !isTokenValid {
		errMessage = "Token already expired! Please enter valid information."
		errCode = http.StatusUnauthorized
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Invalidate token
	var sessions, errSession = config.GetRedisArrayStr(fmt.Sprintf("%d", jwtToken.UserId))
	if errSession != nil {
		errMessage = "Token already expired! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var tokenIndex = slices.Index(sessions, token)
	if tokenIndex < 0 {
		errMessage = "Token already expired! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var errDel = config.DeleteRedisStrFromArrayStr(fmt.Sprintf("%d", jwtToken.UserId), int64(tokenIndex))
	if errDel != nil {
		errMessage = "Error occurred when deleting cached session! Please try again later."
		errCode = http.StatusInternalServerError
		err = fmt.Errorf("%s", errMessage)
		return
	}
	return
}
