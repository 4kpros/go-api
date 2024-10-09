package data

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/services/user/model"
)

type UserResponse struct {
	model.User
	Password string `json:"password"`
}

type UsersResponse struct {
	types.PaginatedResponse
	Data []model.User `json:"data" doc:"Array of users" example:"[]"`
}
