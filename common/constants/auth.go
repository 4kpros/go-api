package constants

const SECURITY_AUTH_NAME = "Bearer token"

const AUTH_PROVIDER_GOOGLE = "google"
const AUTH_PROVIDER_FACEBOOK = "facebook"

var AUTH_PROVIDERS = []string{
	AUTH_PROVIDER_GOOGLE,
	AUTH_PROVIDER_FACEBOOK,
}

var JWT_ISSUER_SESSION string
var JWT_ISSUER_SESSION_API_KEY string
var JWT_ISSUER_AUTH_ACTIVATE string
var JWT_ISSUER_AUTH_FORGOT_PASSWORD_CODE string
var JWT_ISSUER_AUTH_FORGOT_PASSWORD_NEW_PASSWORD string

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
}
