package model

import (
	"api/common/types"
	"api/services/school/common/year/data"
	"time"
)

type Year struct {
	types.BaseGormModel
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time `gorm:"not null"`
}

func (academicYear *Year) ToResponse() *data.YearResponse {
	resp := &data.YearResponse{}
	resp.StartDate = academicYear.StartDate
	resp.EndDate = academicYear.EndDate
	return resp
}

func ToResponseList(academicYearList []Year) []data.YearResponse {
	resp := make([]data.YearResponse, len(academicYearList))
	for index, academicYear := range academicYearList {
		resp[index] = *academicYear.ToResponse()
	}
	return resp
}
