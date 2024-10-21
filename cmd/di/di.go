package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/admin/history"
	"api/services/admin/permission"
	"api/services/admin/role"
	"api/services/admin/user"
	"api/services/common/auth"
	"api/services/common/profile"
)

// InjectDependencies Inject all dependencies
func InjectDependencies() {
	var historyRepo = history.NewRepository(config.DB)
	var userRepo = user.NewRepository(config.DB)
	var roleRepo = role.NewRepository(config.DB)
	var permissionRepo = permission.NewRepository(config.DB)
	// History
	api.AllControllers.HistoryController = history.NewController(
		history.NewService(
			historyRepo,
		),
	)
	// Auth
	api.AllControllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
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
	// Profile
	api.AllControllers.ProfileController = profile.NewProfileController(
		profile.NewProfileService(
			userRepo,
		),
	)
}
