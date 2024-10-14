package model

import (
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/data"
	"gorm.io/gorm"
)

type User struct {
	types.BaseGormModel
	Email          string     `gorm:"default:null"`
	PhoneNumber    uint64     `gorm:"default:null"`
	Provider       string     `gorm:"default:null"`
	ProviderUserId string     `gorm:"default:null"`
	IsActivated    bool       `gorm:"default:null"`
	ActivatedAt    *time.Time `gorm:"default:null"`
	RoleId         int64      `gorm:"default:null"`
	Language       string     `gorm:"default:en"`
	Password       string     `gorm:"default:null"`

	UserInfo   UserInfo `gorm:"default:null;foreignKey:UserInfoId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	UserInfoId int64    `gorm:"default:null"`

	Mfa   MFA   `gorm:"default:null;foreignKey:MfaId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	MfaId int64 `gorm:"default:null"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password, err = utils.EncodeArgon2id(user.Password)
	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	user.Password, err = utils.EncodeArgon2id(user.Password)
	return
}

func (user *User) ToResponse() *data.UserResponse {
	resp := &data.UserResponse{}
	resp.ID = user.ID
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
	resp.DeletedAt = user.DeletedAt
	resp.Email = user.Email
	resp.PhoneNumber = user.PhoneNumber
	resp.Provider = user.Provider
	resp.ProviderUserId = user.ProviderUserId
	resp.IsActivated = user.IsActivated
	resp.ActivatedAt = user.ActivatedAt
	resp.RoleId = user.RoleId
	resp.Language = user.Language
	resp.UserInfo = *user.UserInfo.ToResponse()
	resp.Mfa = *user.Mfa.ToResponse()
	return resp
}

func ToResponseList(userList []User) []data.UserResponse {
	resp := make([]data.UserResponse, len(userList))
	for index, user := range userList {
		resp[index] = *user.ToResponse()
	}
	return resp
}
