package user

import (
	"github.com/danielgtaylor/huma/v2"
)

func SetupEndpoints(
	humaApi *huma.API,
	controller *UserController,
) {
	const requireAuth = true

	// Create
	// router.Post(humaApi, "Create user by email", "/email", requireAuth, controller.CreateWithEmail)
	// router.Post(humaApi, "Create user by phone number", "/phone", requireAuth, controller.CreateWithPhoneNumber)

	// // Update
	// router.Put(humaApi, "Update user", "/:id", requireAuth, controller.UpdateUser)
	// router.Put(humaApi, "Update user info", "/info/:id", requireAuth, controller.UpdateUserInfo)

	// // Get
	// router.Get(humaApi, "Get user", "/:id", requireAuth, controller.FindById)
	// router.Get(humaApi, "Get user info", "/info/:id", requireAuth, controller.FindUserInfoById)
	// router.Get(humaApi, "Get all users", "/", requireAuth, controller.FindAll)

	// // Delete
	// router.Delete(humaApi, "Delete user", "/:id", requireAuth, controller.Delete)
}
