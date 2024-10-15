package di

import (
	"github.com/4kpros/go-api/cmd/api"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/admin"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/history"
	"github.com/4kpros/go-api/services/permission"
	"github.com/4kpros/go-api/services/profile"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
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
