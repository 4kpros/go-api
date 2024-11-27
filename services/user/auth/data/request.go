package data

// Login
type LoginDevice struct {
	Platform   string `json:"platform" required:"true" minLength:"2" maxLength:"30" doc:"Platform name" example:"Android"`
	DeviceName string `json:"deviceName" required:"true" minLength:"2" maxLength:"50" doc:"Device name" example:"Google Pixel 5"`
	App        string `json:"app" required:"true" minLength:"2" maxLength:"50" doc:"Application used to login" example:"Chrome"`
}
type LoginWithEmailRequest struct {
	Email         string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected" example:"false"`
}
type LoginWithPhoneNumberRequest struct {
	PhoneNumber   uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected" example:"false"`
}
type LoginWithProviderRequest struct {
	Provider string `json:"provider" required:"true" doc:"Provider" minLength:"2" maxLength:"30" example:"google"`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3" example:""`
}
type LoginRequest struct {
	Email         string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber   uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected" example:"false"`
}

// Register
type RegisterWithEmailRequest struct {
	Email    string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	Password string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}
type RegisterWithPhoneNumberRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	Password    string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}
type RegisterRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	Password    string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}

// Activate account
type ActivateAccountRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token" example:""`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number" example:""`
}

// Forgot password
type ForgotPasswordWithEmailInitRequest struct {
	Email string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
}
type ForgotPasswordWithPhoneNumberInitRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
}
type ForgotPasswordInitRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
}
type ForgotPasswordCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1" example:""`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number" example:""`
}
type ForgotPasswordNewPasswordRequest struct {
	Token       string `json:"token" required:"true" minLength:"3" doc:"Received token on step 2" example:""`
	NewPassword string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}
