package model

import (
	"api/common/types"
	"api/services/school/secondary/subject/data"
)

type Subject struct {
	types.BaseGormModel
	SchoolID          int64              `gorm:"not null"`
	ClassID           int64              `gorm:"not null"`
	Name              string             `gorm:"not null"`
	Description       string             `gorm:"default:null"`
	Coefficient       int                `gorm:"default:1"`
	Program           string             `gorm:"default null"`
	Requirements      string             `gorm:"default null"`
	SubjectProfessors []SubjectProfessor `gorm:"default:null;foreignKey:SubjectID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (subject *Subject) ToResponse() *data.SubjectResponse {
	resp := &data.SubjectResponse{}
	resp.ID = subject.ID
	resp.CreatedAt = subject.CreatedAt
	resp.UpdatedAt = subject.UpdatedAt
	resp.DeletedAt = subject.DeletedAt

	resp.SchoolID = subject.SchoolID
	resp.ClassID = subject.ClassID
	resp.Name = subject.Name
	resp.Description = subject.Description
	resp.Coefficient = subject.Coefficient
	resp.Program = subject.Program
	resp.Requirements = subject.Requirements
	resp.SubjectProfessors = ToSubjectProfessorResponseList(subject.SubjectProfessors)
	return resp
}

func ToSubjectResponseList(subjectList []Subject) []data.SubjectResponse {
	resp := make([]data.SubjectResponse, len(subjectList))
	for index, subject := range subjectList {
		resp[index] = *subject.ToResponse()
	}
	return resp
}
