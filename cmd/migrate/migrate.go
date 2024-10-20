package migrate

import (
	"api/common/helpers"
	"api/config"
	historyModel "api/services/admin/history/model"
	permissionModel "api/services/admin/permission/model"
	roleModel "api/services/admin/role/model"
	userModel "api/services/admin/user/model"
)

// Start Loads and applies all migrations.
func Start() {
	helpers.LogMigrations(
		config.DB.AutoMigrate(
			&historyModel.History{},
			&roleModel.Role{},
			&permissionModel.PermissionFeature{},
			&permissionModel.PermissionTable{},
			&userModel.User{},
		),
	)
}
