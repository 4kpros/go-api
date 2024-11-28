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

func (schoolDirector *SchoolDirector) ToResponse() *data.SchoolDirectorResponse {
	schoolDirectorResp := &data.SchoolDirectorResponse{
		SchoolID: schoolDirector.SchoolID,
		UserID:   schoolDirector.UserID,
	}
	return schoolDirectorResp
}

func ToSchoolDirectorResponseList(schoolDirectorList []SchoolDirector) []data.SchoolDirectorResponse {
	resp := make([]data.SchoolDirectorResponse, len(schoolDirectorList))
	for index, school := range schoolDirectorList {
		resp[index] = *school.ToResponse()
	}
	return resp
}
