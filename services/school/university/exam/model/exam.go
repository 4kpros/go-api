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

func (item *Exam) ToResponse() *data.ExamResponse {
	resp := &data.ExamResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.SchoolID = item.SchoolID
	resp.TeachingUnitID = item.TeachingUnitID
	resp.Type = item.Type
	resp.Percentage = item.Percentage
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Exam) []data.ExamResponse {
	resp := make([]data.ExamResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
