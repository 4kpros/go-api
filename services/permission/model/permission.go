package model

import (
	"github.com/4kpros/go-api/common/types"
)

type Permission struct {
	types.BaseGormModel
	RoleId int64  `json:"roleId" doc:"Role id" example:"1"`
	Table  string `json:"table" doc:"Table name" example:"history"`
	Create bool   `json:"create" doc:"Create permission" example:""`
	Read   bool   `json:"read" doc:"Create permission" example:""`
	Update bool   `json:"update" doc:"Create permission" example:""`
	Delete bool   `json:"delete" doc:"Create permission" example:""`
}
