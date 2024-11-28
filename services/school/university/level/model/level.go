package model

import (
	"api/common/types"
	"api/services/school/university/level/data"
)

type Level struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	DomainID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (level *Level) ToResponse() *data.LevelResponse {
	resp := &data.LevelResponse{}
	resp.DomainID = level.DomainID
	resp.Name = level.Name
	resp.Description = level.Description
	return resp
}

func ToResponseList(levelList []Level) []data.LevelResponse {
	resp := make([]data.LevelResponse, len(levelList))
	for index, level := range levelList {
		resp[index] = *level.ToResponse()
	}
	return resp
}
