package model

import (
	"api/common/types"
	"api/services/user/data"
)

type UserInfo struct {
	types.BaseGormModel
	UserName  string `gorm:"default:null"`
	FirstName string `gorm:"default:null"`
	LastName  string `gorm:"default:null"`
	Address   string `gorm:"default:null"`
	Image     string `gorm:"default:null"`
	Language  string `gorm:"default:en"`
}

func (userInfo *UserInfo) ToResponse() *data.UserInfoResponse {
	userInfoResp := &data.UserInfoResponse{}
	userInfoResp.UserName = userInfo.UserName
	userInfoResp.FirstName = userInfo.FirstName
	userInfoResp.LastName = userInfo.LastName
	userInfoResp.Address = userInfo.Address
	userInfoResp.Image = userInfo.Image
	userInfoResp.Language = userInfo.Language
	return userInfoResp
}
