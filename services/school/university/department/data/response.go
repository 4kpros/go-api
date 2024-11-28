package data

import (
	"api/common/types"
)

type DepartmentResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"false" doc:"School id"`
	FacultyID   int64  `json:"facultyID" required:"false" doc:"Faculty id"`
	Name        string `json:"name" required:"false" doc:"Department name"`
	Description string `json:"description" required:"false" doc:"Department description"`
}

type DepartmentResponseList struct {
	types.PaginatedResponse
	Data []DepartmentResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
