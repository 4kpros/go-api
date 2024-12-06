package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolDirector struct {
	types.BaseGormModel
	SchoolID int64 `gorm:"not null"`
	UserID   int64 `gorm:"not null"`
}

func (item *SchoolDirector) ToResponse() *data.SchoolDirectorResponse {
	resp := &data.SchoolDirectorResponse{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}
	return resp
}

func ToSchoolDirectorResponseList(itemList []SchoolDirector) []data.SchoolDirectorResponse {
	resp := make([]data.SchoolDirectorResponse, len(itemList))
	for index, school := range itemList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
