package data

import (
	"api/common/types"
)

type PermissionResponse struct {
	types.BaseGormModelResponse
	RoleId           int64    `json:"roleId" required:"false" doc:"Role id"`
	FeatureName      string   `json:"featureName" required:"false" doc:"Feature name"`
	TablePermissions []string `json:"data" required:"false" doc:"List of tables with theirs permissions" example:"[]"`
}

type PermissionList struct {
	types.PaginatedResponse
	Data []PermissionResponse `json:"data" required:"false" doc:"List of permissions" example:"[]"`
}
