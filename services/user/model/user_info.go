package model

import (
	"github.com/4kpros/go-api/common/types"
)

type UserInfo struct {
	types.BaseGormModel
	UserName  string `json:"userName" doc:"User name" example:""`
	FirstName string `json:"firstName" doc:"First name" example:""`
	LastName  string `json:"lastName" doc:"Last name or family name" example:""`
	Address   string `json:"address" doc:"Address" example:""`
	Image     string `json:"image" doc:"Thumbnail" example:""`
}
