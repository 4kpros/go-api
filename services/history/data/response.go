package data

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/history/model"
)

type HistoriesResponse struct {
	types.PaginatedResponse
	Data []model.History `json:"data" required:"false" doc:"Array of histories" example:"[]"`
}
