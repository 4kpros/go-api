package data

import "time"

type YearID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Academic year id" example:"1"`
}

type YearRequest struct {
	StartDate *time.Time `json:"startDate" required:"true" doc:"Academic year start date" example:""`
	EndDate   *time.Time `json:"endDate" required:"true" doc:"Academic year end date" example:""`
}
