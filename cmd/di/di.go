package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/history"
	"api/services/school/common/school"
	"api/services/user/auth"
	"api/services/user/permission"
	"api/services/user/profile"
	"api/services/user/role"
	"api/services/user/user"
)

// InjectDependencies Inject all dependencies
func InjectDependencies() {
	var historyRepo = history.NewRepository(config.DB)
	var userRepo = user.NewRepository(config.DB)
	var roleRepo = role.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)
	var schoolRepo = school.NewRepository(config.DB)

	// Auth
	api.AllControllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
		),
	)
	// Profile
	api.AllControllers.ProfileController = profile.NewController(
		profile.NewService(
			userRepo,
		),
	)

	// History
	api.AllControllers.HistoryController = history.NewController(
		history.NewService(
			historyRepo,
		),
	)
	// Role
	api.AllControllers.RoleController = role.NewController(
		role.NewService(
			roleRepo,
		),
	)
	// Permission
	api.AllControllers.PermissionController = permission.NewController(
		permission.NewService(
			permissionRepo,
		),
	)
	// User
	api.AllControllers.UserController = user.NewController(
		user.NewService(
			userRepo,
		),
	)
	// School
	api.AllControllers.SchoolController = school.NewController(
		school.NewService(
			schoolRepo,
		),
	)

	// School Director
	api.AllControllers.SchoolController = school.NewController(
		school.NewService(
			schoolRepo,
		),
	)
}
