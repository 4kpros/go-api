package di

import (
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/history"
	"github.com/4kpros/go-api/services/permission"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
	"github.com/danielgtaylor/huma/v2"
)

func InitRepositories() (
	authRepo *auth.AuthRepository,

	historyRepo *history.HistoryRepository,
	roleRepo *role.RoleRepository,
	permissionRepo *permission.PermissionRepository,
	userRepo *user.UserRepository,
) {
	authRepo = auth.NewAuthRepository(config.DB)

	historyRepo = history.NewHistoryRepository(config.DB)
	roleRepo = role.NewRoleRepository(config.DB)
	permissionRepo = permission.NewPermissionRepository(config.DB)
	userRepo = user.NewUserRepository(config.DB)
	return
}

func InitServices(
	authRepo *auth.AuthRepository,

	historyRepo *history.HistoryRepository,
	roleRepo *role.RoleRepository,
	permissionRepo *permission.PermissionRepository,
	userRepo *user.UserRepository,
) (
	authSvc *auth.AuthService,

	historySvc *history.HistoryService,
	roleSvc *role.RoleService,
	permissionSvc *permission.PermissionService,
	userSvc *user.UserService,
) {
	authSvc = auth.NewAuthService(*authRepo)

	historySvc = history.NewHistoryService(*historyRepo)
	roleSvc = role.NewRoleService(*roleRepo)
	permissionSvc = permission.NewPermissionService(*permissionRepo)
	userSvc = user.NewUserService(*userRepo)
	return
}

func InitControllers(
	authSvc *auth.AuthService,

	historySvc *history.HistoryService,
	roleSvc *role.RoleService,
	permissionSvc *permission.PermissionService,
	userSvc *user.UserService,
) (
	authCtrl *auth.AuthController,

	historyCtrl *history.HistoryController,
	roleCtrl *role.RoleController,
	permissionCtrl *permission.PermissionController,
	userCtrl *user.UserController,
) {
	authCtrl = auth.NewAuthController(*authSvc)

	historyCtrl = history.NewHistoryController(*historySvc)
	roleCtrl = role.NewRoleController(*roleSvc)
	permissionCtrl = permission.NewPermissionController(*permissionSvc)
	userCtrl = user.NewUserController(*userSvc)
	return
}

func InitRouters(
	humaApi *huma.API,

	authCtrl *auth.AuthController,

	historyCtrl *history.HistoryController,
	roleCtrl *role.RoleController,
	permissionCtrl *permission.PermissionController,
	userCtrl *user.UserController,
) {
	auth.SetupEndpoints(humaApi, authCtrl)

	history.SetupEndpoints(humaApi, historyCtrl)
	role.SetupEndpoints(humaApi, roleCtrl)
	permission.SetupEndpoints(humaApi, permissionCtrl)
	user.SetupEndpoints(humaApi, userCtrl)
}
