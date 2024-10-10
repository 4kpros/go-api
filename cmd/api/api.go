package api

import (
	"fmt"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/middleware"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/di"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"

	"github.com/gin-gonic/gin"
)

// Used to setup and start API: dependency injection, setup API doc, middlewares and security
func Start() {
	// Inject Dependencies step 1
	// Repositories
	authRepo, historyRepo, roleRepo, permissionRepo, userRepo :=
		di.InitRepositories()
	// Services
	authSvc, historySvc, roleSvc, permissionSvc, userSvc :=
		di.InitServices(
			authRepo, historyRepo, roleRepo, permissionRepo, userRepo,
		)
	// Controllers
	authCtrl, historyCtrl, roleCtrl, permissionCtrl, userCtrl :=
		di.InitControllers(
			authSvc, historySvc, roleSvc, permissionSvc, userSvc,
		)

	// Setup gin for your API
	gin.SetMode(config.Env.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	ginGroup := engine.Group(config.Env.ApiGroup)
	// OpenAPI documentation based on huma
	humaConfig := huma.DefaultConfig(constants.OPEN_API_TITLE, constants.OPEN_API_VERSION)
	humaConfig.DocsPath = ""
	humaConfig.Servers = []*huma.Server{
		{URL: config.Env.ApiGroup},
	}
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		constants.SECURITY_AUTH_NAME: {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
			Description:  "Bearer token used to access some resources",
		},
	}
	humaConfig.Info.Description = constants.OPEN_API_DESCRIPTION
	humaApi := humagin.NewWithGroup(engine, ginGroup, humaConfig)
	humaApi.UseMiddleware(
		middleware.SecureHeadersMiddleware(humaApi),
		middleware.RateLimitMiddleware(humaApi),
		middleware.AuthMiddleware(humaApi),
	)
	ginGroup.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", []byte(config.OpenAPITemplates.Scalar))
	})

	// Inject Dependencies step 2
	// Routers
	di.InitRouters(
		&humaApi, authCtrl, historyCtrl, roleCtrl, permissionCtrl, userCtrl,
	)

	// Start to listen
	formattedPort := fmt.Sprintf(":%d", config.Env.ApiPort)
	engine.Run(formattedPort)
}
