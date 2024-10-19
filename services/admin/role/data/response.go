package data

import (
	"api/common/types"
)

type RoleResponse struct {
	types.BaseGormModelResponse
	Name        string `json:"name" doc:"Role name"`
	Description string `json:"description" doc:"Role description"`
}

type RoleResponseList struct {
	types.PaginatedResponse
	Data []RoleResponse `json:"data" required:"false" doc:"List of roles" example:"[]"`
}
