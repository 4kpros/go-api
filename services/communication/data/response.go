package data

import (
	"api/common/types"
)

type CommunicationResponse struct {
	types.BaseGormModelResponse
	Subject       string `json:"subject" required:"true" doc:"Subject"`
	Message       string `json:"message" required:"true" doc:"Message"`
	AudienceType  string `json:"audienceType" required:"true" doc:"Audience type"`
	AudienceValue string `json:"audienceValue" required:"true" doc:"Audience value"`
}

type CommunicationResponseList struct {
	types.PaginatedResponse
	Data []CommunicationResponse `json:"data" required:"false" doc:"List of communications" example:"[]"`
}
