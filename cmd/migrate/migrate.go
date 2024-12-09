package migrate

import (
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
	permissionModel "api/services/user/permission/model"
	roleModel "api/services/user/role/model"
	userModel "api/services/user/user/model"
)

// Start Loads and applies all migrations.
func Start() error {
	err := config.DB.AutoMigrate(
		// History
		&historyModel.History{},

		// User
		&permissionModel.Permission{},
		&userModel.User{},
		&roleModel.Role{},
		&userModel.UserMfa{},
		&userModel.UserInfo{},

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

	return err
}
