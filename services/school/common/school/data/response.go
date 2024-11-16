package data

import (
	"api/common/types"
	"time"
)

type SchoolResponse struct {
	types.BaseGormModelResponse
	Name       string              `json:"name" doc:"School name"`
	Type       string              `json:"type" doc:"School description"`
	SchoolInfo *SchoolInfoResponse `json:"schoolInfo" doc:"School description"`
}

type SchoolInfoResponse struct {
	FullName    string `json:"fullName" doc:"School name"`
	Description string `json:"description" doc:"Description"`
	Devise      string `json:"devise" doc:"Devise"`

	Founders  string    `json:"founders" doc:"Founder name"`
	FoundedAt time.Time `json:"foundedAt" doc:"Founded date time"`

	Address           string  `json:"address" doc:"Address"`
	LocationLongitude float64 `json:"locationLongitude" doc:"Location longitude"`
	LocationLatitude  float64 `json:"locationLatitude" doc:"Location latitude"`

	Image string `json:"image" doc:"School logo"`
}

type SchoolResponseList struct {
	types.PaginatedResponse
	Data []SchoolResponse `json:"data" required:"false" doc:"List of schools" example:"[]"`
}
