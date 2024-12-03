package model

import (
	"api/common/types"
	"api/services/user/role/data"
)

type Role struct {
	types.BaseGormModel
	Name        string `gorm:"unique;not null"`
	Feature     string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *Role) ToResponse() *data.RoleResponse {
	resp := &data.RoleResponse{}
	if item == nil {
		return resp
	}
	resp = &data.RoleResponse{
		Name:        item.Name,
		Feature:     item.Feature,
		Description: item.Description,
	}

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt
	return resp
}

func ToResponseList(itemList []Role) []data.RoleResponse {
	resp := make([]data.RoleResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
