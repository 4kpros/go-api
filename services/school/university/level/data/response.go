package data

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type LevelResponse struct {
	types.BaseGormModelResponse
	School      *data.SchoolResponse `json:"school" doc:"School"`
	Name        string               `json:"name" required:"false" doc:"Level name"`
	Description string               `json:"description" required:"false" doc:"Level description"`
}

type LevelResponseList struct {
	types.PaginatedResponse
	Data []LevelResponse `json:"data" required:"false" doc:"List of levels" example:"[]"`
}
