package model

import (
	"api/common/types"
	"api/services/user/user/data"
)

type UserMfa struct {
	types.BaseGormModel
	Email         bool `gorm:"default:false"`
	PhoneNumber   bool `gorm:"default:false"`
	Authenticator bool `gorm:"default:false"`
}

func (userMfa *UserMfa) ToResponse() *data.UserMfaResponse {
	userMfaResp := &data.UserMfaResponse{
		Email:         userMfa.Email,
		PhoneNumber:   userMfa.PhoneNumber,
		Authenticator: userMfa.Authenticator,
	}
	return userMfaResp
}
