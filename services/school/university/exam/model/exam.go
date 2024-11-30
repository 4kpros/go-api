package model

import (
	"api/common/types"
	"api/services/school/university/exam/data"
)

type Exam struct {
	types.BaseGormModel
	SchoolID       int64  `gorm:"not null"`
	TeachingUnitID int64  `gorm:"not null"`
	Type           string `gorm:"not null"`
	Percentage     int    `gorm:"not null"`
	Description    string `gorm:"not null"`
}

func (exam *Exam) ToResponse() *data.ExamResponse {
	resp := &data.ExamResponse{}
	resp.ID = exam.ID
	resp.CreatedAt = exam.CreatedAt
	resp.UpdatedAt = exam.UpdatedAt
	resp.DeletedAt = exam.DeletedAt

	resp.SchoolID = exam.SchoolID
	resp.TeachingUnitID = exam.TeachingUnitID
	resp.Type = exam.Type
	resp.Percentage = exam.Percentage
	resp.Description = exam.Description
	return resp
}

func ToResponseList(examList []Exam) []data.ExamResponse {
	resp := make([]data.ExamResponse, len(examList))
	for index, exam := range examList {
		resp[index] = *exam.ToResponse()
	}
	return resp
}
