package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/history"
	"api/services/school/common/school"
	"api/services/school/common/year"
	"api/services/school/university/department"
	"api/services/school/university/domain"
	"api/services/school/university/faculty"
	"api/services/user/auth"
	"api/services/user/permission"
	"api/services/user/profile"
	"api/services/user/role"
	"api/services/user/user"
)

// InjectDependencies Inject all dependencies
func InjectDependencies() {
	// History
	var historyRepo = history.NewRepository(config.DB)
	api.AllControllers.HistoryController = history.NewController(
		history.NewService(
			historyRepo,
		),
	)

	// User
	var userRepo = user.NewRepository(config.DB)
	var roleRepo = role.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)
	api.AllControllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
		),
	)
	api.AllControllers.RoleController = role.NewController(
		role.NewService(
			roleRepo,
		),
	)
	api.AllControllers.PermissionController = permission.NewController(
		permission.NewService(
			permissionRepo,
		),
	)
	api.AllControllers.UserController = user.NewController(
		user.NewService(
			userRepo,
		),
	)
	api.AllControllers.ProfileController = profile.NewController(
		profile.NewService(
			userRepo,
		),
	)

	// School
	var yearRepo = year.NewRepository(config.DB)
	var schoolRepo = school.NewRepository(config.DB)
	api.AllControllers.YearController = year.NewController(
		year.NewService(
			yearRepo,
		),
	)
	api.AllControllers.SchoolController = school.NewController(
		school.NewService(
			schoolRepo,
		),
	)
	api.AllControllers.SchoolController = school.NewController(
		school.NewService(
			schoolRepo,
		),
	)

	// Secondary

	// University
	var facultyRepo = faculty.NewRepository(config.DB)
	var departmentRepo = department.NewRepository(config.DB)
	var domainRepo = domain.NewRepository(config.DB)
	api.AllControllers.FacultyController = faculty.NewController(
		faculty.NewService(
			facultyRepo,
		),
	)
	api.AllControllers.DepartmentController = department.NewController(
		department.NewService(
			departmentRepo,
		),
	)
	api.AllControllers.DomainController = domain.NewController(
		domain.NewService(
			domainRepo,
		),
	)
}
