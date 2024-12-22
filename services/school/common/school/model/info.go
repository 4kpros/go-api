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

	Founder   string     `gorm:"default:null"`
	FoundedAt *time.Time `gorm:"default:null"`

	Address           string  `gorm:"default:null"`
	LocationLongitude float64 `gorm:"default:0"`
	LocationLatitude  float64 `gorm:"default:0"`

	Logo string `gorm:"default:null"`

	Image1 string `gorm:"default:null"`
	Image2 string `gorm:"default:null"`
	Image3 string `gorm:"default:null"`
	Image4 string `gorm:"default:null"`
}

func (item *SchoolInfo) ToResponse() *data.SchoolInfoResponse {
	resp := &data.SchoolInfoResponse{}
	if item == nil {
		return resp
	}
	resp.FullName = item.FullName
	resp.Description = item.Description
	resp.Slogan = item.Slogan

	resp.PhoneNumber1 = item.PhoneNumber1
	resp.PhoneNumber2 = item.PhoneNumber2
	resp.PhoneNumber3 = item.PhoneNumber3

	resp.Email1 = item.Email1
	resp.Email2 = item.Email2
	resp.Email3 = item.Email3

	resp.Founder = item.Founder
	resp.FoundedAt = item.FoundedAt

	resp.Address = item.Address
	resp.LocationLongitude = item.LocationLongitude
	resp.LocationLatitude = item.LocationLatitude

	resp.Logo = item.Logo

	resp.Image1 = item.Image1
	resp.Image2 = item.Image2
	resp.Image3 = item.Image3
	resp.Image4 = item.Image4
	return resp
}

func FromInfoRequest(item *data.SchoolInfoRequest) *SchoolInfo {
	resp := &SchoolInfo{
		FullName:    item.FullName,
		Description: item.Description,
		Slogan:      item.Slogan,

		PhoneNumber1: item.PhoneNumber1,
		PhoneNumber2: item.PhoneNumber2,
		PhoneNumber3: item.PhoneNumber3,

		Email1: item.Email1,
		Email2: item.Email2,
		Email3: item.Email3,

		Founder:   item.Founder,
		FoundedAt: item.FoundedAt,

		Address:           item.Address,
		LocationLongitude: item.LocationLongitude,
		LocationLatitude:  item.LocationLatitude,

		Logo: item.Logo,

		Image1: item.Image1,
		Image2: item.Image2,
		Image3: item.Image3,
		Image4: item.Image4,
	}
	return resp
}
