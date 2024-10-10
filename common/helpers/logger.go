package helpers

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func SetupLogger() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}

func LogMigrations(err error) {
	if err != nil {
		Logger.Error(
			"Failed to migrate some tables!",
			zap.String("Error", err.Error()),
		)
		return
	}
	Logger.Info(
		"Migration done for all tables!",
	)
}
