package model

import (
	"api/common/types"
	"api/services/school/university/level/data"
)

type Level struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (item *Level) ToResponse() *data.LevelResponse {
	resp := &data.LevelResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Level) []data.LevelResponse {
	resp := make([]data.LevelResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
