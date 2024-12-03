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
	resp = &data.PermissionResponse{
		TableName: item.TableName,
		Create:    item.Create,
		Read:      item.Read,
		Update:    item.Update,
		Delete:    item.Delete,
	}
	return resp
}
