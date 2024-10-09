package di

import (
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/history"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
	"github.com/danielgtaylor/huma/v2"
)

func InitRepositories() (
	historyRepo *history.HistoryRepository,
	roleRepo *role.RoleRepository,
	authRepo *auth.AuthRepository,
	userRepo *user.UserRepository,
) {
	tmpHistory := history.NewHistoryRepositoryImpl(config.DB)
	historyRepo = &tmpHistory

	tmpRole := role.NewRoleRepositoryImpl(config.DB)
	roleRepo = &tmpRole

	tmpAuth := auth.NewAuthRepositoryImpl(config.DB)
	authRepo = &tmpAuth

	tmpUser := user.NewUserRepositoryImpl(config.DB)
	userRepo = &tmpUser
	return
}

func InitServices(
	historyRepo *history.HistoryRepository,
	roleRepo *role.RoleRepository,
	authRepo *auth.AuthRepository,
	userRepo *user.UserRepository,
) (
	historySvc *history.HistoryService,
	roleSvc *role.RoleService,
	authSvc *auth.AuthService,
	userSvc *user.UserService,
) {
	tmpHistory := history.NewHistoryServiceImpl(*historyRepo)
	historySvc = &tmpHistory

	tmpRole := role.NewRoleServiceImpl(*roleRepo)
	roleSvc = &tmpRole

	tmpAuth := auth.NewAuthServiceImpl(*authRepo)
	authSvc = &tmpAuth

	tmpUser := user.NewUserServiceImpl(*userRepo)
	userSvc = &tmpUser
	return
}

func InitControllers(
	historySvc *history.HistoryService,
	roleSvc *role.RoleService,
	authSvc *auth.AuthService,
	userSvc *user.UserService,
) (
	historyCtrl *history.HistoryController,
	roleCtrl *role.RoleController,
	authCtrl *auth.AuthController,
	userCtrl *user.UserController,
) {
	historyCtrl = history.NewHistoryController(*historySvc)
	roleCtrl = role.NewRoleController(*roleSvc)
	authCtrl = auth.NewAuthController(*authSvc)
	userCtrl = user.NewUserController(*userSvc)
	return
}

func InitRouters(
	humaApi *huma.API,
	historyCtrl *history.HistoryController,
	roleCtrl *role.RoleController,
	authCtrl *auth.AuthController,
	userCtrl *user.UserController,
) {
	history.SetupEndpoints(humaApi, historyCtrl)
	role.SetupEndpoints(humaApi, roleCtrl)
	auth.SetupEndpoints(humaApi, authCtrl)
	user.SetupEndpoints(humaApi, userCtrl)
}
