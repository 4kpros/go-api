package api

import (
	"fmt"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/di"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"

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
	ginGroup := engine.Group(config.AppEnv.ApiGroup)

	// OpenAPI documentation
	humaConfig := huma.DefaultConfig(constants.OPEN_API_TITLE, constants.OPEN_API_VERSION)
	humaConfig.DocsPath = ""
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	humaConfig.Info.Description = constants.OPEN_API_DESCRIPTION
	humaApi := humagin.NewWithGroup(nil, ginGroup, humaConfig)
	// humaApi.UseMiddleware(middleware.ErrorsMiddleware, middleware.SecureApiMiddleware)
	ginGroup.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", []byte(config.AppTemplate.Scalar))
	})

	// Inject Dependencies
	roleRepo, authRepo, userRepo :=
		di.InitRepositories() // Repositories
	roleSvc, authSvc, userSvc :=
		di.InitServices(
			roleRepo, authRepo, userRepo,
		) // Services
	roleCtrl, authCtrl, userCtrl :=
		di.InitControllers(
			roleSvc, authSvc, userSvc,
		) // Controllers
	di.InitRouters(
		&humaApi, roleCtrl, authCtrl, userCtrl,
	) // Routers

	// Run gin
	formattedPort := fmt.Sprintf(":%d", config.AppEnv.ApiPort)
	engine.Run(formattedPort)
}
