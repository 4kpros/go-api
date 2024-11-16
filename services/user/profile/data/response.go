package data

import (
	"api/common/types"
	"api/services/user/user/model"
	"time"
)

type UserProfileResponse struct {
	types.BaseGormModelResponse
	Email       string `json:"email" required:"false" doc:"Email"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number"`
	RoleId      int64  `json:"roleId" required:"false" doc:"RoleId id"`

	LoginMethod    string     `json:"loginMethod" required:"false" doc:"How the user should login ? with email, phone number or external provider?"`
	Provider       string     `json:"provider" required:"false" doc:"Provider name"`
	ProviderUserId string     `json:"providerUserId" required:"false" doc:"User id from the provider"`
	IsActivated    bool       `json:"isActivated" required:"false" doc:"Is user account activated ?"`
	ActivatedAt    *time.Time `json:"activatedAt" required:"false" doc:"Activation date time"`

	UserInfo *UserProfileInfoResponse `json:"info" required:"false" doc:"Additional user info(e.g. address, first name, last name, ...)" `
	UserMfa  *UserProfileMfaResponse  `json:"mfa" required:"false" doc:"Multiple factor authenticator enabled by the user"`
}

type UserProfileInfoResponse struct {
	Username  string `json:"username" required:"false" doc:"User name"`
	FirstName string `json:"firstName" required:"false" doc:"First name"`
	LastName  string `json:"lastName" required:"false" doc:"Last name or family name"`

	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string     `json:"address" required:"false" doc:"Address"`
	Language      string     `json:"language" required:"false" doc:"Language"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}

type UserProfileMfaResponse struct {
	Email         bool `json:"email" required:"false" doc:"Is 2FA enabled with email ?"`
	PhoneNumber   bool `json:"phoneNumber" required:"false" doc:"Is 2FA enabled with phone number ?"`
	Authenticator bool `json:"authenticator" required:"false" doc:"Is 2FA enabled with authenticator ?"`
}

func FromUser(user *model.User) *UserProfileResponse {
	resp := &UserProfileResponse{}
	resp.ID = user.ID
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
	resp.DeletedAt = user.DeletedAt

	resp.Email = user.Email
	resp.PhoneNumber = user.PhoneNumber
	resp.RoleId = user.RoleId

	resp.LoginMethod = user.LoginMethod
	resp.Provider = user.Provider
	resp.ProviderUserId = user.ProviderUserId
	resp.IsActivated = user.IsActivated
	resp.ActivatedAt = user.ActivatedAt

	resp.UserInfo = FromUserInfo(user.UserInfo)
	resp.UserMfa = FromUserMfa(user.UserMfa)
	return resp
}

func FromUserInfo(userInfo *model.UserInfo) *UserProfileInfoResponse {
	resp := &UserProfileInfoResponse{}
	resp.Username = userInfo.Username
	resp.FirstName = userInfo.FirstName
	resp.LastName = userInfo.LastName
	resp.Birthday = userInfo.Birthday
	resp.BirthLocation = userInfo.BirthLocation
	resp.Address = userInfo.Address
	resp.Language = userInfo.Language
	resp.Image = userInfo.Image
	return resp
}

func FromUserMfa(userMfa *model.UserMfa) *UserProfileMfaResponse {
	resp := &UserProfileMfaResponse{}
	resp.Email = userMfa.Email
	resp.PhoneNumber = userMfa.PhoneNumber
	resp.Authenticator = userMfa.Authenticator
	return resp
}
