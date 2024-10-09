package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth/data"
	"github.com/4kpros/go-api/services/user/model"
)

type AuthService interface {
	SignIn(deviceName string, input *data.SignInRequest) (accessToken string, accessExpires *time.Time, errCode int, err error)
	SignInWithProvider(deviceName string, input *data.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error)

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

func (service *AuthServiceImpl) SignIn(deviceName string, input *data.SignInRequest) (accessToken string, accessExpires *time.Time, errCode int, err error) {
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
	var isPasswordMatches, errCompare = utils.CompareToArgon2id(input.Password, userFound.Password)
	if errCompare != nil || !isPasswordMatches {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if account is activated
	if !userFound.IsActivated {
		var randomCode, _ = utils.GenerateRandomCode(5)
		var expires = utils.NewExpiresDateDefault()
		var jwt = &types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Issuer:  utils.JwtIssuerActivate,
			Code:    randomCode,
			Role:    userFound.Role,
		}
		var newJwt, tokenStr, errEncrypt = utils.EncryptJWTToken(
			jwt,
			config.Keys.JwtPrivateKey,
			true,
		)
		if errEncrypt != nil || newJwt == nil {
			errCode = http.StatusInternalServerError
			err = errEncrypt
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
				fmt.Sprintf("Your code: %d \nYour token: %s", randomCode, tokenStr),
				input.Email,
			)
		} else {
			// TODO send code to phone number
		}

		return
	}

	// Generate new token
	var expires = utils.NewExpiresDateSignIn(input.StayConnected)
	var _, tokenStr, errEncrypt = utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Device:  deviceName,
			Issuer:  utils.JwtIssuerSession,
			Role:    userFound.Role,
		},
		config.Keys.JwtPrivateKey,
		false,
	)
	if errEncrypt != nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
	}
	accessToken = tokenStr
	accessExpires = expires
	return
}

func (service *AuthServiceImpl) SignInWithProvider(deviceName string, input *data.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error) {
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
	var expires = utils.NewExpiresDateSignIn(true)
	var _, tokenStr, errEncrypt = utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Device:  deviceName,
			Issuer:  utils.JwtIssuerSession,
			Role:    userFound.Role,
		},
		config.Keys.JwtPrivateKey,
		false,
	)
	if errEncrypt != nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
	}
	accessToken = tokenStr
	accessExpires = expires
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
	var expires = utils.NewExpiresDateDefault()
	var randomCode, _ = utils.GenerateRandomCode(5)
	var jwt = &types.JwtToken{
		UserId:  userFound.ID,
		Expires: *expires,
		Issuer:  utils.JwtIssuerActivate,
		Code:    randomCode,
		Role:    userFound.Role,
	}
	var newJwt, tokenStr, errEncrypt = utils.EncryptJWTToken(
		jwt,
		config.Keys.JwtPrivateKey,
		true,
	)
	if errEncrypt != nil || newJwt == nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
	}

	fmt.Printf("\nGenerated code: %d \n", randomCode)
	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go config.SendMail(
			"Reset password",
			fmt.Sprintf("Your code: %d \nYour token: %s", randomCode, tokenStr),
			input.Email,
		)
	} else {
		// TODO send code to phone number
	}
	return
}

func (service *AuthServiceImpl) ActivateAccount(input *data.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information
	var errMessage string
	var jwt, errDecrypt = utils.DecryptJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecrypt != nil {
		errCode = http.StatusNotFound
		err = errDecrypt
		return
	}
	if jwt == nil || jwt.UserId <= 0 || jwt.Issuer != utils.JwtIssuerActivate {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if code is valid
	fmt.Printf("\nRequested code: %d\n", input.Code)
	if jwt.Code <= 0 || jwt.Code != input.Code {
		errMessage = "Invalid code! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwt.UserId)
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
	config.DeleteRedisVal(utils.GetCachedKey(jwt))

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
	var newJwt, tokenStr, errEncrypt = utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Code:    randomCode,
			Issuer:  utils.JwtIssuerResetCode,
			Role:    userFound.Role,
		},
		config.Keys.JwtPrivateKey,
		true,
	)
	if errEncrypt != nil || newJwt == nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
		return
	}
	token = tokenStr

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
	// Extract token information
	var errMessage string
	var jwt, errDecrypt = utils.DecryptJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecrypt != nil {
		errCode = http.StatusNotFound
		err = errDecrypt
		return
	}
	if jwt == nil || jwt.UserId <= 0 || jwt.Issuer != utils.JwtIssuerResetCode {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if code is valid
	if jwt.Code <= 0 || jwt.Code != input.Code {
		errMessage = "Invalid code! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwt.UserId)
	var userFound, errFound = service.Repository.GetById(userId)
	if errFound != nil || userFound == nil {
		errMessage = "User not found! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Invalidate token
	config.DeleteRedisVal(utils.GetCachedKey(jwt))

	// Generate new token
	var expires = utils.NewExpiresDateDefault()
	var newJwt, tokenStr, errEncrypt = utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Issuer:  utils.JwtIssuerResetNewPassword,
			Role:    userFound.Role,
		},
		config.Keys.JwtPrivateKey,
		true,
	)
	if errEncrypt != nil || newJwt == nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
	}
	token = tokenStr
	return
}

func (service *AuthServiceImpl) ResetPasswordNewPassword(input *data.ResetPasswordNewPasswordRequest) (errCode int, err error) {
	// Extract token information
	var errMessage string
	var jwt, errDecrypt = utils.DecryptJWTToken(input.Token, config.Keys.JwtPublicKey)
	if errDecrypt != nil {
		errCode = http.StatusNotFound
		err = errDecrypt
		return
	}
	if jwt == nil || jwt.UserId <= 0 || jwt.Issuer != utils.JwtIssuerResetNewPassword {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwt.UserId)
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
	config.DeleteRedisVal(utils.GetCachedKey(jwt))
	return
}

func (service *AuthServiceImpl) SignOut(token string) (errCode int, err error) {
	// Extract token information
	var errMessage string
	var jwt, errDecrypt = utils.DecryptJWTToken(token, config.Keys.JwtPublicKey)
	if errDecrypt != nil {
		errCode = http.StatusNotFound
		err = errDecrypt
		return
	}
	if jwt == nil || jwt.UserId <= 0 || jwt.Issuer != utils.JwtIssuerResetNewPassword {
		errMessage = "Invalid or expired token! Please enter valid information."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Invalidate token
	config.DeleteRedisVal(utils.GetCachedKey(jwt))
	return
}
