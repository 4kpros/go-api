package model

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/data"
)

type UserMfa struct {
	types.BaseGormModel
	Email         bool `gorm:"default:false"`
	PhoneNumber   bool `gorm:"default:false"`
	Authenticator bool `gorm:"default:false"`
}

func (userMfa *UserMfa) ToResponse() *data.UserMfaResponse {
	userMfaResp := &data.UserMfaResponse{}
	userMfaResp.Email = userMfa.Email
	userMfaResp.PhoneNumber = userMfa.PhoneNumber
	userMfaResp.Authenticator = userMfa.Authenticator
	return userMfaResp
}
