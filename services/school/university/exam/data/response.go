package data

import (
	"api/common/types"
)

type ExamResponse struct {
	types.BaseGormModelResponse
	SchoolID       int64  `json:"schoolID" required:"true" doc:"School id"`
	TeachingUnitID int64  `json:"teachingUnitID" required:"false" doc:"Teaching unit id"`
	Type           string `json:"type" required:"false" doc:"Type"`
	Percentage     int    `json:"percentage" required:"false" doc:"Percentage"`
	Description    string `json:"description" required:"false" doc:"Description"`
}

type ExamResponseList struct {
	types.PaginatedResponse
	Data []ExamResponse `json:"data" required:"false" doc:"List of exams" example:"[]"`
}
