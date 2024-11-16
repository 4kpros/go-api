package data

import (
	"api/common/types"
)

type PermissionFeatureResponse struct {
	types.BaseGormModelResponse
	RoleId      int64  `json:"roleId" required:"false" doc:"Role id"`
	FeatureName string `json:"featureName" required:"false" minLength:"2" doc:"Feature name"`
	IsEnabled   bool   `json:"isEnabled" required:"false" doc:"Is this feature enabled ?"`
}

type PermissionTableResponse struct {
	TableName string `json:"tableName" required:"false" doc:"Table name"`
	Create    bool   `json:"create" required:"false" doc:"Create permission"`
	Read      bool   `json:"read" required:"false" doc:"Read permission"`
	Update    bool   `json:"update" required:"false" doc:"Update permission"`
	Delete    bool   `json:"delete" required:"false" doc:"Delete permission"`
}

type PermissionFeatureTableResponse struct {
	*PermissionFeatureResponse
	*PermissionTableResponse
}

type PermissionListResponse struct {
	types.PaginatedResponse
	Data []PermissionFeatureTableResponse `json:"data" required:"false" doc:"List of all permissions" example:"[]"`
}
