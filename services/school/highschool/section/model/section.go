package model

import (
	"api/common/types"
	"api/services/school/common/school/model"
	"api/services/school/highschool/section/data"
)

type HighschoolSection struct {
	types.BaseGormModel
	SchoolID int64         `gorm:"default:null"`
	School   *model.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolSection) ToResponse() *data.SectionResponse {
	if item == nil {
		return nil
	}
	resp := &data.SectionResponse{}
	resp.School = item.School.ToResponse()
	resp.Name = item.Name
	resp.Description = item.Description

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []HighschoolSection) []data.SectionResponse {
	resp := make([]data.SectionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
