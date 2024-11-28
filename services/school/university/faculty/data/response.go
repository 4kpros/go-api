package data

import (
	"api/common/types"
)

type FacultyResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"false" doc:"School ID"`
	Name        string `json:"name" required:"false" doc:"Faculty name"`
	Description string `json:"description" required:"false" doc:"Faculty description"`
}

type FacultyResponseList struct {
	types.PaginatedResponse
	Data []FacultyResponse `json:"data" required:"false" doc:"List of faculties" example:"[]"`
}
