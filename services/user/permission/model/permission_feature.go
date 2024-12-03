package model

import (
	"api/common/types"
	"api/services/user/permission/data"
)

type PermissionFeature struct {
	types.BaseGormModel
	RoleID  int64  `gorm:"default:null"`
	Feature string `gorm:"default:null"`
}

func (permission *PermissionFeature) ToResponse() *data.PermissionFeatureResponse {
	resp := &data.PermissionFeatureResponse{
		RoleID:  permission.RoleID,
		Feature: permission.Feature,
	}

	resp.ID = permission.ID
	resp.CreatedAt = permission.CreatedAt
	resp.UpdatedAt = permission.UpdatedAt
	resp.DeletedAt = permission.DeletedAt
	return resp
}
