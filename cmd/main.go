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

	// Load templates
	errTemplate := config.LoadTemplates()
	if errTemplate != nil {
		initError = errTemplate
		helpers.Logger.Warn(
			"Failed to load all template files!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"All template files loaded!",
		)
	}
}

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
