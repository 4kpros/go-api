package data

type PermissionId struct {
	Id int64 `json:"id" path:"id" required:"true" minLength:"1" doc:"Permission id" example:"1"`
}

type CreatePermissionRequest struct {
	RoleId int64  `json:"roleId" required:"true" minLength:"1" doc:"Role id" example:"1"`
	Table  string `json:"table" required:"true" minLength:"2" doc:"Table name" example:"history"`
	Create bool   `json:"create" required:"true" doc:"Create permission" example:"false"`
	Read   bool   `json:"read" required:"true" doc:"Read permission" example:"false"`
	Update bool   `json:"update" required:"true" doc:"Update permission" example:"false"`
	Delete bool   `json:"delete" required:"true" doc:"Delete permission" example:"false"`
}

type UpdatePermissionRequest struct {
	Create bool `json:"create" required:"true" doc:"Create permission" example:"false"`
	Read   bool `json:"read" required:"true" doc:"Read permission" example:"false"`
	Update bool `json:"update" required:"true" doc:"Update permission" example:"false"`
	Delete bool `json:"delete" required:"true" doc:"Delete permission" example:"false"`
}
