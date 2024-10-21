package model

import (
	"api/common/types"
	"api/services/admin/permission/data"
)

type PermissionFeature struct {
	types.BaseGormModel
	RoleId      int64  `gorm:"default:null"`
	FeatureName string `gorm:"default:null"`
	IsEnabled   bool   `gorm:"default:false"`
}

func (permission *PermissionFeature) ToResponse() *data.PermissionFeatureResponse {
	resp := &data.PermissionFeatureResponse{
		RoleId:      permission.RoleId,
		FeatureName: permission.FeatureName,
		IsEnabled:   permission.IsEnabled,
	}
	resp.ID = permission.ID
	resp.CreatedAt = permission.CreatedAt
	resp.UpdatedAt = permission.UpdatedAt
	resp.DeletedAt = permission.DeletedAt
	return resp
}
