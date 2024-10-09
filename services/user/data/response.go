package data

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/model"
)

type UsersResponse struct {
	types.PaginatedResponse
	Data []model.User `json:"data" required:"false" doc:"Array of users" example:"[]"`
}
