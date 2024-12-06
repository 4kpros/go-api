package data

import (
	"api/common/types"
)

type StudentResponse struct {
	types.BaseGormModelResponse
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id"`
	UserID   int64 `json:"userID" required:"true" doc:"User id"`
	LevelID  int64 `json:"levelID" required:"true" doc:"Level id"`
}

type StudentResponseList struct {
	types.PaginatedResponse
	Data []StudentResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
