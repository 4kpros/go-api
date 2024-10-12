package data

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/role/model"
)

type RolesResponse struct {
	types.PaginatedResponse
	Data []model.Role `json:"data" required:"false" doc:"List of roles" example:"[]"`
}
