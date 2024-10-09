package helpers

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func SetupLogger() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
}

func PrintMigrationLogs(err error, modelName string) {
	var message string
	if err != nil {
		message = fmt.Sprintf("Failed to migrate << %s >> table !", modelName)
		Logger.Error(
			message,
			zap.String("Error", err.Error()),
		)
		return
	}
	message = fmt.Sprintf("Migration for table << %s >> Done !", modelName)
	Logger.Info(
		message,
	)
}
