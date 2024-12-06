package data

import (
	"api/common/types"
)

type PupilResponse struct {
	types.BaseGormModelResponse
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id"`
	UserID   int64 `json:"userID" required:"true" doc:"User id"`
	ClassID  int64 `json:"classID" required:"true" doc:"Class id"`
}

type PupilResponseList struct {
	types.PaginatedResponse
	Data []PupilResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
