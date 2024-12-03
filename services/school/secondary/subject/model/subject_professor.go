package model

import (
	"api/common/types"
	"api/services/school/secondary/subject/data"
)

type SubjectProfessor struct {
	types.BaseGormModel
	SubjectID int64 `gorm:"not null"`
	UserID    int64 `gorm:"not null"`
}

func (item *SubjectProfessor) ToResponse() *data.SubjectProfessorResponse {
	resp := &data.SubjectProfessorResponse{}
	resp.SubjectID = item.SubjectID
	resp.UserID = item.UserID
	return resp
}

func ToSubjectProfessorResponseList(itemList []SubjectProfessor) []data.SubjectProfessorResponse {
	resp := make([]data.SubjectProfessorResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
