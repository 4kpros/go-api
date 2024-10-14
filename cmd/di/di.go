package di

import (
	"github.com/4kpros/go-api/cmd/api"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/history"
	"github.com/4kpros/go-api/services/permission"
	"github.com/4kpros/go-api/services/profile"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
)

// Inject all dependencies
func InjectDependencies() {
	var userRepo = user.NewUserRepository(config.DB)
	// Auth
	api.Controllers.AuthController = auth.NewAuthController(
		auth.NewAuthService(
			userRepo,
		),
	)
	// History
	api.Controllers.HistoryController = history.NewHistoryController(
		history.NewHistoryService(
			history.NewHistoryRepository(config.DB),
		),
	)
	// Role
	api.Controllers.RoleController = role.NewRoleController(
		role.NewRoleService(
			role.NewRoleRepository(config.DB),
		),
	)
	// Permission
	api.Controllers.PermissionController = permission.NewPermissionController(
		permission.NewPermissionService(
			permission.NewPermissionRepository(config.DB),
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
}
