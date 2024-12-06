package model

import (
	"api/common/types"
	"api/services/school/university/department/data"
)

type Department struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	FacultyID   int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (item *Department) ToResponse() *data.DepartmentResponse {
	resp := &data.DepartmentResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.FacultyID = item.FacultyID
	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Department) []data.DepartmentResponse {
	resp := make([]data.DepartmentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
