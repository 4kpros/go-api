package data

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/permission/model"
)

type PermissionsResponse struct {
	types.PaginatedResponse
	Data []model.Permission `json:"data" required:"false" doc:"Array of permissions" example:"[]"`
}
