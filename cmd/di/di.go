package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/admin"
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
	var permissionRepo = permission.NewPermissionRepository(config.DB)
	// Auth
	api.Controllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
		),
	)
	// History
	api.Controllers.HistoryController = history.NewController(
		history.NewService(
			historyRepo,
		),
	)
	// Role
	api.Controllers.RoleController = role.NewController(
		role.NewService(
			roleRepo,
		),
	)
	// Permission
	api.Controllers.PermissionController = permission.NewController(
		permission.NewService(
			permissionRepo,
		),
	)
	// User
	api.Controllers.UserController = user.NewController(
		user.NewService(
			userRepo,
		),
	)
	// Profile
	api.Controllers.ProfileController = profile.NewProfileController(
		profile.NewProfileService(
			userRepo,
		),
	)
	// Admin
	api.Controllers.AdminController = admin.NewController(
		admin.NewAdminService(
			userRepo,
			roleRepo,
			permissionRepo,
		),
	)
}
