package migrate

import (
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/config"
	historyModel "github.com/4kpros/go-api/services/history/model"
	permissionModel "github.com/4kpros/go-api/services/permission/model"
	roleModel "github.com/4kpros/go-api/services/role/model"
	userModel "github.com/4kpros/go-api/services/user/model"
)

// Used to load all migrations
func Start() {
	var err = config.DB.AutoMigrate(
		&historyModel.History{},
		&roleModel.Role{},
		&permissionModel.Permission{},
		&userModel.User{},
	)
	helpers.LogMigrations(err)
}
