package model

import (
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"gorm.io/gorm"
)

type User struct {
	types.BaseGormModel
	Email          string     `json:"email" doc:"Email" minLength:"3" maxLength:"20" example:"example@domain.com"`
	PhoneNumber    uint64     `json:"phoneNumber" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Provider       string     `json:"provider" doc:"Provider" example:"google"`
	ProviderUserId string     `json:"providerUserId" doc:"User id from provider"  example:"121"`
	IsActivated    bool       `json:"isActivated" doc:"Is user account activated ?"  example:"false"`
	ActivatedAt    *time.Time `json:"activatedAt"`
	RoleId         int64      `json:"roleId" doc:"Role id" example:"1"`
	Language       string     `json:"language" doc:"Language with 2 letter" minLength:"2" maxLength:"2" example:"en"`
	Password       string     `json:"password"`
	UserInfo       UserInfo   `json:"userInfo,omitempty" gorm:"foreignKey:UserInfoId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	UserInfoId     int64      `json:"_" gorm:"default:null"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password, err = utils.EncodeArgon2id(user.Password)
	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	user.Password, err = utils.EncodeArgon2id(user.Password)
	return
}
