package model

import (
	"api/common/types"
	"api/services/user/permission/data"
)

type Permission struct {
	types.BaseGormModel
	RoleID    int64  `gorm:"default:null"`
	TableName string `gorm:"default:null"`
	Create    bool   `gorm:"default:null"`
	Read      bool   `gorm:"default:null"`
	Update    bool   `gorm:"default:null"`
	Delete    bool   `gorm:"default:null"`
}

func (item *Permission) ToResponse() *data.PermissionResponse {
	resp := &data.PermissionResponse{}
	if item == nil {
		return resp
	}
	resp.TableName = item.TableName
	resp.Create = item.Create
	resp.Read = item.Read
	resp.Update = item.Update
	resp.Delete = item.Delete
	return resp
}
