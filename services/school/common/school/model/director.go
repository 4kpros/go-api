package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolDirector struct {
	types.BaseGormModel
	SchoolId int64 `gorm:"not null"`
	UserId   int64 `gorm:"not null"`
}

func (schoolDirector *SchoolDirector) ToResponse() *data.SchoolDirectorResponse {
	schoolDirectorResp := &data.SchoolDirectorResponse{
		SchoolId: schoolDirector.SchoolId,
		UserId:   schoolDirector.UserId,
	}
	return schoolDirectorResp
}
