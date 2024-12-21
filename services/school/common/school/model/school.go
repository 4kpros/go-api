package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type School struct {
	types.BaseGormModel
	Name string `gorm:"unique;not null"`
	Type string `gorm:"not null"`

	SchoolConfigID int64         `gorm:"default:null"`
	Config         *SchoolConfig `gorm:"default:null;foreignKey:SchoolConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SchoolInfoID int64       `gorm:"default:null"`
	Info         *SchoolInfo `gorm:"default:null;foreignKey:SchoolInfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Directors []SchoolDirector `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *School) ToResponse() *data.SchoolResponse {
	resp := &data.SchoolResponse{
		Name:      item.Name,
		Type:      item.Type,
		Info:      item.Info.ToResponse(),
		Config:    item.Config.ToResponse(),
		Directors: ToSchoolDirectorResponseList(item.Directors),
	}

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToSchoolResponseList(itemList []School) []data.SchoolResponse {
	resp := make([]data.SchoolResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
