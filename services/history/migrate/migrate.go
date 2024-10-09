package migrate

import (
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/history/model"
)

func Migrate() {
	err := config.DB.AutoMigrate(&model.History{})
	helpers.PrintMigrationLogs(err, "History")
}
