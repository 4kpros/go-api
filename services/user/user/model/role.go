package model

import (
	"api/common/types"
	"api/services/user/user/data"
)

type UserRole struct {
	types.BaseGormModel
	UserID int64 `gorm:"unique;not null"`
	RoleID int64 `gorm:"not null"`
}

func (item *UserRole) ToResponse() *data.UserRoleResponse {
	resp := &data.UserRoleResponse{}
	if item == nil {
		return resp
	}
	resp = &data.UserRoleResponse{
		ID: item.RoleID,
	}
	return resp
}
