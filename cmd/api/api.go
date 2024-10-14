package api

import (
	"fmt"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/middlewares"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/services/auth"
	"github.com/4kpros/go-api/services/history"
	"github.com/4kpros/go-api/services/permission"
	"github.com/4kpros/go-api/services/profile"
	"github.com/4kpros/go-api/services/role"
	"github.com/4kpros/go-api/services/user"
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
	engine.SetTrustedProxies([]string{"127.0.0.1"})
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
		middlewares.SecureHeadersMiddleware(humaApi),
		middlewares.AuthMiddleware(humaApi),
	)

	// Register endpoints
	// Register endpoint for docs with support for custom template
	ginGroup.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", []byte(*config.OpenAPITemplates.Scalar))
	})
	registerEndpoints(&humaApi)

	// Start to listen
	formattedPort := fmt.Sprintf(":%d", config.Env.AppPort)
	engine.Run(formattedPort)
}
