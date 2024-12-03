package model

import (
	"api/common/types"
	"api/services/school/university/faculty/data"
)

type Faculty struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (item *Faculty) ToResponse() *data.FacultyResponse {
	resp := &data.FacultyResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.SchoolID = item.SchoolID
	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Faculty) []data.FacultyResponse {
	resp := make([]data.FacultyResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
