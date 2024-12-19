package data

import (
	"api/common/types"
)

type SectionResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"false" doc:"School id"`
	Name        string `json:"name" required:"false" doc:"Section name"`
	Description string `json:"description" required:"false" doc:"Section description"`
}

type SectionResponseList struct {
	types.PaginatedResponse
	Data []SectionResponse `json:"data" required:"false" doc:"List of sections" example:"[]"`
}
