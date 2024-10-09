package data

import "time"

// Sign in
type SignInResponse struct {
	AccessToken string    `json:"accessToken"`
	Expires     time.Time `json:"expires"`
}

// Sign up
type SignUpResponse struct {
	Message string `json:"message"`
}

// Activate account
type ActivateAccountResponse struct {
	ActivatedAt time.Time `json:"activatedAt"`
}

// Reset password
type ResetPasswordInitResponse struct {
	Token string `json:"token"`
}
type ResetPasswordCodeResponse struct {
	Token string `json:"token"`
}
type ResetPasswordNewPasswordResponse struct {
	Message string `json:"message"`
}

// Sign out
type SignOutResponse struct {
	Message string `json:"message"`
}
