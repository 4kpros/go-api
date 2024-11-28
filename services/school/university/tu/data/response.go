package data

import (
	"api/common/types"
)

type TeachingUnitResponse struct {
	types.BaseGormModelResponse
	SchoolID               int64                           `json:"schoolID" required:"false" doc:"School id"`
	DomainID               int64                           `json:"domainID" required:"false" doc:"Domain id"`
	LevelID                int64                           `json:"levelID" required:"false" doc:"Level id"`
	Name                   string                          `json:"name" required:"false" doc:"Name"`
	Description            string                          `json:"description" required:"false" doc:"Description"`
	Credit                 int                             `json:"credit" required:"false" doc:"Credit"`
	Semester               int                             `json:"semester" required:"false" doc:"Semester"`
	Program                string                          `json:"Program" required:"false" doc:"Program"`
	Requirements           string                          `json:"Requirements" required:"false" doc:"Requirements"`
	TeachingUnitProfessors []TeachingUnitProfessorResponse `json:"teachingUnitProfessors" required:"false" doc:"Teaching unit list"`
}

type TeachingUnitProfessorResponse struct {
	types.BaseGormModelResponse
	TeachingUnitID int64 `json:"teachingUnitID" required:"false" doc:"Teaching unit id"`
	UserID         int64 `json:"userID" required:"false" doc:"User id"`
}

type TeachingUnitResponseList struct {
	types.PaginatedResponse
	Data []TeachingUnitResponse `json:"data" required:"false" doc:"List of teaching unit" example:"[]"`
}
