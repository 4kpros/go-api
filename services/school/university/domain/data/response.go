package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	departmentData "api/services/school/university/department/data"
)

type DomainResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolResponse         `json:"school" doc:"School"`
	Department  *departmentData.DepartmentResponse `json:"department" doc:"Department"`
	Name        string                             `json:"name" required:"false" doc:"Department name"`
	Description string                             `json:"description" required:"false" doc:"Department description"`
}

type DomainResponseList struct {
	types.PaginatedResponse
	Data []DomainResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
