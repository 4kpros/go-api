package model

import (
	"time"

	"api/common/types"
	"api/services/user/user/data"
)

type UserInfo struct {
	types.BaseGormModel
	UserID        int64      `gorm:"unique;not null"`
	Username      string     `gorm:"default:null"`
	FirstName     string     `gorm:"default:null"`
	LastName      string     `gorm:"default:null"`
	Birthday      *time.Time `gorm:"default:null"`
	BirthLocation string     `gorm:"default:null"`
	Address       string     `gorm:"default:null"`
	Language      string     `gorm:"default:null"`
	Image         string     `gorm:"default:null"`
}

func (item *UserInfo) ToResponse() *data.UserInfoResponse {
	resp := &data.UserInfoResponse{}
	if item == nil {
		return resp
	}
	resp = &data.UserInfoResponse{
		UserID:        item.UserID,
		Username:      item.Username,
		FirstName:     item.FirstName,
		LastName:      item.LastName,
		Birthday:      item.Birthday,
		BirthLocation: item.BirthLocation,
		Address:       item.Address,
		Language:      item.Language,
		Image:         item.Image,
	}
	return resp
}
