package migrate

import (
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/role/model"
)

func Migrate() {
	var err = config.DB.AutoMigrate(&model.Role{})
	helpers.PrintMigrationLogs(err, "Role")
}
