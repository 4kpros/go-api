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

func (item *TeachingUnit) ToResponse() *data.TeachingUnitResponse {
	resp := &data.TeachingUnitResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.SchoolID = item.SchoolID
	resp.DomainID = item.DomainID
	resp.LevelID = item.LevelID
	resp.Name = item.Name
	resp.Description = item.Description
	resp.Credit = item.Credit
	resp.Semester = item.Semester
	resp.Program = item.Program
	resp.Requirements = item.Requirements
	resp.TeachingUnitProfessors = ToTeachingUnitProfessorResponseList(item.TeachingUnitProfessors)
	return resp
}

func ToTeachingUnitResponseList(itemList []TeachingUnit) []data.TeachingUnitResponse {
	resp := make([]data.TeachingUnitResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
