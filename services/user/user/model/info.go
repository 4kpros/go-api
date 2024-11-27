package model

import (
	"time"

	"api/common/types"
	"api/services/user/user/data"
)

type UserInfo struct {
	types.BaseGormModel
	Username      string     `gorm:"default:null"`
	FirstName     string     `gorm:"default:null"`
	LastName      string     `gorm:"default:null"`
	Birthday      *time.Time `gorm:"default:null"`
	BirthLocation string     `gorm:"default:null"`
	Address       string     `gorm:"default:null"`
	Language      string     `gorm:"default:null"`
	Image         string     `gorm:"default:null"`
}

func (userInfo *UserInfo) ToResponse() *data.UserInfoResponse {
	userInfoResp := &data.UserInfoResponse{
		Username:      userInfo.Username,
		FirstName:     userInfo.FirstName,
		LastName:      userInfo.LastName,
		Birthday:      userInfo.Birthday,
		BirthLocation: userInfo.BirthLocation,
		Address:       userInfo.Address,
		Language:      userInfo.Language,
		Image:         userInfo.Image,
	}
	return userInfoResp
}
