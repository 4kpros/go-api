package helpers

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// Enable the logger to print beautiful log messages.
func EnableLogger() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}

// Custom log message for migrations.
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
