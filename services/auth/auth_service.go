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
	"github.com/4kpros/go-api/services/user"
	"github.com/4kpros/go-api/services/user/model"
)

type AuthService struct {
	Repository *user.UserRepository
}

func NewAuthService(repository *user.UserRepository) *AuthService {
	return &AuthService{Repository: repository}
}

// Login with email or phone number
func (service *AuthService) SignIn(input *data.SignInRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, activateAccountToken string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		userFound, err = service.Repository.GetByEmail(input.Email)
		errMessage = "Invalid email or password! Please enter valid information."
	} else {
		errMessage = "Invalid phone number or password! Please enter valid information."
		userFound, err = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if err != nil || userFound == nil || userFound.Email != input.Email {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}
	isPasswordMatches, err := utils.CompareArgon2id(input.Password, userFound.Password)
	if err != nil || !isPasswordMatches {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if account is activated
	if userFound.IsActivated {
		// Generate new token
		var accessJwtToken *types.JwtToken
		accessJwtToken, accessToken, err = utils.EncodeJWTToken(
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
			config.AppendToRedisStringList,
		)
		if err != nil || accessJwtToken == nil || len(accessToken) <= 0 {
			errCode = http.StatusInternalServerError
			err = constants.HTTP_500_ERROR_MESSAGE("encode JWT token")
		}

		accessExpires = &accessJwtToken.ExpiresAt.Time
		return
	}

	// For non activated user account, generate new random code and token with
	// issuer JWT_ISSUER_AUTH_ACTIVATE and send code to email or phone number
	randomCode := 0
	randomCode, err = utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate random code")
		return
	}

	// Generate new token
	var activateAccountJwtToken *types.JwtToken
	activateAccountJwtToken, activateAccountToken, err = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JWT_ISSUER_AUTH_ACTIVATE,
		utils.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || activateAccountJwtToken == nil || len(activateAccountToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("encode jwt token")
	}
	errCode = http.StatusForbidden
	err = fmt.Errorf("%s", "Account found but not activated! Please activate your account to start using your services.")

	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go utils.SendMail(
			fmt.Sprintf("%s - Activate your account", config.Env.AppName),
			fmt.Sprintf("The code to activate your account is %d", randomCode),
			input.Email,
		)
	} else {
		go utils.SendSMS(
			fmt.Sprintf("The code to activate your account is %d.", randomCode),
			fmt.Sprintf("+%d", input.PhoneNumber),
		)
	}
	return
}

// Login with provider like Google and Facebook
func (service *AuthService) SignInWithProvider(input *data.SignInWithProviderRequest, device *data.SignInDevice) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Validate provider token and update user
	var user = &model.User{}
	var expires int64 = 0
	if input.Provider == constants.AUTH_PROVIDER_GOOGLE {
		googleUser, errGoogleUser := utils.VerifyGoogleIDToken(input.Token)
		if errGoogleUser != nil || googleUser == nil || len(googleUser.ID) <= 0 {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
			return
		}
		if googleUser.Expires <= time.Now().Unix() {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Token already expired! Please enter valid information.")
			return
		}
		expires = googleUser.Expires
		user.FromGoogleUser(googleUser)
	} else if input.Provider == constants.AUTH_PROVIDER_FACEBOOK {
		facebookUser, errFacebookUser := utils.VerifyFacebookToken(input.Token)
		if errFacebookUser != nil || facebookUser == nil || len(facebookUser.ID) <= 0 {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
			return
		}
		if facebookUser.Expires <= time.Now().Unix() {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Token already expired! Please enter valid information.")
			return
		}
		expires = facebookUser.Expires
		user.FromFacebookUser(facebookUser)
	} else {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
		return
	}
	user.Provider = input.Provider

	// Save user if it's not in database
	userFound, err := service.Repository.GetByProvider(input.Provider, user.ProviderUserId)
	if err != nil || userFound == nil {
		userFound, err = service.Repository.Create(
			&model.User{
				Provider:       input.Provider,
				ProviderUserId: user.ProviderUserId,
				SignInMethod:   constants.AUTH_LOGIN_METHOD_PROVIDER,
			},
		)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.HTTP_500_ERROR_MESSAGE("create user on database")
			return
		}
		var userInfo *model.UserInfo
		userInfo, err = service.Repository.CreateUserInfo(&user.UserInfo)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.HTTP_500_ERROR_MESSAGE("create user info on database")
			return
		}
		userFound.UserInfo = *userInfo
	}

	// Generate new token
	expiresTime := time.Unix(expires, 0)
	jwtToken, accessToken, err := utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: device.Platform,
			Device:   device.DeviceName,
			App:      device.App,
		},
		constants.JWT_ISSUER_SESSION,
		&expiresTime,
		config.Keys.JwtPrivateKey,
		config.AppendToRedisStringList,
	)
	if err != nil || jwtToken == nil || len(accessToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("encode JWT token")
	}
	accessExpires = &jwtToken.ExpiresAt.Time
	return
}

// Register with email or phone number
func (service *AuthService) SignUp(input *data.SignUpRequest) (activateAccountToken string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		userFound, err = service.Repository.GetByEmail(input.Email)
		errMessage = "user email"
	} else {
		errMessage = "user phone number"
		userFound, err = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if err != nil {
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
	userFound.SignInMethod = constants.AUTH_LOGIN_METHOD_DEFAULT
	createdUser, err := service.Repository.Create(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user on database")
		return
	}

	// Since the new user account is not activated, we generate code with
	// issuer JWT_ISSUER_AUTH_ACTIVATE and send code to email or phone number
	randomCode := 0
	randomCode, err = utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate random code")
		return
	}

	// Generate new token
	var activateAccountJwtToken *types.JwtToken
	activateAccountJwtToken, activateAccountToken, err = utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   createdUser.ID,
			RoleId:   createdUser.RoleId,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JWT_ISSUER_AUTH_ACTIVATE,
		utils.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || activateAccountJwtToken == nil || len(activateAccountToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("encode jwt token")
	}

	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go utils.SendMail(
			fmt.Sprintf("%s - Activate your account", config.Env.AppName),
			fmt.Sprintf("The code to activate your account is %d", randomCode),
			input.Email,
		)
	} else {
		go utils.SendSMS(
			fmt.Sprintf("The code to activate your account is %d.", randomCode),
			fmt.Sprintf("+%d", input.PhoneNumber),
		)
	}
	return
}

// Activate user account
func (service *AuthService) ActivateAccount(input *data.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information and validate the token
	errMessage := "Invalid or expired token! Please enter valid information."
	jwtToken, err := utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_AUTH_ACTIVATE {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	isTokenValid := utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisString)
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

	// Check if account is activated
	userFound, err := service.Repository.GetById(jwtToken.UserId)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}
	if userFound.IsActivated {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User account is already activated! Please sign in and start using our services.")
		return
	}

	// Create user info and MFA
	userInfo, err := service.Repository.CreateUserInfo(&model.UserInfo{})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user info")
		return
	}
	userMfa, err := service.Repository.CreateUserMfa(&model.UserMfa{})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user MFA")
		return
	}

	// Update account
	tmpActivatedAt := time.Now()
	userFound.ActivatedAt = &tmpActivatedAt
	userFound.IsActivated = true
	userFound.UserInfoId = userInfo.ID
	userFound.UserInfo = *userInfo
	userFound.UserMfaId = userMfa.ID
	userFound.UserMfa = *userMfa
	updatedUser, err := service.Repository.UpdateUserActivation(userFound.ID, userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user" + err.Error())
		return
	}
	activatedAt = updatedUser.ActivatedAt

	// Invalidate token
	config.DeleteRedisString(utils.GetJWTCachedKey(jwtToken))

	// Send welcome message
	if utils.IsEmailValid(updatedUser.Email) {
		go utils.SendMail(
			fmt.Sprintf("%s - Welcome", config.Env.AppName),
			fmt.Sprintf("Welcome"),
			updatedUser.Email,
		)
	}
	return
}

// Forgot password step 1: request forgot password
func (service *AuthService) ForgotPasswordInit(input *data.ForgotPasswordInitRequest) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errMessage string
	if utils.IsEmailValid(input.Email) {
		errMessage = "User with this email"
		userFound, err = service.Repository.GetByEmail(input.Email)
	} else {
		errMessage = "User with this phone number"
		userFound, err = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if err != nil || userFound.ID <= 0 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE(errMessage)
		return
	}

	// Generate new random code
	randomCode, err := utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate random code")
	}
	expires := utils.NewExpiresDateDefault()
	newJwtToken, newToken, err := utils.EncodeJWTToken(
		&types.JwtToken{
			UserId:   userFound.ID,
			RoleId:   userFound.RoleId,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JWT_ISSUER_AUTH_FORGOT_PASSWORD_CODE,
		expires,
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate new JWT token")
		return
	}
	token = newToken

	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go utils.SendMail(
			fmt.Sprintf("%s - Forgot your password", config.Env.AppName),
			fmt.Sprintf("The code to reset your password is %d", randomCode),
			input.Email,
		)
	} else {
		go utils.SendSMS(
			fmt.Sprintf("The code to reset your password is %d.", randomCode),
			fmt.Sprintf("+%d", input.PhoneNumber),
		)
	}
	return
}

// Forgot password step 2: validate sended code
func (service *AuthService) ForgotPasswordCode(input *data.ForgotPasswordCodeRequest) (token string, errCode int, err error) {
	// Extract token information and validate the token
	errMessage := "Invalid or expired token! Please enter valid information."
	jwtToken, err := utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_AUTH_FORGOT_PASSWORD_CODE {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	isTokenValid := utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisString)
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
	userFound, err := service.Repository.GetById(jwtToken.UserId)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Invalidate token
	config.DeleteRedisString(utils.GetJWTCachedKey(jwtToken))

	// Generate new token
	newJwtToken, newToken, err := utils.EncodeJWTToken(
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
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("generate new JWT token")
	}
	token = newToken
	return
}

// Forgot password step 3: setup new password
func (service *AuthService) ForgotPasswordNewPassword(input *data.ForgotPasswordNewPasswordRequest) (errCode int, err error) {
	// Extract token information and validate the token
	errMessage := "Invalid or expired token! Please enter valid information."
	jwtToken, err := utils.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	if jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_AUTH_FORGOT_PASSWORD_NEW_PASSWORD {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}
	isTokenValid := utils.ValidateJWTToken(input.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMessage)
		return
	}

	// Check if user exists
	userFound, err := service.Repository.GetById(jwtToken.UserId)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user password
	userUpdated, err := service.Repository.UpdateUserPassword(jwtToken.UserId, input.NewPassword)
	if err != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update password")
		return
	}

	// Invalidate token
	_, err = config.DeleteRedisString(utils.GetJWTCachedKey(jwtToken))
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("create user from database")
		return
	}
	return
}

// Logout user with provided token
func (service *AuthService) SignOut(token string) (errCode int, err error) {
	// Extract token information and validate the token
	jwtToken, err := utils.DecodeJWTToken(token, config.Keys.JwtPublicKey)
	if err != nil || jwtToken == nil || jwtToken.UserId <= 0 || jwtToken.Issuer != constants.JWT_ISSUER_SESSION {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}
	isTokenValid := utils.ValidateJWTToken(token, jwtToken, config.CheckValueInRedisList(token))
	if !isTokenValid {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}

	// Invalidate token
	sessions, err := config.GetRedisStringList(utils.GetJWTCachedKey(jwtToken))
	if err != nil {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}
	tokenIndex := slices.Index(sessions, token)
	if tokenIndex < 0 {
		errCode = http.StatusUnauthorized
		err = constants.HTTP_401_ERROR_MESSAGE()
		return
	}
	err = config.RemoveFromRedisStringList(fmt.Sprintf("%d", jwtToken.UserId), int64(tokenIndex))
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("delete cached session")
		return
	}
	return
}
