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
	"api/services/communication"
	"api/services/contact"
	"api/services/history"
	"api/services/school/common/director"
	"api/services/school/common/school"
	"api/services/school/common/year"
	"api/services/school/highschool/class"
	"api/services/school/highschool/pupil"
	"api/services/school/highschool/section"
	"api/services/school/highschool/specialty"
	"api/services/school/highschool/subject"
	"api/services/school/highschool/test"
	"api/services/school/university/department"
	"api/services/school/university/domain"
	"api/services/school/university/exam"
	"api/services/school/university/faculty"
	"api/services/school/university/level"
	"api/services/school/university/student"
	"api/services/school/university/tu"
	"api/services/user/auth"
	"api/services/user/permission"
	"api/services/user/profile"
	"api/services/user/role"
	"api/services/user/user"
)

type Controllers struct {
	// Others service
	CommunicationController     *communication.Controller
	ContactControllerController *contact.Controller
	HistoryController           *history.Controller

	// User service
	AuthController       *auth.Controller
	RoleController       *role.Controller
	PermissionController *permission.Controller
	UserController       *user.Controller
	ProfileController    *profile.Controller

	// School service
	YearController     *year.Controller
	SchoolController   *school.Controller
	DirectorController *director.Controller
	// Secondary
	SectionController   *section.Controller
	SpecialtyController *specialty.Controller
	ClassController     *class.Controller
	SubjectController   *subject.Controller
	PupilController     *pupil.Controller
	TestController      *test.Controller
	// Faculty
	FacultyController    *faculty.Controller
	DepartmentController *department.Controller
	DomainController     *domain.Controller
	LevelController      *level.Controller
	TUController         *tu.Controller
	ExamController       *exam.Controller
	StudentController    *student.Controller
}

var AllControllers = &Controllers{}

// Register all API endpoints
func registerEndpoints(humaApi *huma.API) {
	// Others service
	communication.RegisterEndpoints(humaApi, AllControllers.CommunicationController)
	contact.RegisterEndpoints(humaApi, AllControllers.ContactControllerController)
	history.RegisterEndpoints(humaApi, AllControllers.HistoryController)

	// User service
	auth.RegisterEndpoints(humaApi, AllControllers.AuthController)
	role.RegisterEndpoints(humaApi, AllControllers.RoleController)
	permission.RegisterEndpoints(humaApi, AllControllers.PermissionController)
	user.RegisterEndpoints(humaApi, AllControllers.UserController)
	profile.RegisterEndpoints(humaApi, AllControllers.ProfileController)

	// School service
	year.RegisterEndpoints(humaApi, AllControllers.YearController)
	school.RegisterEndpoints(humaApi, AllControllers.SchoolController)
	director.RegisterEndpoints(humaApi, AllControllers.DirectorController)
	// Highschool
	section.RegisterEndpoints(humaApi, AllControllers.SectionController)
	specialty.RegisterEndpoints(humaApi, AllControllers.SpecialtyController)
	class.RegisterEndpoints(humaApi, AllControllers.ClassController)
	subject.RegisterEndpoints(humaApi, AllControllers.SubjectController)
	pupil.RegisterEndpoints(humaApi, AllControllers.PupilController)
	test.RegisterEndpoints(humaApi, AllControllers.TestController)
	// University
	faculty.RegisterEndpoints(humaApi, AllControllers.FacultyController)
	department.RegisterEndpoints(humaApi, AllControllers.DepartmentController)
	domain.RegisterEndpoints(humaApi, AllControllers.DomainController)
	level.RegisterEndpoints(humaApi, AllControllers.LevelController)
	student.RegisterEndpoints(humaApi, AllControllers.StudentController)
	tu.RegisterEndpoints(humaApi, AllControllers.TUController)
	exam.RegisterEndpoints(humaApi, AllControllers.ExamController)
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
	// Custom hook to remove schema links
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
		middlewares.PermissionMiddleware(
			humaApi,
			AllControllers.RoleController.Service.Repository,
			AllControllers.PermissionController.Service.Repository,
		),
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
