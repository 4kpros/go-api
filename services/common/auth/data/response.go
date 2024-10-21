package data

import "time"

// Login
type LoginResponse struct {
	AccessToken          string     `json:"accessToken" required:"false" doc:"Access token"`
	Expires              *time.Time `json:"expires" required:"false" doc:"Access token expiration date time"`
	ActivateAccountToken string     `json:"activateAccountToken" required:"false" doc:"Token to account token"`
}

// Register
type RegisterResponse struct {
	ActivateAccountToken string `json:"activateAccountToken" required:"false" doc:"Token to account token"`
	Message              string `json:"message" required:"false" doc:"Response message"`
}

// Activate account
type ActivateAccountResponse struct {
	ActivatedAt *time.Time `json:"activatedAt" required:"false" doc:"Account activation date time"`
}

// Forgot password
type ForgotPasswordInitResponse struct {
	Token string `json:"token" required:"false" doc:"Token used to validate code"`
}
type ForgotPasswordCodeResponse struct {
	Token string `json:"token" required:"false" doc:"Token used to set new password"`
}
type ForgotPasswordNewPasswordResponse struct {
	Message string `json:"message" required:"false" doc:"Response message"`
}

// Logout
type LogoutResponse struct {
	Message string `json:"message" required:"false" doc:"Response message"`
}
