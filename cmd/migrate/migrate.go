package migrate

import (
	"api/common/helpers"
	"api/config"
	historyModel "api/services/history/model"
	permissionModel "api/services/permission/model"
	roleModel "api/services/role/model"
	userModel "api/services/user/model"
)

// Loads and applies all migrations.
func Start() {
	helpers.LogMigrations(
		config.DB.AutoMigrate(
			&historyModel.History{},
			&roleModel.Role{},
			&permissionModel.Permission{},
			&userModel.User{},
		),
	)
}
