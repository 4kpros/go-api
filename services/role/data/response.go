package data

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/role/model"
)

type DeleteResponse struct {
	AffectedRows int64 `json:"affectedRows" doc:"Number of row affected with this update" example:"1"`
}

type GetAllResponse struct {
	types.PaginatedResponse
	Data []model.Role `json:"data" doc:"Number of row affected with this update" example:"[]"`
}
