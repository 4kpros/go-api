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

func (item *History) ToResponse() *data.HistoryResponse {
	if item == nil {
		return nil
	}
	resp := &data.HistoryResponse{}
	resp.Action = item.Action
	resp.UserID = item.UserID
	resp.Table = item.Table
	resp.RowID = item.RowID

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(historyList []History) []data.HistoryResponse {
	resp := make([]data.HistoryResponse, len(historyList))
	for index, history := range historyList {
		resp[index] = *history.ToResponse()
	}
	return resp
}
