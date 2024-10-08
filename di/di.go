package di

import (
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
	"github.com/danielgtaylor/huma/v2"
)

func InitRepositories() (
	roleRepo *role.RoleRepository,
	authRepo *auth.AuthRepository,
	userRepo *user.UserRepository,
) {
	tmpRole := role.NewRoleRepositoryImpl(config.DB)
	tmpAuth := auth.NewAuthRepositoryImpl(config.DB)
	tmpUser := user.NewUserRepositoryImpl(config.DB)
	roleRepo = &tmpRole
	authRepo = &tmpAuth
	userRepo = &tmpUser
	return
}

func InitServices(
	roleRepo *role.RoleRepository,
	authRepo *auth.AuthRepository,
	userRepo *user.UserRepository,
) (
	roleSer *role.RoleService,
	authSer *auth.AuthService,
	userSer *user.UserService,
) {
	tmpRole := role.NewRoleServiceImpl(*roleRepo)
	tmpAuth := auth.NewAuthServiceImpl(*authRepo)
	tmpUser := user.NewUserServiceImpl(*userRepo)
	roleSer = &tmpRole
	authSer = &tmpAuth
	userSer = &tmpUser
	return
}

func InitControllers(
	roleSer *role.RoleService,
	authSer *auth.AuthService,
	userSer *user.UserService,
) (
	roleContr *role.RoleController,
	authContr *auth.AuthController,
	userContr *user.UserController,
) {
	tmpRole := *role.NewRoleController(*roleSer)
	tmpAuth := *auth.NewAuthController(*authSer)
	tmpUser := *user.NewUserController(*userSer)
	roleContr = &tmpRole
	authContr = &tmpAuth
	userContr = &tmpUser
	return
}

func InitRouters(
	humaApi *huma.API,
	//
	roleContr *role.RoleController,
	authContr *auth.AuthController,
	userContr *user.UserController,
) {
	role.SetupEndpoints(humaApi, roleContr)
	auth.SetupEndpoints(humaApi, authContr)
	user.SetupEndpoints(humaApi, userContr)
	return
}
