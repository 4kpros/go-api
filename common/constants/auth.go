package constants

// API security name
const SecurityAuthName = "Bearer token"

// Authentication - login methods
const AuthLoginMethodDefault = "default"
const AuthLoginMethodProvider = "provider"

var AUTH_LOGIN_METHODS = []string{
	AuthLoginMethodDefault,
	AuthLoginMethodProvider,
}

// Authentication - providers
const AuthProviderGoogle = "google"
const AuthProviderFacebook = "facebook"

var AuthProviders = []string{
	AuthProviderGoogle,
	AuthProviderFacebook,
}
var AuthLoginWithFacebookRequiredScopes = []string{
	"email",
	"public_profile",
}

// Multiple Factor Authentication
const AuthMfaMethodEmail = "email"
const AuthMfaMethodPhone = "phone"
const AuthMfaMethodAuthenticator = "authenticator"

var AuthMfaMethods = []string{
	AuthMfaMethodEmail,
	AuthMfaMethodPhone,
	AuthMfaMethodAuthenticator,
}

// JWT
var JwtIssuerSession string
var JwtIssuerSessionApiKey string

var JwtIssuerAuthActivate string
var JwtIssuerAuthForgotPasswordCode string
var JwtIssuerAuthForgotPasswordNewPassword string
var JwtIssuerAuths []string

// InitializeJwtIssuerConst Initializes the JWT issuer with the provided passphrase.
// Needs to be called inside the "init" function in the "main.go" file.
func InitializeJwtIssuerConst(
	sessionPassPhrase string,
	sessionApiKeyPassPhrase string,
	sessionAuthPassPhrase string,
) {
	JwtIssuerSession = sessionPassPhrase + "issuer_session"
	JwtIssuerSessionApiKey = sessionApiKeyPassPhrase + "issuer_session_api_key"
	JwtIssuerAuthActivate = sessionAuthPassPhrase + "issuer_auth_activate"
	JwtIssuerAuthForgotPasswordCode = sessionAuthPassPhrase + "issuer_auth_forgot_password_code"
	JwtIssuerAuthForgotPasswordNewPassword = sessionAuthPassPhrase + "issuer_auth_forgot_password_new_password"

	JwtIssuerAuths = []string{
		JwtIssuerAuthActivate,
		JwtIssuerAuthForgotPasswordCode,
		JwtIssuerAuthForgotPasswordNewPassword,
	}
}
