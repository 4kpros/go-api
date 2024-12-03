package data

type PermissionPathRequest struct {
	RoleID int64 `json:"roleID" path:"roleID" required:"true" doc:"Role id" example:"1"`
}

type UpdatePermissionFeatureRequest struct {
	Feature string `json:"feature" required:"true" doc:"Feature permission"`
}

type UpdatePermissionTableRequest struct {
	TableName string `json:"tableName" required:"true" minLength:"2" doc:"Table name" example:"users"`
	Create    bool   `json:"create" required:"true" doc:"Create permission" example:"false"`
	Read      bool   `json:"read" required:"true" doc:"Read permission" example:"false"`
	Update    bool   `json:"update" required:"true" doc:"Update permission" example:"false"`
	Delete    bool   `json:"delete" required:"true" doc:"Delete permission" example:"false"`
}
