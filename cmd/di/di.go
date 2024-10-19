package di

import (
	"api/cmd/api"
	"api/config"
	"api/services/admin"
	history2 "api/services/admin/history"
	permission2 "api/services/admin/permission"
	role2 "api/services/admin/role"
	user2 "api/services/admin/user"
	auth2 "api/services/common/auth"
	profile2 "api/services/common/profile"
)

// Inject all dependencies
func InjectDependencies() {
	var historyRepo = history2.NewHistoryRepository(config.DB)
	var userRepo = user2.NewUserRepository(config.DB)
	var roleRepo = role2.NewRoleRepository(config.DB)
	var permissionRepo = permission2.NewPermissionRepository(config.DB)
	// Auth
	api.Controllers.AuthController = auth2.NewAuthController(
		auth2.NewAuthService(
			userRepo,
		),
	)
	// History
	api.Controllers.HistoryController = history2.NewHistoryController(
		history2.NewHistoryService(
			historyRepo,
		),
	)
	// Role
	api.Controllers.RoleController = role2.NewRoleController(
		role2.NewRoleService(
			roleRepo,
		),
	)
	// Permission
	api.Controllers.PermissionController = permission2.NewPermissionController(
		permission2.NewPermissionService(
			permissionRepo,
		),
	)
	// User
	api.Controllers.UserController = user2.NewUserController(
		user2.NewUserService(
			userRepo,
		),
	)
	// Profile
	api.Controllers.ProfileController = profile2.NewProfileController(
		profile2.NewProfileService(
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
