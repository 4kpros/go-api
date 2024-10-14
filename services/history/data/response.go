package data

import (
	"github.com/4kpros/go-api/common/types"
)

type HistoryResponse struct {
	types.BaseGormModelResponse
	Action string `json:"roleId" required:"false" doc:"Action name"`
	UserId int64  `json:"table" required:"false" doc:"User id"`
	Table  string `json:"create" required:"false" doc:"Table name"`
	RowId  int64  `json:"read" required:"false" doc:"Row id inside table"`
}

type HistoryList struct {
	types.PaginatedResponse
	Data []HistoryResponse `json:"data" required:"false" doc:"List of history" example:"[]"`
}
