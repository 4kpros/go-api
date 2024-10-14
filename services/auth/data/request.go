package data

// Login
type SignInDevice struct {
	Platform   string `json:"platform" required:"true" minLength:"2" maxLength:"30" doc:"Platform name" example:"Android"`
	DeviceName string `json:"deviceName" required:"true" minLength:"2" maxLength:"50" doc:"Device name" example:"Google Pixel 5"`
	App        string `json:"app" required:"true" minLength:"2" maxLength:"50" doc:"Application used to login" example:"Chrome"`
}
type SignInWithEmailRequest struct {
	Email         string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected" example:"false"`
}
type SignInWithPhoneNumberRequest struct {
	PhoneNumber   uint64 `json:"phoneNumber" required:"true" minLength:"10" maxLength:"50" doc:"Phone number" example:"237690909090"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected" example:"false"`
}
type SignInWithProviderRequest struct {
	Provider string `json:"provider" required:"true" doc:"Provider" minLength:"2" maxLength:"30" example:"google"`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3" maxLength:"30" example:""`
}
type SignInRequest struct {
	Email         string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber   uint64 `json:"phoneNumber" required:"true" minLength:"10" maxLength:"50" doc:"Phone number" example:"237690909090"`
	Password      string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	StayConnected bool   `json:"stayConnected" required:"false" doc:"Stay connected" example:"false"`
}

// Register
type SignUpWithEmailRequest struct {
	Email    string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	Password string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}
type SignUpWithPhoneNumberRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minLength:"10" maxLength:"50" doc:"Phone number" example:"237690909090"`
	Password    string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}
type SignUpRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minLength:"10" maxLength:"50" doc:"Phone number" example:"237690909090"`
	Password    string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}

// Activate account
type ActivateAccountRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on sign in or sign up" example:""`
	Code  int    `json:"code" required:"true" minLength:"5" maxLength:"8" doc:"Received Code by email or phone number" example:""`
}

// Forgot password
type ForgotPasswordWithEmailInitRequest struct {
	Email string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
}
type ForgotPasswordWithPhoneNumberInitRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minLength:"10" maxLength:"50" doc:"Phone number" example:"237690909090"`
}
type ForgotPasswordInitRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" minLength:"10" maxLength:"50" doc:"Phone number" example:"237690909090"`
}
type ForgotPasswordCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1" example:""`
	Code  int    `json:"code" required:"true" minLength:"5" maxLength:"8" doc:"Received Code by email or phone number" example:""`
}
type ForgotPasswordNewPasswordRequest struct {
	Token       string `json:"token" required:"true" minLength:"3" doc:"Received token on step 2" example:""`
	NewPassword string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}

// Sign out
type SignOutRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Valid access token returned wen user sign in" example:""`
}
