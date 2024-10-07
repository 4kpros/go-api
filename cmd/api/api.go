package api

import (
	"fmt"

	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/di"

	"github.com/gin-gonic/gin"
)

func Start() {
	// Setup gin for your API
	gin.SetMode(config.AppEnv.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.Use(middleware.ErrorsHandler())
	apiGroup := engine.Group(config.AppEnv.ApiGroup)

	// Inject Dependencies
	authRepo, userRepo :=
		di.InitRepositories() // Repositories
	authSer, userSer :=
		di.InitServices(
			authRepo, userRepo,
		) // Services
	authContr, userContr :=
		di.InitControllers(
			authSer, userSer,
		) // Controllers
	di.InitRouters(
		apiGroup, authContr, userContr,
	) // Routers

	// Run gin
	formattedPort := fmt.Sprintf(":%d", config.AppEnv.ApiPort)
	engine.Run(formattedPort)
}
