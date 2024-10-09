package model

import (
	"github.com/4kpros/go-api/common/types"
)

type UserInfo struct {
	types.BaseGormModel
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	Image     string `json:"image"`
}
