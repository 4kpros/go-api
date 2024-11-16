package migrate

import (
	"api/common/helpers"
	"api/config"
	historyModel "api/services/history/model"
	schoolModel "api/services/school/common/school/model"
	permissionModel "api/services/user/permission/model"
	roleModel "api/services/user/role/model"
	userModel "api/services/user/user/model"
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
			&userModel.UserMfa{},
			&userModel.UserMfa{},

			&schoolModel.School{},
			&schoolModel.SchoolInfo{},
			&schoolModel.SchoolDirector{},
		),
	)
}
