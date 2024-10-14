package model

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/data"
)

type UserInfo struct {
	types.BaseGormModel
	UserName  string `gorm:"default:null"`
	FirstName string `gorm:"default:null"`
	LastName  string `gorm:"default:null"`
	Address   string `gorm:"default:null"`
	Image     string `gorm:"default:null"`
}

func (userInfo *UserInfo) ToResponse() *data.UserInfoResponse {
	userInfoResp := &data.UserInfoResponse{}
	userInfoResp.UserName = userInfo.UserName
	userInfoResp.FirstName = userInfo.FirstName
	userInfoResp.LastName = userInfo.LastName
	userInfoResp.Address = userInfo.Address
	userInfoResp.Image = userInfo.Image
	return userInfoResp
}
