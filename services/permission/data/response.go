package data

import (
	"github.com/4kpros/go-api/common/types"
)

type PermissionResponse struct {
	types.BaseGormModelResponse
	RoleId int64  `json:"roleId" required:"false" doc:"Role id"`
	Table  string `json:"table" required:"false" doc:"Table name"`
	Create bool   `json:"create" required:"false" doc:"Create permission"`
	Read   bool   `json:"read" required:"false" doc:"Read permission"`
	Update bool   `json:"update" required:"false" doc:"Update permission"`
	Delete bool   `json:"delete" required:"false" doc:"Delete permission"`
}

type PermissionList struct {
	types.PaginatedResponse
	Data []PermissionResponse `json:"data" required:"false" doc:"List of permissions" example:"[]"`
}
