package model

import (
	"api/common/types"
	"api/services/school/secondary/section/data"
)

type Section struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (item *Section) ToResponse() *data.SectionResponse {
	resp := &data.SectionResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SchoolID = item.SchoolID
	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Section) []data.SectionResponse {
	resp := make([]data.SectionResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
