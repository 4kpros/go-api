package data

import "time"

type UpdateProfileEmailInitRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
}

type UpdateProfileEmailCheckCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1" example:""`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number" example:""`
}

type UpdateProfileEmailSetRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
}

type UpdateProfilePhoneNumberInitRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
}

type UpdateProfilePhoneNumberCheckCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1" example:""`
	Code  int    `json:"code" required:"true" doc:"Received Code on your phone number" example:""`
}

type UpdateProfilePhoneNumberSetRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
}

type UpdateProfilePasswordInitRequest struct {
}

type UpdateProfilePasswordCheckCodeRequest struct {
	Token string `json:"token" required:"true" minLength:"3" doc:"Received token on step 1" example:""`
	Code  int    `json:"code" required:"true" doc:"Received Code by email or phone number" example:""`
}
type UpdateProfilePasswordSetRequest struct {
	Token       string `json:"token" required:"true" minLength:"3" doc:"Received token on step 2" example:""`
	NewPassword string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
}

type UpdateProfileInfoRequest struct {
	UserName  string `json:"userName" required:"false" minLength:"2" maxLength:"30" doc:"User name" example:"meta_human"`
	FirstName string `json:"firstName" required:"false" minLength:"2" maxLength:"30" doc:"First name" example:"John"`
	LastName  string `json:"lastName" required:"false" minLength:"2" maxLength:"30" doc:"Last name" example:"Doe"`

	Birthday      time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string    `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string    `json:"address" required:"false" minLength:"2" maxLength:"30" doc:"Address" example:"No City"`
	Language      string    `json:"language" required:"false" minLength:"2" maxLength:"2" doc:"Language code with 2 letter" example:"en"`
	Image         string    `json:"image" required:"false" doc:"Thumbnail"`
}

type UpdateProfileMfaRequest struct {
	Method string `json:"method" required:"true" minLength:"2" maxLength:"30" doc:"Method to update MFA" example:"email"`
	Value  bool   `json:"value" required:"true" doc:"Method status" example:"false"`
}
