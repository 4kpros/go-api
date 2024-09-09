package di

import (
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/docs"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRepositories() (
	authRepo *auth.AuthRepository,
	userRepo *user.UserRepository,
) {
	tmpAuth := auth.NewAuthRepositoryImpl(config.DB)
	tmpUser := user.NewUserRepositoryImpl(config.DB)
	authRepo = &tmpAuth
	userRepo = &tmpUser
	return
}

func InitServices(
	authRepo *auth.AuthRepository,
	userRepo *user.UserRepository,
) (
	authSer *auth.AuthService,
	userSer *user.UserService,
) {
	tmpAuth := auth.NewAuthServiceImpl(*authRepo)
	tmpUser := user.NewUserServiceImpl(*userRepo)
	authSer = &tmpAuth
	userSer = &tmpUser
	return
}

func InitControllers(
	authSer *auth.AuthService,
	userSer *user.UserService,
) (
	authContr *auth.AuthController,
	userContr *user.UserController,
) {
	tmpAuth := *auth.NewAuthController(*authSer)
	tmpUser := *user.NewUserController(*userSer)
	authContr = &tmpAuth
	userContr = &tmpUser
	return
}

func InitRouters(
	routerGroup *gin.RouterGroup,
	authContr *auth.AuthController,
	userContr *user.UserController,
) {
	// Add swagger
	docs.SwaggerInfo.BasePath = config.AppEnv.ApiGroup
	routerGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth.SetupEndpoints(routerGroup, authContr)
	user.SetupEndpoints(routerGroup, userContr)
}
