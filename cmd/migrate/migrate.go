package migrate

import (
	"api/common/constants"
	"api/common/helpers"
	"api/config"
	historyModel "api/services/history/model"
	schoolModel "api/services/school/common/school/model"
	yearModel "api/services/school/common/year/model"
	classModel "api/services/school/secondary/class/model"
	pupilModel "api/services/school/secondary/pupil/model"
	sectionModel "api/services/school/secondary/section/model"
	subjectModel "api/services/school/secondary/subject/model"
	testModel "api/services/school/secondary/test/model"
	departmentModel "api/services/school/university/department/model"
	domainModel "api/services/school/university/domain/model"
	examModel "api/services/school/university/exam/model"
	facultyModel "api/services/school/university/faculty/model"
	levelModel "api/services/school/university/level/model"
	studentModel "api/services/school/university/student/model"
	tuModel "api/services/school/university/tu/model"
	"api/services/user/permission"
	permissionData "api/services/user/permission/data"
	permissionModel "api/services/user/permission/model"
	"api/services/user/role"
	roleModel "api/services/user/role/model"
	"api/services/user/user"
	userModel "api/services/user/user/model"
	"time"
)

// Start Loads and applies all migrations.
func Start() error {
	err := config.DB.AutoMigrate(
		// History
		&historyModel.History{},

		// User
		&roleModel.Role{},
		&permissionModel.PermissionFeature{},
		&permissionModel.PermissionTable{},
		&userModel.User{},
		&userModel.UserMfa{},
		&userModel.UserMfa{},

		// School
		&yearModel.Year{},
		&schoolModel.School{},
		&schoolModel.SchoolInfo{},
		&schoolModel.SchoolConfig{},
		&schoolModel.SchoolDirector{},

		// Secondary
		&sectionModel.Section{},
		&classModel.Class{},
		&subjectModel.Subject{},
		&subjectModel.SubjectProfessor{},
		&pupilModel.Pupil{},
		&testModel.Test{},

		// University
		&facultyModel.Faculty{},
		&departmentModel.Department{},
		&domainModel.Domain{},
		&levelModel.Level{},
		&tuModel.TeachingUnit{},
		&tuModel.TeachingUnitProfessor{},
		&studentModel.Student{},
		&examModel.Exam{},
	)
	helpers.LogMigrations(
		err,
	)

	if err == nil {
		err = loadFixures()
	}

	return err
}

// loadFixures loads initial database values
func loadFixures() (err error) {
	var roleRepo = role.NewRepository(config.DB)
	var userRepo = user.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)

	// Create roles
	roleRepo.Create(&roleModel.Role{
		Name:        config.Env.RoleAdmin,
		Description: "Administrator role",
	})
	roleRepo.Create(&roleModel.Role{
		Name:        config.Env.RoleDefault,
		Description: "Default role",
	})
	roleAdmin, err := roleRepo.GetByName(config.Env.RoleAdmin)
	if err != nil {
		return
	}

	// Find admin user
	foundAdmin, _ := userRepo.GetByEmail(config.Env.UserAdminEmail)
	if foundAdmin != nil && foundAdmin.Email == config.Env.UserAdminEmail {
		return
	}

	// Create new admin
	userInfo, err := userRepo.CreateUserInfo(&userModel.UserInfo{
		Username: "Admin",
		Language: "en",
	})
	if err != nil {
		return
	}
	userMfa, err := userRepo.CreateUserMfa(&userModel.UserMfa{})
	if err != nil {
		return
	}
	tmpActivatedAt := time.Now()
	_, err = userRepo.Create(&userModel.User{
		Email:    config.Env.UserAdminEmail,
		Password: config.Env.UserAdminPassword,
		RoleID:   roleAdmin.ID,

		LoginMethod: constants.AuthLoginMethodDefault,
		IsActivated: true,
		ActivatedAt: &tmpActivatedAt,

		UserInfoID: userInfo.ID,
		UserMfaID:  userMfa.ID,
	})

	// Add admin permissions
	permissionRepo.UpdateByRoleID(
		roleAdmin.ID,
		constants.FeatureAdmin,
		permissionData.UpdatePermissionTableRequest{
			TableName: "*",
			Create:    true,
			Read:      true,
			Update:    true,
			Delete:    true,
		},
	)

	return
}
