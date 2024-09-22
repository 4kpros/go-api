package auth

import (
	"github.com/4kpros/go-api/common/router"
	"github.com/gin-gonic/gin"
)

func SetupEndpoints(
	routerGroup *gin.RouterGroup,
	controller *AuthController,
) {
	group := routerGroup.Group("/auth")
	const requireAuth = false
	allowedRoles := []string{} // Emty list => allow all roles

	// Sign in
	router.POST(group, "/signin-email", controller.SignInWithEmail, requireAuth, allowedRoles)
	router.POST(group, "/signin-phone", controller.SignInWithEmail, requireAuth, allowedRoles)
	router.POST(group, "/signin-provider", controller.SignInWithProvider, requireAuth, allowedRoles)

	// Sign up
	router.POST(group, "/signup-email", controller.SignUpWithEmail, requireAuth, allowedRoles)
	router.POST(group, "/signup-phone", controller.SignUpWithPhoneNumber, requireAuth, allowedRoles)

	// Activate account
	router.POST(group, "/activate", controller.ActivateAccount, requireAuth, allowedRoles)

	// Reset password
	router.POST(group, "/reset/init-email", controller.ResetPasswordEmailInit, requireAuth, allowedRoles)
	router.POST(group, "/reset/init-phone", controller.ResetPasswordPhoneNumberInit, requireAuth, allowedRoles)
	router.POST(group, "/reset/code", controller.ResetPasswordCode, requireAuth, allowedRoles)
	router.POST(group, "/reset/new-password", controller.ResetPasswordNewPassword, requireAuth, allowedRoles)

	// Logout
	router.POST(group, "/signout", controller.SignOut, true, allowedRoles)
}
