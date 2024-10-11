package data

import "time"

// Login
type SignInResponse struct {
	AccessToken string     `json:"accessToken" required:"false" doc:"Access token" example:""`
	Expires     *time.Time `json:"expires" required:"false" doc:"Token expiration date time" example:""`
}

// Register
type SignUpResponse struct {
	Message string `json:"message" required:"false" doc:"Message" example:""`
}

// Activate account
type ActivateAccountResponse struct {
	ActivatedAt *time.Time `json:"activatedAt" required:"false" doc:"Activation date" example:""`
}

// Forgot password
type ForgotPasswordInitResponse struct {
	Token string `json:"token" required:"false" doc:"Token" example:""`
}
type ForgotPasswordCodeResponse struct {
	Token string `json:"token" required:"false" doc:"Token" example:""`
}
type ForgotPasswordNewPasswordResponse struct {
	Message string `json:"message" required:"false" doc:"Message" example:""`
}

// Logout
type SignOutResponse struct {
	Message string `json:"message" required:"false" doc:"Message" example:""`
}
