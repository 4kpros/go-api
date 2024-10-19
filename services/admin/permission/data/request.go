package data

type PermissionId struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Permission id" example:"1"`
}

type UpdatePermissionRequest struct {
	RoleId           int64    `json:"roleId" required:"true" doc:"Role id" example:"1"`
	FeatureName      string   `json:"featureName" required:"true" minLength:"2" doc:"Feature name" example:"feature-admin"`
	TablePermissions []string `json:"data" required:"true" doc:"List of tables with theirs permissions" example:"[\"users crud\", \"roles crud\"]"`
}
