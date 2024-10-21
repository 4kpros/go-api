package config

import (
	"api/common/constants"

	"github.com/spf13/viper"
)

type Environment struct {
	// Application config
	AppPort  int    `mapstructure:"APP_PORT"`
	AppName  string `mapstructure:"APP_NAME"`
	Hostname string `mapstructure:"HOST_NAME"`

	// API config
	ApiGroup     string `mapstructure:"API_GROUP"`
	GinMode      string `mapstructure:"GIN_MODE"`
	AllowedHosts string `mapstructure:"ALLOWED_HOSTS"`

	// Redis for fast memory key-value storage
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisUserName string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase int    `mapstructure:"REDIS_DB"`

	// Postgres database
	PostGresHost     string `mapstructure:"POSTGRES_HOST"`
	PostGresPort     int    `mapstructure:"POSTGRES_PORT"`
	PostGresUserName string `mapstructure:"POSTGRES_USERNAME"`
	PostGresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostGresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostGresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostGresTimeZone string `mapstructure:"POSTGRES_TIME_ZONE"`

	// Argon2id to hash password
	ArgonMemoryLeft  int `mapstructure:"ARGON_PARAM_MEMORY_L"`
	ArgonMemoryRight int `mapstructure:"ARGON_PARAM_MEMORY_R"`
	ArgonIterations  int `mapstructure:"ARGON_PARAM_ITERATIONS"`
	ArgonSaltLength  int `mapstructure:"ARGON_PARAM_SALT_LENGTH"`
	ArgonKeyLength   int `mapstructure:"ARGON_PARAM_KEY_LENGTH"`

	// JWT for authentication and user session
	JwtExpiresLogin                  int    `mapstructure:"JWT_EXPIRES_LOGIN"`
	JwtExpiresLoginStayConnected     int    `mapstructure:"JWT_EXPIRES_LOGIN_STAY_CONNECTED"`
	JwtExpiresDefault                int    `mapstructure:"JWT_EXPIRES_DEFAULT"`
	JwtIssuerSessionPassphrase       string `mapstructure:"JWT_ISSUER_SESSION_PASSPHRASE"`
	JwtIssuerSessionApiKeyPassphrase string `mapstructure:"JWT_ISSUER_SESSION_API_KEY_PASSPHRASE"`
	JwtIssuerAuthPassphrase          string `mapstructure:"JWT_ISSUER_AUTH_PASSPHRASE"`

	JwtIssuerProfileUpdatePasswordPassphrase string `mapstructure:"JWT_ISSUER_PROFILE_UPDATE_PASSWORD_PASSPHRASE"`
	JwtIssuerProfileUpdateEmailPassphrase string `mapstructure:"JWT_ISSUER_PROFILE_UPDATE_EMAIL_PASSPHRASE"`
	JwtIssuerProfileUpdatePhoneNumberPassphrase string `mapstructure:"JWT_ISSUER_PROFILE_UPDATE_PHONE_NUMBER_PASSPHRASE"`

	profileUpdatePasswordPassPhrase string,
	profileUpdateEmailPassPhrase string,
	profileUpdatePhoneNumberPassPhrase string,
	// reCAPTCHA
	GoogleReCAPTCHASiteKey string  `mapstructure:"GOOGLE_RECAPTCHA_SITE_KEY"`
	GoogleReCAPTCHAScore   float32 `mapstructure:"GOOGLE_RECAPTCHA_SCORE"`

	// SMTP
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpUsername string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpSender   string `mapstructure:"SMTP_SENDER"`

	// SMS
	TwilioAccountSid   string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioApiKey       string `mapstructure:"TWILIO_API_KEY"`
	TwilioApiSecret    string `mapstructure:"TWILIO_API_SECRET"`
	TwilioSenderNumber string `mapstructure:"TWILIO_SENDER_NUMBER"`

	// Login with Google
	GooglePlusClientId string `mapstructure:"GOOGLE_PLUS_CLIENT_ID"`

	// Login with Facebook
	FacebookAppName       string `mapstructure:"FACEBOOK_APP_NAME"`
	FacebookAppId         string `mapstructure:"FACEBOOK_APP_ID"`
	FacebookClientSecret  string `mapstructure:"FACEBOOK_CLIENT_SECRET"`
	FacebookDebugTokenUrl string `mapstructure:"FACEBOOK_DEBUG_TOKEN_URL"`
	FacebookProfileUrl    string `mapstructure:"FACEBOOK_PROFILE_URL"`
}

var Env = &Environment{}

// LoadEnv Loads environment variables.
func LoadEnv() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(Env)
		if err == nil {
			// Initialize the JWT issuer passphrase after the environment file is loaded
			constants.InitializeJwtIssuerConst(
				Env.JwtIssuerSessionPassphrase,
				Env.JwtIssuerSessionApiKeyPassphrase,
				Env.JwtIssuerAuthPassphrase,
				Env.JwtIssuerProfileUpdatePasswordPassphrase,
				Env.JwtIssuerProfileUpdateEmailPassphrase,
				Env.JwtIssuerProfileUpdatePhoneNumberPassphrase,
			)
		}
	}
	return err
}
