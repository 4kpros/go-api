package api

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"api/common/constants"
	"api/config"
	"api/middlewares"
	"api/services/admin/history"
	"api/services/admin/permission"
	"api/services/admin/role"
	"api/services/admin/school"
	"api/services/admin/user"
	"api/services/common/auth"
	"api/services/common/profile"
)

type Controllers struct {
	AuthController       *auth.Controller
	HistoryController    *history.Controller
	RoleController       *role.Controller
	PermissionController *permission.Controller
	UserController       *user.Controller
	ProfileController    *profile.Controller
	SchoolController     *school.Controller
}

var AllControllers = &Controllers{}

// Register all API endpoints
func registerEndpoints(humaApi *huma.API) {
	auth.RegisterEndpoints(humaApi, AllControllers.AuthController)
	history.RegisterEndpoints(humaApi, AllControllers.HistoryController)
	role.RegisterEndpoints(humaApi, AllControllers.RoleController)
	permission.RegisterEndpoints(humaApi, AllControllers.PermissionController)
	user.RegisterEndpoints(humaApi, AllControllers.UserController)
	profile.RegisterEndpoints(humaApi, AllControllers.ProfileController)
	school.RegisterEndpoints(humaApi, AllControllers.SchoolController)
}

// Start Set up and start the API: set up API documentation,
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
	humaConfig := huma.DefaultConfig(constants.OpenApiTitle, constants.OpenApiVersion)
	// Custom CreateHooks to remove $schema links
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
		constants.SecurityAuthName: {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
			Description:  "Bearer token used to access some resources",
		},
	}
	humaConfig.Info.Description = constants.OpenApiDescription
	humaApi := humagin.NewWithGroup(engine, ginGroup, humaConfig)
	// Register middlewares
	humaApi.UseMiddleware(
		middlewares.HeadersMiddleware(humaApi),
		middlewares.CorsMiddleware(humaApi),
		middlewares.AuthMiddleware(humaApi),
		middlewares.PermissionMiddleware(humaApi, AllControllers.PermissionController.Service.Repository),
	)

	// Register endpoints
	// Serve static files as favicon
	engine.StaticFS("/assets", http.Dir(constants.AssetAppPath))
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
