package model

import (
	"time"

	"api/common/types"
	"api/services/admin/user/data"
)

type UserInfo struct {
	types.BaseGormModel
	UserName      string    `gorm:"default:null"`
	FirstName     string    `gorm:"default:null"`
	LastName      string    `gorm:"default:null"`
	Birthday      time.Time `gorm:"default:null"`
	BirthLocation string    `gorm:"default:null"`
	Address       string    `gorm:"default:null"`
	Language      string    `gorm:"default:null"`
	Image         string    `gorm:"default:null"`
}

func (userInfo *UserInfo) ToResponse() *data.UserInfoResponse {
	userInfoResp := &data.UserInfoResponse{}
	userInfoResp.UserName = userInfo.UserName
	userInfoResp.FirstName = userInfo.FirstName
	userInfoResp.LastName = userInfo.LastName
	userInfoResp.Birthday = userInfo.Birthday
	userInfoResp.BirthLocation = userInfo.BirthLocation
	userInfoResp.Address = userInfo.Address
	userInfoResp.Language = userInfo.Language
	userInfoResp.Image = userInfo.Image
	return userInfoResp
}
