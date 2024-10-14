package data

import (
	"github.com/4kpros/go-api/common/types"
)

type RoleResponse struct {
	types.BaseGormModelResponse
	Name        string `json:"name" doc:"Role name" example:"Client"`
	Description string `json:"description" doc:"Role description" example:"Client role used to allow roles to access your services"`
}

type RoleResponseList struct {
	types.PaginatedResponse
	Data []RoleResponse `json:"data" required:"false" doc:"List of roles" example:"[]"`
}
