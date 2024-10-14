package data

import (
	"time"

	"github.com/4kpros/go-api/common/types"
)

type UserResponse struct {
	types.BaseGormModelResponse
	Email       string `json:"email" required:"false" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`
	RoleId      int64  `json:"roleId" required:"false" doc:"Role id" example:"1"`

	SignInMethod   string     `json:"signInMethod" required:"false" doc:"How the user should login ? with email, phone number or external provider?"`
	Provider       string     `json:"provider" required:"false" doc:"Provider name"`
	ProviderUserId string     `json:"providerUserId" required:"false" doc:"User id from the provider"`
	IsActivated    bool       `json:"isActivated" required:"false" doc:"Is user account activated ?"`
	ActivatedAt    *time.Time `json:"activatedAt" required:"false" doc:"Activation date time"`

	UserInfo UserInfoResponse `json:"info" required:"false" doc:"Additional user info(e.g. address, first name, last name, ...)" `
	UserMfa  UserMfaResponse  `json:"mfa" required:"false" doc:"Multiple factor authenticator enabled by the user"`
}

type UserInfoResponse struct {
	UserName  string `json:"userName" required:"false" doc:"User name"`
	FirstName string `json:"firstName" required:"false" doc:"First name"`
	LastName  string `json:"lastName" required:"false" doc:"Last name or family name"`
	Address   string `json:"address" required:"false" doc:"Address"`
	Image     string `json:"image" required:"false" doc:"Thumbnail"`
	Language  string `json:"language" required:"false" doc:"Language"`
}

type UserMfaResponse struct {
	Email         bool `json:"email" required:"false" doc:"Is 2FA enabled with email ?"`
	PhoneNumber   bool `json:"phoneNumber" required:"false" doc:"Is 2FA enabled with phone number ?"`
	Authenticator bool `json:"authenticator" required:"false" doc:"Is 2FA enabled with authenticator ?"`
}

type UserResponseList struct {
	types.PaginatedResponse
	Data []UserResponse `json:"data" required:"false" doc:"List of users" example:"[]"`
}
