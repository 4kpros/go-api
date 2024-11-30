package model

import (
	"api/common/types"
	"api/services/history/data"
)

type History struct {
	types.BaseGormModel
	Action string `gorm:"default:null"`
	UserID int64  `gorm:"default:null"`
	Table  string `gorm:"default:null"`
	RowID  int64  `gorm:"default:null"`
}

func (history *History) ToResponse() *data.HistoryResponse {
	resp := &data.HistoryResponse{
		Action: history.Action,
		UserID: history.UserID,
		Table:  history.Table,
		RowID:  history.RowID,
	}

	resp.ID = history.ID
	resp.CreatedAt = history.CreatedAt
	resp.UpdatedAt = history.UpdatedAt
	resp.DeletedAt = history.DeletedAt
	return resp
}

func ToResponseList(historyList []History) []data.HistoryResponse {
	resp := make([]data.HistoryResponse, len(historyList))
	for index, history := range historyList {
		resp[index] = *history.ToResponse()
	}
	return resp
}
