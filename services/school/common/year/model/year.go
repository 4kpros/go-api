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
	resp := &data.YearResponse{
		StartDate: year.StartDate,
		EndDate:   year.EndDate,
	}
	resp.ID = year.ID
	resp.CreatedAt = year.CreatedAt
	resp.UpdatedAt = year.UpdatedAt
	resp.DeletedAt = year.DeletedAt
	return resp
}

func ToResponseList(yearList []Year) []data.YearResponse {
	resp := make([]data.YearResponse, len(yearList))
	for index, year := range yearList {
		resp[index] = *year.ToResponse()
	}
	return resp
}
