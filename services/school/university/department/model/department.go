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

func (department *Department) ToResponse() *data.DepartmentResponse {
	resp := &data.DepartmentResponse{}
	resp.FacultyID = department.FacultyID
	resp.Name = department.Name
	resp.Description = department.Description
	return resp
}

func ToResponseList(departmentList []Department) []data.DepartmentResponse {
	resp := make([]data.DepartmentResponse, len(departmentList))
	for index, department := range departmentList {
		resp[index] = *department.ToResponse()
	}
	return resp
}
