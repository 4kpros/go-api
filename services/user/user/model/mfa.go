package model

import (
	"api/common/types"
	"api/services/user/user/data"
)

type UserMfa struct {
	types.BaseGormModel
	UserID        int64 `gorm:"unique;not null"`
	Email         bool  `gorm:"default:false"`
	PhoneNumber   bool  `gorm:"default:false"`
	Authenticator bool  `gorm:"default:false"`
}

func (item *UserMfa) ToResponse() *data.UserMfaResponse {
	resp := &data.UserMfaResponse{}
	if item == nil {
		return resp
	}
	resp = &data.UserMfaResponse{
		UserID:        item.UserID,
		Email:         item.Email,
		PhoneNumber:   item.PhoneNumber,
		Authenticator: item.Authenticator,
	}
	return resp
}
