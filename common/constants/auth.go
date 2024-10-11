package constants

const SECURITY_AUTH_NAME = "bearerAuth"

const AUTH_PROVIDER_GOOGLE = "google"
const AUTH_PROVIDER_FACEBOOK = "facebook"

var AUTH_PROVIDERS = []string{
	AUTH_PROVIDER_GOOGLE,
	AUTH_PROVIDER_FACEBOOK,
}

var JWT_ISSUER_SESSION string
var JWT_ISSUER_SESSION_GENERATED string
var JWT_ISSUER_ACTIVATE string
var JWT_ISSUER_FORGOT_PASSWORD_CODE string
var JWT_ISSUER_FORGOT_PASSWORD_NEW_PASSWORD string

// Initializes the JWT issuer with the provided passphrase.
// Needs to be called inside the "init" function in the "main.go" file.
func InitializeJwtIssuerConst(passPhrase string) {
	JWT_ISSUER_SESSION = passPhrase + "issuer_session"
	JWT_ISSUER_SESSION_GENERATED = passPhrase + "issuer_session_generated"
	JWT_ISSUER_ACTIVATE = passPhrase + "issuer_activate"
	JWT_ISSUER_FORGOT_PASSWORD_CODE = passPhrase + "issuer_forgot_password_code"
	JWT_ISSUER_FORGOT_PASSWORD_NEW_PASSWORD = passPhrase + "issuer_forgot_password_new_password"
}
