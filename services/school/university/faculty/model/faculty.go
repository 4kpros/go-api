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

func (year *Faculty) ToResponse() *data.FacultyResponse {
	resp := &data.FacultyResponse{}
	resp.SchoolID = year.SchoolID
	resp.Name = year.Name
	resp.Description = year.Description
	return resp
}

func ToResponseList(yearList []Faculty) []data.FacultyResponse {
	resp := make([]data.FacultyResponse, len(yearList))
	for index, year := range yearList {
		resp[index] = *year.ToResponse()
	}
	return resp
}
