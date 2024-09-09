package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth/data/request"
	"github.com/4kpros/go-api/services/user/model"
)

type AuthService interface {
	SignIn(deviceName string, reqData *request.SignInRequest) (validateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error)
	SignInWithProvider(deviceName string, reqData *request.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error)

	SignUp(reqData *request.SignUpRequest) (token string, errCode int, err error)

	ActivateAccount(reqData *request.ActivateAccountRequest) (date *time.Time, errCode int, err error)

	ResetPasswordInit(reqData *request.ResetPasswordInitRequest) (token string, errCode int, err error)
	ResetPasswordCode(reqData *request.ResetPasswordCodeRequest) (token string, errCode int, err error)
	ResetPasswordNewPassword(reqData *request.ResetPasswordNewPasswordRequest) (errCode int, err error)
}

type AuthServiceImpl struct {
	Repository AuthRepository
}

func NewAuthServiceImpl(repository AuthRepository) AuthService {
	return &AuthServiceImpl{Repository: repository}
}

func (service *AuthServiceImpl) SignIn(deviceName string, reqData *request.SignInRequest) (activateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if len(reqData.Email) > 0 {
		userFound, errFound = service.Repository.FindByEmail(reqData.Email)
		errMessage = "Invalid email or password! Please enter valid information."
	} else {
		errMessage = "Invalid phone number or password! Please enter valid information."
		userFound, errFound = service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	}
	isPasswordMatches, errCompare := utils.CompareToArgon2id(reqData.Password, userFound.Password)
	if errFound != nil || errCompare != nil || userFound == nil || userFound.Email != reqData.Email || !isPasswordMatches {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if account is activated
	if !userFound.IsActivated {
		var expires = utils.NewExpiresDateDefault()
		var jwt = &types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Issuer:  utils.JwtIssuerActivate,
		}
		var randomCode, _ = utils.GenerateRandomCode(utils.GetCachedKey(jwt), 5)
		jwt.Code = randomCode
		newJwt, tokenStr, errEncrypt := utils.EncryptJWTToken(
			jwt,
			config.AppPem.JwtPrivateKey,
			true,
		)
		if errEncrypt != nil || newJwt == nil {
			errCode = http.StatusInternalServerError
			err = errEncrypt
			return
		}
		activateAccountToken = tokenStr
		errMessage = "Account found but not activated! Please activate your account before."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", errMessage)

		// Send code to email or phone number
		if len(reqData.Email) > 0 {
			// TODO send code to email
		} else {
			// TODO send code to phone number
		}

		return
	}

	// Generate new token
	var expires = utils.NewExpiresDateSignIn(reqData.StayConnected)
	_, tokenStr, errEncrypt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Device:  deviceName,
			Issuer:  utils.JwtIssuerSession,
		},
		config.AppPem.JwtPrivateKey,
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

func (service *AuthServiceImpl) SignInWithProvider(deviceName string, reqData *request.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Validate provider token and update user
	var errMessage string
	if !types.AuthProviders[reqData.Provider] {
		errMessage = "Invalid provider! Please enter valid information."
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", errMessage)
		return
	}
	var providerUserId = "Test"
	if len(providerUserId) <= 0 {
		errMessage = "Invalid provider or token! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}
	userFound, errFound := service.Repository.FindByProvider(reqData.Provider, providerUserId)
	if errFound != nil || userFound == nil || userFound.Provider != reqData.Provider {
		// Save new user
		user := &model.User{
			Provider:       reqData.Provider,
			ProviderUserId: providerUserId,
		}
		user.Role = types.RoleCustomer
		err = service.Repository.Create(user)
		if err != nil {
			errCode = http.StatusInternalServerError
			return
		}
	}

	// Generate new token
	var expires = utils.NewExpiresDateSignIn(true)
	_, tokenStr, errEncrypt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Device:  deviceName,
			Issuer:  utils.JwtIssuerSession,
		},
		config.AppPem.JwtPrivateKey,
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

func (service *AuthServiceImpl) SignUp(reqData *request.SignUpRequest) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if len(reqData.Email) > 0 {
		userFound, errFound = service.Repository.FindByEmail(reqData.Email)
		errMessage = "User with this email already exists! Please enter valid information."
	} else {
		errMessage = "User with this phone number already exists! Please enter valid information."
		userFound, errFound = service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	}
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if userFound != nil && userFound.Email == reqData.Email {
		errCode = http.StatusFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Create new user
	userFound.Email = reqData.Email
	userFound.PhoneNumber = reqData.PhoneNumber
	userFound.Password = reqData.Password
	userFound.Role = types.RoleCustomer
	err = service.Repository.Create(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Generate new token
	var expires = utils.NewExpiresDateDefault()
	var jwt = &types.JwtToken{
		UserId:  userFound.ID,
		Expires: *expires,
		Issuer:  utils.JwtIssuerActivate,
	}
	var randomCode, _ = utils.GenerateRandomCode(utils.GetCachedKey(jwt), 5)
	jwt.Code = randomCode
	newJwt, tokenStr, errEncrypt := utils.EncryptJWTToken(
		jwt,
		config.AppPem.JwtPrivateKey,
		true,
	)
	if errEncrypt != nil || newJwt == nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
	}
	token = tokenStr

	// Send code to email or phone number
	if len(reqData.Email) > 0 {
		// TODO send code to email
	} else {
		// TODO send code to phone number
	}

	return
}

func (service *AuthServiceImpl) ActivateAccount(reqData *request.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information
	var errMessage string
	jwt, errDecrypt := utils.DecryptJWTToken(reqData.Token, config.AppPem.JwtPublicKey)
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
	if jwt.Code != reqData.Code {
		errMessage = "Invalid code! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwt.UserId)
	userFound, errFound := service.Repository.FindById(userId)
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
	tmpActivatedAt := time.Now()
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

func (service *AuthServiceImpl) ResetPasswordInit(reqData *request.ResetPasswordInitRequest) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errFound error
	var errMessage string
	if len(reqData.Email) > 0 {
		userFound, errFound = service.Repository.FindByEmail(reqData.Email)
		errMessage = "User not found! Please enter valid information."
	} else {
		errMessage = "User not found! Please enter valid information."
		userFound, errFound = service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
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
	cacheKey := utils.GetCachedKey(&types.JwtToken{
		UserId: userFound.ID,
		Issuer: utils.JwtIssuerResetCode,
	})
	randomCode, errRandomCode := utils.GenerateRandomCode(cacheKey, 5)
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
		},
		config.AppPem.JwtPrivateKey,
		true,
	)
	if errEncrypt != nil || newJwt == nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
		return
	}
	token = tokenStr

	// Send code to email or phone number
	if len(reqData.Email) > 0 {
		// TODO send code to email
	} else {
		// TODO send code to phone number
	}

	return
}

func (service *AuthServiceImpl) ResetPasswordCode(reqData *request.ResetPasswordCodeRequest) (token string, errCode int, err error) {
	// Extract token information
	var errMessage string
	jwt, errDecrypt := utils.DecryptJWTToken(reqData.Token, config.AppPem.JwtPublicKey)
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
	if jwt.Code != reqData.Code {
		errMessage = "Invalid code! Please enter valid information."
		errCode = http.StatusPreconditionFailed
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwt.UserId)
	userFound, errFound := service.Repository.FindById(userId)
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
	newJwt, tokenStr, errEncrypt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  userFound.ID,
			Expires: *expires,
			Issuer:  utils.JwtIssuerResetNewPassword,
		},
		config.AppPem.JwtPrivateKey,
		true,
	)
	if errEncrypt != nil || newJwt == nil {
		errCode = http.StatusInternalServerError
		err = errEncrypt
	}
	token = tokenStr

	return
}

func (service *AuthServiceImpl) ResetPasswordNewPassword(reqData *request.ResetPasswordNewPasswordRequest) (errCode int, err error) {
	// Extract token information
	var errMessage string
	jwt, errDecrypt := utils.DecryptJWTToken(reqData.Token, config.AppPem.JwtPublicKey)
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
	userFound, errFound := service.Repository.FindById(userId)
	if errFound != nil || userFound == nil {
		errMessage = "User not found! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Update user password
	userUpdated, errUpdate := service.Repository.UpdatePasswordById(userId, reqData.NewPassword)
	if errUpdate != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = errUpdate
		return
	}

	// Invalidate token
	config.DeleteRedisVal(utils.GetCachedKey(jwt))

	return
}
