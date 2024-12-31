package data

import (
	"api/common/types"
)

type DocumentResponse struct {
	types.BaseGormModelResponse
	SchoolID       int64  `json:"schoolID" required:"false" doc:"School id"`
	YearID         int64  `json:"yearID" required:"false" doc:"Year id"`
	SubjectID      int64  `json:"subjectID" required:"false" doc:"Subject id"`
	TeachingUnitID int64  `json:"teachingUnitID" required:"false" doc:"Teaching unit id"`
	Type           string `json:"type" required:"false" doc:"Type"`
	URL            string `json:"url" required:"false" doc:"URL"`
	Name           string `json:"name" required:"false" doc:"Name"`
	Description    string `json:"description" required:"false" doc:"Description"`
}

type DocumentResponseList struct {
	types.PaginatedResponse
	Data []DocumentResponse `json:"data" required:"false" doc:"List of documents" example:"[]"`
}
