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
	Slogan      string `gorm:"default:null"`

	PhoneNumber1 int64 `gorm:"default:null"`
	PhoneNumber2 int64 `gorm:"default:null"`
	PhoneNumber3 int64 `gorm:"default:null"`

	Email1 string `gorm:"default:null"`
	Email2 string `gorm:"default:null"`
	Email3 string `gorm:"default:null"`

	Founder   string    `gorm:"default:null"`
	FoundedAt time.Time `gorm:"default:null"`

	Address           string  `gorm:"default:null"`
	LocationLongitude float64 `gorm:"default:0"`
	LocationLatitude  float64 `gorm:"default:0"`

	Logo string `gorm:"default:null"`

	Image1 string `gorm:"default:null"`
	Image2 string `gorm:"default:null"`
	Image3 string `gorm:"default:null"`
	Image4 string `gorm:"default:null"`
}

func (schoolInfo *SchoolInfo) ToResponse() *data.SchoolInfoResponse {
	schoolInfoResp := &data.SchoolInfoResponse{
		FullName:    schoolInfo.FullName,
		Description: schoolInfo.Description,
		Slogan:      schoolInfo.Slogan,

		PhoneNumber1: schoolInfo.PhoneNumber1,
		PhoneNumber2: schoolInfo.PhoneNumber2,
		PhoneNumber3: schoolInfo.PhoneNumber3,

		Email1: schoolInfo.Email1,
		Email2: schoolInfo.Email2,
		Email3: schoolInfo.Email3,

		Founder:   schoolInfo.Founder,
		FoundedAt: schoolInfo.FoundedAt,

		Address:           schoolInfo.Address,
		LocationLongitude: schoolInfo.LocationLongitude,
		LocationLatitude:  schoolInfo.LocationLatitude,

		Logo: schoolInfo.Logo,

		Image1: schoolInfo.Image1,
		Image2: schoolInfo.Image2,
		Image3: schoolInfo.Image3,
		Image4: schoolInfo.Image4,
	}
	return schoolInfoResp
}
