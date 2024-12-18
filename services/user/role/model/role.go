package model

import (
	"api/common/types"
	"api/services/user/role/data"
)

type Role struct {
	types.BaseGormModel
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"default:null"`
}

func (role *Role) ToResponse() *data.RoleResponse {
	resp := &data.RoleResponse{}
	resp.ID = role.ID
	resp.CreatedAt = role.CreatedAt
	resp.UpdatedAt = role.UpdatedAt
	resp.DeletedAt = role.DeletedAt
	resp.Name = role.Name
	resp.Description = role.Description
	return resp
}

func ToResponseList(roleList []Role) []data.RoleResponse {
	resp := make([]data.RoleResponse, len(roleList))
	for index, role := range roleList {
		resp[index] = *role.ToResponse()
	}
	return resp
}
