package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/admin"
	"api/services/auth"
	"api/services/history"
	"api/services/permission"
	"api/services/profile"
	"api/services/role"
	"api/services/user"
)

// Inject all dependencies
func InjectDependencies() {
	var historyRepo = history.NewHistoryRepository(config.DB)
	var userRepo = user.NewUserRepository(config.DB)
	var roleRepo = role.NewRoleRepository(config.DB)
	var permissionRepo = permission.NewPermissionRepository(config.DB)
	// Auth
	api.Controllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
		),
	)
	// History
	api.Controllers.HistoryController = history.NewHistoryController(
		history.NewHistoryService(
			historyRepo,
		),
	)
	// Role
	api.Controllers.RoleController = role.NewRoleController(
		role.NewRoleService(
			roleRepo,
		),
	)
	// Permission
	api.Controllers.PermissionController = permission.NewPermissionController(
		permission.NewPermissionService(
			permissionRepo,
		),
	)
	// User
	api.Controllers.UserController = user.NewUserController(
		user.NewUserService(
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
	api.Controllers.AdminController = admin.NewAdminController(
		admin.NewAdminService(
			userRepo,
			roleRepo,
			permissionRepo,
		),
	)
}
