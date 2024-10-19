package api

import (
	history2 "api/services/admin/history"
	permission2 "api/services/admin/permission"
	role2 "api/services/admin/role"
	user2 "api/services/admin/user"
	auth2 "api/services/common/auth"
	profile2 "api/services/common/profile"
	"fmt"
	"net/http"

	"api/common/constants"
	"api/config"
	"api/middlewares"
	"api/services/admin"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"

	"github.com/gin-gonic/gin"
)

type APIControllers struct {
	AuthController       *auth2.Controller
	HistoryController    *history2.HistoryController
	RoleController       *role2.RoleController
	PermissionController *permission2.PermissionController
	UserController       *user2.UserController
	ProfileController    *profile2.ProfileController
	AdminController      *admin.AdminController
}

var Controllers = &APIControllers{}

// Register all API endpoints
func registerEndpoints(humaApi *huma.API) {
	auth2.RegisterEndpoints(humaApi, Controllers.AuthController)
	history2.RegisterEndpoints(humaApi, Controllers.HistoryController)
	role2.RegisterEndpoints(humaApi, Controllers.RoleController)
	permission2.RegisterEndpoints(humaApi, Controllers.PermissionController)
	user2.RegisterEndpoints(humaApi, Controllers.UserController)
	profile2.RegisterEndpoints(humaApi, Controllers.ProfileController)
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
	engine.StaticFS("/assets", http.Dir(constants.ASSET_APP_PATH))
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
