package data

import (
	"api/common/types"
)

type ContactResponse struct {
	types.BaseGormModelResponse
	Subject string `json:"subject" required:"true" doc:"Subject"`
	Email   string `json:"email" required:"true" doc:"Email"`
	Message string `json:"message" required:"true" doc:"Message"`
}

type ContactResponseList struct {
	types.PaginatedResponse
	Data []ContactResponse `json:"data" required:"false" doc:"List of contacts" example:"[]"`
}
