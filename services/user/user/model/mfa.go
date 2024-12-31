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

func (item *UserMfa) ToResponse() *data.UserMfaResponse {
	resp := &data.UserMfaResponse{}
	if item == nil {
		return resp
	}
	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber
	resp.Authenticator = item.Authenticator
	return resp
}
