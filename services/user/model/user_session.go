package model

import (
	"github.com/4kpros/go-api/common/types"
)

type UserSession struct {
	types.BaseGormModel
	Token  string `json:"Token"`
	UserId string `json:"userId"`
}
