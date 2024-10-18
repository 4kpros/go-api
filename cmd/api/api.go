package api

import (
	"fmt"

	"api/common/constants"
	"api/config"
	"api/middlewares"
	"api/services/admin"
	"api/services/auth"
	"api/services/history"
	"api/services/permission"
	"api/services/profile"
	"api/services/role"
	"api/services/user"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"

	"github.com/gin-gonic/gin"
)

type APIControllers struct {
	AuthController       *auth.AuthController
	HistoryController    *history.HistoryController
	RoleController       *role.RoleController
	PermissionController *permission.PermissionController
	UserController       *user.UserController
	ProfileController    *profile.ProfileController
	AdminController      *admin.AdminController
}

var Controllers = &APIControllers{}

// Register all API endpoints
func registerEndpoints(humaApi *huma.API) {
	auth.RegisterEndpoints(humaApi, Controllers.AuthController)
	history.RegisterEndpoints(humaApi, Controllers.HistoryController)
	role.RegisterEndpoints(humaApi, Controllers.RoleController)
	permission.RegisterEndpoints(humaApi, Controllers.PermissionController)
	user.RegisterEndpoints(humaApi, Controllers.UserController)
	profile.RegisterEndpoints(humaApi, Controllers.ProfileController)
	admin.RegisterEndpoints(humaApi, Controllers.AdminController)
}

// Set up and start the API: set up API documentation,
// configure middlewares, and security measures.
func Start() {
	// Set up gin for your API
	gin.SetMode(config.Env.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	ginGroup := engine.Group(config.Env.ApiGroup)

	// OpenAPI documentation based on huma
	humaConfig := huma.DefaultConfig(constants.OPEN_API_TITLE, constants.OPEN_API_VERSION)
	// CUstom CreateHooks to remove $schema links
	humaConfig.CreateHooks = []func(huma.Config) huma.Config{
		func(c huma.Config) huma.Config {
			return c
		},
	}
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
	// Register middlewares
	humaApi.UseMiddleware(
		middlewares.HeadersMiddleware(humaApi),
		middlewares.CorsMiddleware(humaApi),
		middlewares.AuthMiddleware(humaApi),
		middlewares.PermissionMiddleware(humaApi, Controllers.PermissionController.Service.Repository),
	)

	// Register endpoints
	// Serve static files as favicon
	engine.Static("/static", constants.ASSET_PUBLIC_PATH)
	// Register endpoint for docs with support for custom template
	ginGroup.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", []byte(*config.OpenAPITemplates.Scalar))
	})
	registerEndpoints(&humaApi)

	// Start to listen
	formattedPort := fmt.Sprintf(":%d", config.Env.AppPort)
	err = engine.Run(formattedPort)
	if err != nil {
		panic(err)
	}
}
