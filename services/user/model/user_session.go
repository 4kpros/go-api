package model

import (
	"github.com/4kpros/go-api/common/types"
)

type UserSession struct {
	types.BaseGormModel
	Token  string `json:"Token" doc:"Access token" example:""`
	UserId int64  `json:"userId" doc:"User id" example:""`
}
