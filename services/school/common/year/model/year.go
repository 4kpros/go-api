package model

import (
	"api/common/types"
	"api/services/school/common/year/data"
	"time"
)

type Year struct {
	types.BaseGormModel
	Name      string     `gorm:"unique;not null"`
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time `gorm:"not null"`
}

func (item *Year) ToResponse() *data.YearResponse {
	resp := &data.YearResponse{
		Name:      item.Name,
		StartDate: item.StartDate,
		EndDate:   item.EndDate,
	}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Year) []data.YearResponse {
	resp := make([]data.YearResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
