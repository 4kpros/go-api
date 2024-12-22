package model

import (
	"api/common/types"
	"api/services/school/common/director/data"
	schoolModel "api/services/school/common/school/model"
	userModel "api/services/user/user/model"
)

type Director struct {
	types.BaseGormModel

	UserID int64           `gorm:"not null"`
	User   *userModel.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SchoolID int64               `gorm:"not null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Director) ToResponse() *data.DirectorResponse {
	resp := &data.DirectorResponse{}
	if item == nil {
		return resp
	}
	resp.User = *item.User.ToResponse()
	resp.User.Role = nil
	resp.User.Info = nil
	resp.User.Mfa = nil
	resp.School = *item.School.ToResponse()
	resp.School.Config = nil
	resp.School.Info = nil

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Director) []data.DirectorResponse {
	resp := make([]data.DirectorResponse, len(itemList))
	for index, school := range itemList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
