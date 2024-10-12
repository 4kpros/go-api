package model

import (
	"github.com/4kpros/go-api/common/types"
)

type History struct {
	types.BaseGormModel
	Action string `json:"action" doc:"User action. E.g. read, write, update, delete" example:"read"`
	UserId int64  `json:"userId" doc:"User id" example:"1"`
	Table  string `json:"table" doc:"Affected table name" example:"user"`
	RowId  int64  `json:"rowId" doc:"Affected row inside table" example:"1"`
}
