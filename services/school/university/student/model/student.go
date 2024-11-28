package model

import (
	"api/common/types"
	"api/services/school/university/student/data"
)

type Student struct {
	types.BaseGormModel
	SchoolID int64 `gorm:"not null"`
	UserID   int64 `gorm:"not null"`
	LevelID  int64 `gorm:"default:null"`
}

func (student *Student) ToResponse() *data.StudentResponse {
	resp := &data.StudentResponse{}
	resp.SchoolID = student.SchoolID
	resp.UserID = student.UserID
	resp.LevelID = student.LevelID
	return resp
}

func ToResponseList(studentList []Student) []data.StudentResponse {
	resp := make([]data.StudentResponse, len(studentList))
	for index, student := range studentList {
		resp[index] = *student.ToResponse()
	}
	return resp
}
