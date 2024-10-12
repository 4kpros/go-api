package config

import (
	"github.com/spf13/viper"
)

type Environment struct {
	// API config
	ApiPort      int    `mapstructure:"API_PORT"`
	ApiGroup     string `mapstructure:"API_GROUP"`
	GinMode      string `mapstructure:"GIN_MODE"`
	AllowedHosts string `mapstructure:"ALLOWED_HOSTS"`

	// Postgres database
	PostGresHost     string `mapstructure:"POSTGRES_HOST"`
	PostGresPort     int    `mapstructure:"POSTGRES_PORT"`
	PostGresUserName string `mapstructure:"POSTGRES_USERNAME"`
	PostGresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostGresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostGresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostGresTimeZone string `mapstructure:"POSTGRES_TIME_ZONE"`

	// JWT
	JwtExpiresSignIn              int    `mapstructure:"JWT_EXPIRES_SIGN_IN"`
	JwtExpiresSignInStayConnected int    `mapstructure:"JWT_EXPIRES_SIGN_IN_STAY_CONNECTED"`
	JwtExpiresDefault             int    `mapstructure:"JWT_EXPIRES_DEFAULT"`
	JwtIssuerPassphrase           string `mapstructure:"JWT_ISSUER_PASSPHRASE"`

	// Redis for fast key-value database
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisUserName string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase int    `mapstructure:"REDIS_DB"`

	// Crypto Argon2id for passwords
	ArgonMemoryLeft  int `mapstructure:"ARGON_PARAM_MEMORY_L"`
	ArgonMemoryRight int `mapstructure:"ARGON_PARAM_MEMORY_R"`
	ArgonIterations  int `mapstructure:"ARGON_PARAM_ITERATIONS"`
	ArgonSaltLength  int `mapstructure:"ARGON_PARAM_SALT_LENGTH"`
	ArgonKeyLength   int `mapstructure:"ARGON_PARAM_KEY_LENGTH"`

	// SMTP
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpUsername string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpSender   string `mapstructure:"SMTP_SENDER"`

	// Google
	GooglePlusClientId string `mapstructure:"GOOGLE_PLUS_CLIENT_ID"`

	// Facebook
	FacebookAppName       string `mapstructure:"FACEBOOK_APP_NAME"`
	FacebookAppId         string `mapstructure:"FACEBOOK_APP_ID"`
	FacebookClientSecret  string `mapstructure:"FACEBOOK_CLIENT_SECRET"`
	FacebookApiBaseUrl    string `mapstructure:"FACEBOOK_API_BASE_URL"`
	FacebookDebugTokenUrl string `mapstructure:"FACEBOOK_DEBUG_TOKEN_URL"`
	FacebookProfileUrl    string `mapstructure:"FACEBOOK_PROFILE_URL"`
}

var Env = &Environment{}

// Loads environment variables from the specified file.
func LoadEnv(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var err error
	err = viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(Env)
	}
	return err
}
