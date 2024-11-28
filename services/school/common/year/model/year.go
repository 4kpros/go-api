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

func (year *Year) ToResponse() *data.YearResponse {
	resp := &data.YearResponse{}
	resp.StartDate = year.StartDate
	resp.EndDate = year.EndDate
	return resp
}

func ToResponseList(yearList []Year) []data.YearResponse {
	resp := make([]data.YearResponse, len(yearList))
	for index, year := range yearList {
		resp[index] = *year.ToResponse()
	}
	return resp
}
