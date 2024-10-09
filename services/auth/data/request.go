package data

// Sign in
type SignInDevice struct {
	DeviceName string `json:"deviceName" required:"true" doc:"Device name" minLength:"2" maxLength:"30" example:"Android - Pixel 5"`
}
type SignInWithEmailRequest struct {
	Email         string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	Password      string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}
type SignInWithPhoneNumberRequest struct {
	PhoneNumber   int    `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Password      string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}
type SignInWithProviderRequest struct {
	Provider string `json:"provider" required:"true" doc:"Provider" minLength:"2" maxLength:"30" example:"google"`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3" maxLength:"30" example:""`
}
type SignInRequest struct {
	Email         string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	PhoneNumber   int    `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Password      string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected"`
}

// Sign up
type SignUpWithEmailRequest struct {
	Email    string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	Password string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
}
type SignUpWithPhoneNumberRequest struct {
	PhoneNumber int    `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Password    string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
}
type SignUpRequest struct {
	Email       string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	PhoneNumber int    `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Password    string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
}

// Activate account
type ActivateAccountRequest struct {
	Token string `json:"token" required:"true" doc:"Token" minLength:"3" maxLength:"30" example:""`
	Code  int    `json:"code" required:"true" doc:"Email" example:"37044"`
}

// Reset password
type ResetPasswordWithEmailInitRequest struct {
	Email string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
}
type ResetPasswordWithPhoneNumberInitRequest struct {
	PhoneNumber int `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
}
type ResetPasswordInitRequest struct {
	Email       string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	PhoneNumber int    `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
}
type ResetPasswordCodeRequest struct {
	Token string `json:"token" required:"true" doc:"Token" minLength:"3" maxLength:"30" example:""`
	Code  int    `json:"code" required:"true" doc:"Email" example:"37044"`
}
type ResetPasswordNewPasswordRequest struct {
	Token       string `json:"token" required:"true" doc:"Token" minLength:"3" maxLength:"30" example:""`
	NewPassword string `json:"password" required:"true" doc:"Base64 encoded password" minLength:"8" maxLength:"30" example:""`
}

// Sign out
type SignOutRequest struct {
	Token string `json:"token" required:"true" doc:"Token" minLength:"3" maxLength:"30" example:""`
}