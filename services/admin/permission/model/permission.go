package model

import (
	"api/common/types"
	"api/services/admin/permission/data"
)

type Permission struct {
	types.BaseGormModel
	RoleId           int64    `gorm:"default:null"`
	FeatureName      string   `gorm:"default:null"`
	TablePermissions []string `gorm:"default:null"` // E.g. ["roles.crud", "user.----"]
}

func (permission *Permission) ToResponse() *data.PermissionResponse {
	resp := &data.PermissionResponse{
		RoleId:           permission.RoleId,
		FeatureName:      permission.FeatureName,
		TablePermissions: permission.TablePermissions,
	}
	resp.ID = permission.ID
	resp.CreatedAt = permission.CreatedAt
	resp.UpdatedAt = permission.UpdatedAt
	resp.DeletedAt = permission.DeletedAt
	return resp
}

func ToResponseList(permissionList []Permission) []data.PermissionResponse {
	resp := make([]data.PermissionResponse, len(permissionList))
	for index, permission := range permissionList {
		resp[index] = *permission.ToResponse()
	}
	return resp
}
