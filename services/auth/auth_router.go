package auth

import (
	"github.com/danielgtaylor/huma/v2"
)

func SetupEndpoints(
	humaApi *huma.API,
	controller *AuthController,
) {
	const requireAuth = false

	// Sign in
	// router.Post(humaApi, "Login with email", "/signin-email", requireAuth, controller.SignInWithEmail)
	// router.Post(humaApi, "Login with phone number", "/signin-phone", requireAuth, controller.SignInWithEmail)
	// router.Post(humaApi, "Login with provider(Google, Facebook, ...)", "/signin-provider", requireAuth, controller.SignInWithProvider)

	// // Sign up
	// router.Post(humaApi, "Register with email", "/signup-email", requireAuth, controller.SignUpWithEmail)
	// router.Post(humaApi, "Register with phone number", "/signup-phone", requireAuth, controller.SignUpWithPhoneNumber)

	// // Activate account
	// router.Post(humaApi, "Activate account", "/activate", requireAuth, controller.ActivateAccount)

	// // Reset password
	// router.Post(humaApi, "Reset password with email - step 1", "/reset/init-email", requireAuth, controller.ResetPasswordEmailInit)
	// router.Post(humaApi, "Reset password with phone number- step 1", "/reset/init-phone", requireAuth, controller.ResetPasswordPhoneNumberInit)
	// router.Post(humaApi, "Reset password - step 2", "/reset/code", requireAuth, controller.ResetPasswordCode)
	// router.Post(humaApi, "Reset password - step 3", "/reset/new-password", requireAuth, controller.ResetPasswordNewPassword)

	// // Logout
	// router.Post(humaApi, "Log out", "/signout", true, controller.SignOut)
}
