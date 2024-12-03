package model

import (
	"api/common/types"
	"api/services/school/secondary/class/data"
)

type Class struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (item *Class) ToResponse() *data.ClassResponse {
	resp := &data.ClassResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Class) []data.ClassResponse {
	resp := make([]data.ClassResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
