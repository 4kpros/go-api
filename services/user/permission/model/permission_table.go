package model

import (
	"api/common/types"
	"api/services/user/permission/data"
)

type PermissionTable struct {
	types.BaseGormModel
	RoleID    int64  `gorm:"default:null"`
	TableName string `gorm:"default:null"`
	Create    bool   `gorm:"default:null"`
	Read      bool   `gorm:"default:null"`
	Update    bool   `gorm:"default:null"`
	Delete    bool   `gorm:"default:null"`
}

func (permission *PermissionTable) ToResponse() *data.PermissionTableResponse {
	resp := &data.PermissionTableResponse{
		TableName: permission.TableName,
		Create:    permission.Create,
		Read:      permission.Read,
		Update:    permission.Update,
		Delete:    permission.Delete,
	}
	return resp
}
