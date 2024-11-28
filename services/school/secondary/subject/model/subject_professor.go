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

func (subjectProfessor *SubjectProfessor) ToResponse() *data.SubjectProfessorResponse {
	resp := &data.SubjectProfessorResponse{}
	resp.SubjectID = subjectProfessor.SubjectID
	resp.UserID = subjectProfessor.UserID
	return resp
}

func ToSubjectProfessorResponseList(subjectProfessorList []SubjectProfessor) []data.SubjectProfessorResponse {
	resp := make([]data.SubjectProfessorResponse, len(subjectProfessorList))
	for index, subjectProfessor := range subjectProfessorList {
		resp[index] = *subjectProfessor.ToResponse()
	}
	return resp
}
