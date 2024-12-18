package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type School struct {
	types.BaseGormModel
	Name string `gorm:"unique;not null"`
	Type string `gorm:"not null"`

	SchoolInfo   *SchoolInfo `gorm:"default:null;foreignKey:SchoolInfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	SchoolInfoID int64       `gorm:"default:null"`

	SchoolConfig   *SchoolConfig `gorm:"default:null;foreignKey:SchoolConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	SchoolConfigID int64         `gorm:"default:null"`

	SchoolDirectors []SchoolDirector `gorm:"default:null;foreignKey:SchoolId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (school *School) ToResponse() *data.SchoolResponse {
	resp := &data.SchoolResponse{}
	resp.ID = school.ID
	resp.CreatedAt = school.CreatedAt
	resp.UpdatedAt = school.UpdatedAt
	resp.DeletedAt = school.DeletedAt
	resp.Name = school.Name
	resp.Type = school.Type
	resp.SchoolInfo = school.SchoolInfo.ToResponse()
	resp.SchoolConfig = school.SchoolConfig.ToResponse()
	resp.SchoolDirectors = ToSchoolDirectorResponseList(school.SchoolDirectors)
	return resp
}

func ToSchoolResponseList(schoolList []School) []data.SchoolResponse {
	resp := make([]data.SchoolResponse, len(schoolList))
	for index, school := range schoolList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
