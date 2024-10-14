package model

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/permission/data"
)

type Permission struct {
	types.BaseGormModel
	RoleId int64  `gorm:"default:null"`
	Table  string `gorm:"default:null"`
	Create bool   `gorm:"default:false"`
	Read   bool   `gorm:"default:false"`
	Update bool   `gorm:"default:false"`
	Delete bool   `gorm:"default:false"`
}

func (permission *Permission) ToResponse() *data.PermissionResponse {
	resp := &data.PermissionResponse{}
	resp.ID = permission.ID
	resp.CreatedAt = permission.CreatedAt
	resp.UpdatedAt = permission.UpdatedAt
	resp.DeletedAt = permission.DeletedAt
	resp.RoleId = permission.RoleId
	resp.Table = permission.Table
	resp.Create = permission.Create
	resp.Read = permission.Read
	resp.Update = permission.Update
	resp.Delete = permission.Delete
	return resp
}

func ToResponseList(permissionList []Permission) []data.PermissionResponse {
	resp := make([]data.PermissionResponse, len(permissionList))
	for index, permission := range permissionList {
		resp[index] = *permission.ToResponse()
	}
	return resp
}
