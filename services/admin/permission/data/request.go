package data

type GetRoleFeaturePermissionRequest struct {
	RoleId      int64  `json:"roleId" path:"roleId" required:"true" doc:"Role id" example:"1"`
	FeatureName string `json:"featureName" path:"featureName" required:"true" doc:"Feature name" example:"feature-admin"`
}

type GetRolePermissionListRequest struct {
	RoleId int64 `json:"roleId" path:"roleId" required:"true" doc:"Role id" example:"1"`
}

type UpdateRoleFeaturePermissionPathRequest struct {
	RoleId      int64  `json:"roleId" path:"roleId" required:"true" doc:"Role id" example:"1"`
	FeatureName string `json:"featureName" path:"featureName" required:"true" doc:"Feature name" example:"feature-admin"`
}

type UpdateRoleFeaturePermissionBodyRequest struct {
	IsEnabled bool                         `json:"isEnabled" required:"true" doc:"Is this feature enabled ?" example:"false"`
	Table     UpdatePermissionTableRequest `json:"table" required:"true" doc:"Table  permission"`
}

type UpdatePermissionTableRequest struct {
	TableName string `json:"tableName" required:"true" minLength:"2" doc:"Table name" example:"users"`
	Create    bool   `json:"create" required:"true" doc:"Create permission" example:"false"`
	Read      bool   `json:"read" required:"true" doc:"Read permission" example:"false"`
	Update    bool   `json:"update" required:"true" doc:"Update permission" example:"false"`
	Delete    bool   `json:"delete" required:"true" doc:"Delete permission" example:"false"`
}
