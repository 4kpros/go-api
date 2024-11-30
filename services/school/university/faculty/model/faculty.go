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

func (faculty *Faculty) ToResponse() *data.FacultyResponse {
	resp := &data.FacultyResponse{}
	resp.ID = faculty.ID
	resp.CreatedAt = faculty.CreatedAt
	resp.UpdatedAt = faculty.UpdatedAt
	resp.DeletedAt = faculty.DeletedAt

	resp.SchoolID = faculty.SchoolID
	resp.Name = faculty.Name
	resp.Description = faculty.Description
	return resp
}

func ToResponseList(facultyList []Faculty) []data.FacultyResponse {
	resp := make([]data.FacultyResponse, len(facultyList))
	for index, faculty := range facultyList {
		resp[index] = *faculty.ToResponse()
	}
	return resp
}
