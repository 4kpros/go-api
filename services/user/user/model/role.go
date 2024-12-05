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
		UserID: item.UserID,
		RoleID: item.RoleID,
	}

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt
	return resp
}
