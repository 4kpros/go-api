package data

import (
	"api/common/types"
)

type LevelResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"false" doc:"School id"`
	Name        string `json:"name" required:"false" doc:"Level name"`
	Description string `json:"description" required:"false" doc:"Level description"`
}

type LevelResponseList struct {
	types.PaginatedResponse
	Data []LevelResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
