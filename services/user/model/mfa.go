package model

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/data"
)

type MFA struct {
	types.BaseGormModel
	Email         bool `gorm:"default:false"`
	PhoneNumber   bool `gorm:"default:false"`
	Authenticator bool `gorm:"default:false"`
}

func (mfa *MFA) ToResponse() *data.MFAResponse {
	mfaResp := &data.MFAResponse{}
	mfaResp.Email = mfa.Email
	mfaResp.PhoneNumber = mfa.PhoneNumber
	mfaResp.Authenticator = mfa.Authenticator
	return mfaResp
}
