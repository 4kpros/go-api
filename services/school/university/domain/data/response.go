package data

import (
	"api/common/types"
)

type DomainResponse struct {
	types.BaseGormModelResponse
	SchoolID     int64  `json:"schoolID" required:"false" doc:"School id"`
	DepartmentID int64  `json:"DepartmentID" required:"false" doc:"Department id"`
	Name         string `json:"name" required:"false" doc:"Domain name"`
	Description  string `json:"description" required:"false" doc:"Domain description"`
}

type DomainResponseList struct {
	types.PaginatedResponse
	Data []DomainResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
