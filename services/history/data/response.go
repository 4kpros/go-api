package data

import (
	"api/common/types"
)

type HistoryResponse struct {
	types.BaseGormModelResponse
	Action string `json:"roleID" required:"false" doc:"Action name"`
	UserID int64  `json:"table" required:"false" doc:"User id"`
	Table  string `json:"create" required:"false" doc:"Table name"`
	RowID  int64  `json:"read" required:"false" doc:"Row id inside table"`
}

type HistoryList struct {
	types.PaginatedResponse
	Data []HistoryResponse `json:"data" required:"false" doc:"List of history" example:"[]"`
}
