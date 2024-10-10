package data

type PermissionId struct {
	Id int64 `json:"id" path:"id" doc:"Permission id" example:"1"`
}

type CreatePermissionRequest struct {
	RoleId int64  `json:"roleId" doc:"Role id" example:"1"`
	Table  string `json:"table" required:"true" doc:"Table name" example:"history"`
	Create bool   `json:"create" required:"false" doc:"Create permission" example:""`
	Read   bool   `json:"read" required:"false" doc:"Read permission" example:""`
	Update bool   `json:"update" required:"false" doc:"Update permission" example:""`
	Delete bool   `json:"delete" required:"false" doc:"Delete permission" example:""`
}

type UpdatePermissionRequest struct {
	Create bool `json:"create" required:"false" doc:"Create permission" example:""`
	Read   bool `json:"read" required:"false" doc:"Read permission" example:""`
	Update bool `json:"update" required:"false" doc:"Update permission" example:""`
	Delete bool `json:"delete" required:"false" doc:"Delete permission" example:""`
}
