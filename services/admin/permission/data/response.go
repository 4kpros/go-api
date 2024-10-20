package data

import (
	"api/common/types"
)

type PermissionResponse struct {
	types.BaseGormModelResponse
	RoleId      int64                   `json:"roleId" required:"false" doc:"Role id"`
	FeatureName string                  `json:"featureName" required:"false" minLength:"2" doc:"Feature name"`
	IsEnabled   bool                    `json:"isEnabled" required:"false" doc:"Is this feature enabled ?"`
	Table       PermissionTableResponse `json:"table" required:"false" doc:"Table  permission"`
}

type PermissionTableResponse struct {
	TableName string `json:"tableName" required:"false" doc:"Table name"`
	Create    bool   `json:"create" required:"false" doc:"Create permission"`
	Read      bool   `json:"read" required:"false" doc:"Read permission"`
	Update    bool   `json:"update" required:"false" doc:"Update permission"`
	Delete    bool   `json:"delete" required:"false" doc:"Delete permission"`
}

type PermissionListResponse struct {
	types.PaginatedResponse
	Data []PermissionResponse `json:"data" required:"false" doc:"List of permissions" example:"[]"`
}
