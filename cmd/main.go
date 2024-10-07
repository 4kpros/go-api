package main

import (
	"github.com/4kpros/go-api/cmd/api"
	"github.com/4kpros/go-api/cmd/migrate"
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"go.uber.org/zap"
)

var initError error = nil

func init() {
	// Setup logger
	helpers.SetupLogger()

	// Load env variables
	errAppEnv := config.LoadAppEnv(".")
	if errAppEnv != nil {
		initError = errAppEnv
		helpers.Logger.Warn(
			"Failed to load app ENV vars!",
			zap.String("Error", errAppEnv.Error()),
		)
	} else {
		helpers.Logger.Warn(
			"App ENV variables loaded!",
		)
	}

	// Setup argon2id params for crypto
	_, errArgonCryptoParamsUtils := utils.EncryptWithArgon2id("")
	if errArgonCryptoParamsUtils != nil {
		initError = errArgonCryptoParamsUtils
		helpers.Logger.Warn(
			"Failed to setup argon2id params!",
			zap.String("Error", errArgonCryptoParamsUtils.Error()),
		)
	} else {
		helpers.Logger.Warn(
			"Argon2id crypto set ok!",
		)
	}

	// Connect to postgres database
	errPostgresDB := config.ConnectToPostgresDB()
	if errPostgresDB != nil {
		initError = errPostgresDB
		helpers.Logger.Warn(
			"Failed to connect to Postgres database!",
			zap.String("Error", errPostgresDB.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to Postgres database!",
		)
	}

	// Connect to redis
	errRedis := config.ConnectToRedis()
	if errRedis != nil {
		initError = errRedis
		helpers.Logger.Warn(
			"Failed to connect to Redis!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to Redis!",
		)
	}

	// Load pem
	errPem := config.LoadPem()
	if errPem != nil {
		initError = errPem
		helpers.Logger.Warn(
			"Failed to load all pem files!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"All pem files loaded!",
		)
	}
}

// @title API Documentation
// @version 1.0
// @description This is the API documentation

// @contact.name Prosper Abouar
// @contact.email prosper.abouar@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey X-API-Key
// @in header
// @name X-API-Key
// @description Enter the API key to have access

// @securityDefinitions.apikey Bearer
// @in header
// @name Bearer
// @description Enter Bearer with space and your token
func main() {
	if initError != nil {
		helpers.Logger.Warn(
			"There are some errors when initializing app!",
			zap.String("Error", "Please fix previous errors before."),
		)
		return
	}
	migrate.Start()
	api.Start()
}
