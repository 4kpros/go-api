package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	specialtyData "api/services/school/highschool/specialty/data"
)

type ClassResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolResponse       `json:"school" doc:"School"`
	Specialty   *specialtyData.SpecialtyResponse `json:"specialty" doc:"Specialty"`
	Name        string                           `json:"name" required:"false" doc:"Class name"`
	Description string                           `json:"description" required:"false" doc:"Class description"`
}

type ClassResponseList struct {
	types.PaginatedResponse
	Data []ClassResponse `json:"data" required:"false" doc:"List of classes" example:"[]"`
}
