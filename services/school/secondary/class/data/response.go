package data

import (
	"api/common/types"
)

type ClassResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"false" doc:"School id"`
	Name        string `json:"name" required:"false" doc:"Class name"`
	Description string `json:"description" required:"false" doc:"Class description"`
}

type ClassResponseList struct {
	types.PaginatedResponse
	Data []ClassResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
