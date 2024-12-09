package data

import (
	"api/common/types"
)

type TestResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	SubjectID   int64  `json:"subjectID" required:"false" doc:"Subject id"`
	Type        string `json:"type" required:"false" doc:"Type"`
	Percentage  int    `json:"percentage" required:"false" doc:"Percentage"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type TestResponseList struct {
	types.PaginatedResponse
	Data []TestResponse `json:"data" required:"false" doc:"List of tests" example:"[]"`
}
