package main

import (
	"github.com/4kpros/go-api/cmd/api"
	"github.com/4kpros/go-api/cmd/migrate"
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"go.uber.org/zap"
)

var errInit error = nil

func init() {
	// Setup logger
	helpers.SetupLogger()

	// Load env
	errEnv := config.LoadEnv(".")
	if errEnv != nil {
		errInit = errEnv
		helpers.Logger.Error(
			"Failed to load env!",
			zap.String("Error", errEnv.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Env loaded!",
		)
	}

	// Test argon2id with empty password
	_, errArgon2id := utils.EncodeArgon2id("")
	if errArgon2id != nil {
		errInit = errArgon2id
		helpers.Logger.Error(
			"Failed to setup argon2id!",
			zap.String("Error", errArgon2id.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Argon2id setup ok!",
		)
	}

	// Connect database
	errDB := config.ConnectDatabase()
	if errDB != nil {
		errInit = errDB
		helpers.Logger.Error(
			"Failed to connect to database!",
			zap.String("Error", errDB.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to database!",
		)
	}

	// Connect redis
	errRedis := config.ConnectRedis()
	if errRedis != nil {
		errInit = errRedis
		helpers.Logger.Error(
			"Failed to connect to Redis!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"Connected to Redis!",
		)
	}

	// Load keys
	errKeys := config.LoadKeys()
	if errKeys != nil {
		errInit = errKeys
		helpers.Logger.Error(
			"Failed to load all keys!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"All keys loaded!",
		)
	}

	// Load OpenAPI templates
	errOpenAPITemplates := config.LoadOpenAPITemplates()
	if errOpenAPITemplates != nil {
		errInit = errOpenAPITemplates
		helpers.Logger.Error(
			"Failed to load OpenAPI templates!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info(
			"All OpenAPI templates loaded!",
		)
	}
}

func main() {
	if errInit != nil {
		helpers.Logger.Warn(
			"There are some errors when initializing app!",
			zap.String("Error", "Please fix previous errors before."),
		)
		return
	}
	migrate.Start()
	api.Start()
}
