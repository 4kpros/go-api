package model

import (
	"api/common/types"
	"api/services/school/university/tu/data"
)

type TeachingUnit struct {
	types.BaseGormModel
	SchoolID               int64                   `gorm:"not null"`
	DomainID               int64                   `gorm:"not null"`
	LevelID                int64                   `gorm:"not null"`
	Name                   string                  `gorm:"not null"`
	Description            string                  `gorm:"default:null"`
	Credit                 int                     `gorm:"default:1"`
	Semester               int                     `gorm:"default:1"`
	Program                string                  `gorm:"default null"`
	Requirements           string                  `gorm:"default null"`
	TeachingUnitProfessors []TeachingUnitProfessor `gorm:"default:null;foreignKey:TeachingUnitID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (teachingUnit *TeachingUnit) ToResponse() *data.TeachingUnitResponse {
	resp := &data.TeachingUnitResponse{}
	resp.ID = teachingUnit.ID
	resp.CreatedAt = teachingUnit.CreatedAt
	resp.UpdatedAt = teachingUnit.UpdatedAt
	resp.DeletedAt = teachingUnit.DeletedAt

	resp.SchoolID = teachingUnit.SchoolID
	resp.DomainID = teachingUnit.DomainID
	resp.LevelID = teachingUnit.LevelID
	resp.Name = teachingUnit.Name
	resp.Description = teachingUnit.Description
	resp.Credit = teachingUnit.Credit
	resp.Semester = teachingUnit.Semester
	resp.Program = teachingUnit.Program
	resp.Requirements = teachingUnit.Requirements
	resp.TeachingUnitProfessors = ToTeachingUnitProfessorResponseList(teachingUnit.TeachingUnitProfessors)
	return resp
}

func ToTeachingUnitResponseList(teachingUnitList []TeachingUnit) []data.TeachingUnitResponse {
	resp := make([]data.TeachingUnitResponse, len(teachingUnitList))
	for index, teachingUnit := range teachingUnitList {
		resp[index] = *teachingUnit.ToResponse()
	}
	return resp
}
