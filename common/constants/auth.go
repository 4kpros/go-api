package constants

// API security name
const SECURITY_AUTH_NAME = "Bearer token"

// Authentication - login methods
const AUTH_LOGIN_METHOD_DEFAULT = "default"
const AUTH_LOGIN_METHOD_PROVIDER = "provider"

var AUTH_LOGIN_METHODS = []string{
	AUTH_LOGIN_METHOD_DEFAULT,
	AUTH_LOGIN_METHOD_PROVIDER,
}

// Authentication - providers
const AUTH_PROVIDER_GOOGLE = "google"
const AUTH_PROVIDER_FACEBOOK = "facebook"

var AUTH_PROVIDERS = []string{
	AUTH_PROVIDER_GOOGLE,
	AUTH_PROVIDER_FACEBOOK,
}
var AUTH_LOGIN_WITH_FACEBOOK_REQUIRED_SCOPES = []string{
	"email",
	"public_profile",
}

// Multiple Factor Authentication
const AUTH_MFA_METHOD_EMAIL = "email"
const AUTH_MFA_METHOD_PHONE = "phone"
const AUTH_MFA_METHOD_AUTHENTICATOR = "authenticator"

var AUTH_MFA_METHODS = []string{
	AUTH_MFA_METHOD_EMAIL,
	AUTH_MFA_METHOD_PHONE,
	AUTH_MFA_METHOD_AUTHENTICATOR,
}

// JWT
var JWT_ISSUER_SESSION string
var JWT_ISSUER_SESSION_API_KEY string

var JWT_ISSUER_AUTH_ACTIVATE string
var JWT_ISSUER_AUTH_FORGOT_PASSWORD_CODE string
var JWT_ISSUER_AUTH_FORGOT_PASSWORD_NEW_PASSWORD string
var JWT_ISSUER_AUTH []string

// Initializes the JWT issuer with the provided passphrase.
// Needs to be called inside the "init" function in the "main.go" file.
func InitializeJwtIssuerConst(
	sessionPassPhrase string,
	sessionApiKeyPassPhrase string,
	sessionAuthPassPhrase string,
) {
	JWT_ISSUER_SESSION = sessionPassPhrase + "issuer_session"
	JWT_ISSUER_SESSION_API_KEY = sessionApiKeyPassPhrase + "issuer_session_api_key"
	JWT_ISSUER_AUTH_ACTIVATE = sessionAuthPassPhrase + "issuer_auth_activate"
	JWT_ISSUER_AUTH_FORGOT_PASSWORD_CODE = sessionAuthPassPhrase + "issuer_auth_forgot_password_code"
	JWT_ISSUER_AUTH_FORGOT_PASSWORD_NEW_PASSWORD = sessionAuthPassPhrase + "issuer_auth_forgot_password_new_password"

	JWT_ISSUER_AUTH = []string{
		JWT_ISSUER_AUTH_ACTIVATE,
		JWT_ISSUER_AUTH_FORGOT_PASSWORD_CODE,
		JWT_ISSUER_AUTH_FORGOT_PASSWORD_NEW_PASSWORD,
	}
}
