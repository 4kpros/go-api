package data

import (
	"api/common/types"
	"time"
)

type YearResponse struct {
	types.BaseGormModelResponse
	StartDate *time.Time `json:"startDate" doc:"Academic year start date"`
	EndDate   *time.Time `json:"endDate" doc:"Academic year end date"`
}

type YearResponseList struct {
	types.PaginatedResponse
	Data []YearResponse `json:"data" required:"false" doc:"List of academic years" example:"[]"`
}
