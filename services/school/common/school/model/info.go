package model

import (
	"time"

	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolInfo struct {
	types.BaseGormModel
	FullName    string `gorm:"default:null"`
	Description string `gorm:"default:null"`
	Devise      string `gorm:"default:null"`

	Founders  string    `gorm:"default:null"`
	FoundedAt time.Time `gorm:"default:null"`

	Address           string  `gorm:"default:null"`
	LocationLongitude float64 `gorm:"default:0"`
	LocationLatitude  float64 `gorm:"default:0"`

	Image string `gorm:"default:null"`
}

func (schoolInfo *SchoolInfo) ToResponse() *data.SchoolInfoResponse {
	schoolInfoResp := &data.SchoolInfoResponse{
		FullName:          schoolInfo.FullName,
		Description:       schoolInfo.Description,
		Devise:            schoolInfo.Devise,
		Founders:          schoolInfo.Founders,
		FoundedAt:         schoolInfo.FoundedAt,
		Address:           schoolInfo.Address,
		LocationLongitude: schoolInfo.LocationLongitude,
		LocationLatitude:  schoolInfo.LocationLatitude,
		Image:             schoolInfo.Image,
	}
	return schoolInfoResp
}
