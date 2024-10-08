package model

import (
	"github.com/4kpros/go-api/common/types"
)

type Role struct {
	types.BaseGormModel
	Name        string `json:"name" doc:"Role name" example:"Client"`
	Description string `json:"description" doc:"Role description" example:"Client role used to allow users to access your services"`
}
