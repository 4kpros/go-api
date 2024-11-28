package migrate

import (
	"api/common/helpers"
	"api/config"
	historyModel "api/services/history/model"
	schoolModel "api/services/school/common/school/model"
	yearModel "api/services/school/common/year/model"
	departmentModel "api/services/school/university/department/model"
	domainModel "api/services/school/university/domain/model"
	examModel "api/services/school/university/exam/model"
	facultyModel "api/services/school/university/faculty/model"
	levelModel "api/services/school/university/level/model"
	tuModel "api/services/school/university/tu/model"
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

			&yearModel.Year{},

			&schoolModel.School{},
			&schoolModel.SchoolInfo{},
			&schoolModel.SchoolConfig{},
			&schoolModel.SchoolDirector{},

			// Secondary

			// University
			&facultyModel.Faculty{},
			&departmentModel.Department{},
			&domainModel.Domain{},
			&levelModel.Level{},
			&tuModel.TeachingUnit{},
			&tuModel.TeachingUnitProfessor{},
			&examModel.Exam{},
		),
	)
}
