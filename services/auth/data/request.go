package data

// Sign in
type SignInDevice struct {
	DeviceName string
}
type SignInWithEmailRequest struct {
	Email         string
	Password      string
	StayConnected bool
}
type SignInWithPhoneNumberRequest struct {
	PhoneNumber   int
	Password      string
	StayConnected bool
}
type SignInWithProviderRequest struct {
	Provider string
	Token    string
}
type SignInRequest struct {
	Email         string
	PhoneNumber   int
	Password      string
	StayConnected bool
}

// Sign up
type SignUpWithEmailRequest struct {
	Email    string
	Password string
}
type SignUpWithPhoneNumberRequest struct {
	PhoneNumber int
	Password    string
}
type SignUpRequest struct {
	Email       string
	PhoneNumber int
	Password    string
}

// Activate account
type ActivateAccountRequest struct {
	Token string
	Code  int
}

// Reset password
type ResetPasswordWithEmailInitRequest struct {
	Email string
}
type ResetPasswordWithPhoneNumberInitRequest struct {
	PhoneNumber int
}
type ResetPasswordInitRequest struct {
	Email       string
	PhoneNumber int
}
type ResetPasswordCodeRequest struct {
	Token string
	Code  int
}
type ResetPasswordNewPasswordRequest struct {
	Token       string
	NewPassword string
}

// Sign out
type SignOutRequest struct {
	Token string
}
