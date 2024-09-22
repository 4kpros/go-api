package user

import (
	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/router"
	"github.com/gin-gonic/gin"
)

func SetupEndpoints(routerGroup *gin.RouterGroup, controller *UserController) {

	group := routerGroup.Group("/users")
	const requireAuth = true
	allowedRoles := []string{constants.ROLE_ADMIN, "customer"}

	// Create new user
	router.POST(group, "/email", controller.CreateWithEmail, requireAuth, allowedRoles)
	router.POST(group, "/phone", controller.CreateWithPhoneNumber, requireAuth, allowedRoles)

	// Update user and user info
	router.PUT(group, "/:id", controller.UpdateUser, requireAuth, allowedRoles)
	router.PUT(group, "/info/:id", controller.UpdateUserInfo, requireAuth, allowedRoles)

	// Delete user
	router.DELETE(group, "/:id", controller.Delete, requireAuth, allowedRoles)

	// Get user and user info
	router.GET(group, "/:id", controller.FindById, requireAuth, allowedRoles)
	router.GET(group, "/info/:id", controller.FindUserInfoById, requireAuth, allowedRoles)
	router.GET(group, "/", controller.FindAll, requireAuth, allowedRoles)
}
