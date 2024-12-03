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

	LoginMethod    string     `json:"loginMethod" required:"false" doc:"How the user should login ? with email, phone number or external provider?"`
	Provider       string     `json:"provider" required:"false" doc:"Provider name"`
	ProviderUserID string     `json:"providerUserID" required:"false" doc:"User id from the provider"`
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

type UserLoginResponse struct {
	LoginMethod string `json:"loginMethod" required:"false" doc:"Login method"`
	Provider    string `json:"provider" required:"false" doc:"Provider"`
	Role        string `json:"role" required:"false" doc:"Role"`
	Feature     string `json:"feature" required:"false" doc:"Feature"`
	Username    string `json:"username" required:"false" doc:"Username"`
	FirstName   string `json:"firstName" required:"false" doc:"First name"`
	LastName    string `json:"lastName" required:"false" doc:"Last name"`
	Image       string `json:"image" required:"false" doc:"Image"`
}

func FromUser(item *model.User) *UserProfileResponse {
	resp := &UserProfileResponse{}
	if item == nil {
		return resp
	}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber

	resp.LoginMethod = item.LoginMethod
	resp.Provider = item.Provider
	resp.ProviderUserID = item.ProviderUserID
	resp.IsActivated = item.IsActivated
	resp.ActivatedAt = item.ActivatedAt

	resp.UserInfo = FromUserInfo(item.UserInfo)
	resp.UserMfa = FromUserMfa(item.UserMfa)
	return resp
}

func FromUserInfo(item *model.UserInfo) *UserProfileInfoResponse {
	resp := &UserProfileInfoResponse{}
	if item == nil {
		return resp
	}
	resp.Username = item.Username
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.Birthday = item.Birthday
	resp.BirthLocation = item.BirthLocation
	resp.Address = item.Address
	resp.Language = item.Language
	resp.Image = item.Image
	return resp
}

func FromUserMfa(item *model.UserMfa) *UserProfileMfaResponse {
	resp := &UserProfileMfaResponse{}
	if item == nil {
		return resp
	}
	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber
	resp.Authenticator = item.Authenticator
	return resp
}
