package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type School struct {
	types.BaseGormModel
	Name         string      `gorm:"unique;not null"`
	Type         string      `gorm:"not null"`
	SchoolInfo   *SchoolInfo `gorm:"default:null;foreignKey:SchoolInfoId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	SchoolInfoId int64       `gorm:"default:null"`
}

type SchoolDirector struct {
	types.BaseGormModel
	SchoolId int64 `gorm:"not null"`
	UserId   int64 `gorm:"not null"`
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
	return resp
}

func ToResponseList(schoolList []School) []data.SchoolResponse {
	resp := make([]data.SchoolResponse, len(schoolList))
	for index, school := range schoolList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
