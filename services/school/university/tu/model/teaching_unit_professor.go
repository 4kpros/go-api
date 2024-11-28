package model

import (
	"api/common/types"
	"api/services/school/university/tu/data"
)

type TeachingUnitProfessor struct {
	types.BaseGormModel
	TeachingUnitID int64 `gorm:"not null"`
	UserID         int64 `gorm:"not null"`
}

func (teachingUnitProfessor *TeachingUnitProfessor) ToResponse() *data.TeachingUnitProfessorResponse {
	resp := &data.TeachingUnitProfessorResponse{}
	resp.TeachingUnitID = teachingUnitProfessor.TeachingUnitID
	resp.UserID = teachingUnitProfessor.UserID
	return resp
}

func ToTeachingUnitProfessorResponseList(teachingUnitProfessorList []TeachingUnitProfessor) []data.TeachingUnitProfessorResponse {
	resp := make([]data.TeachingUnitProfessorResponse, len(teachingUnitProfessorList))
	for index, teachingUnitProfessor := range teachingUnitProfessorList {
		resp[index] = *teachingUnitProfessor.ToResponse()
	}
	return resp
}
