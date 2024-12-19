package data

type CommunicationID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Communication id" example:"1"`
}

type CommunicationRequest struct {
	Subject       string `json:"subject" required:"true" doc:"Subject" example:""`
	Message       string `json:"message" required:"true" doc:"Message" example:""`
	AudienceType  string `json:"audienceType" required:"true" doc:"Audience type" example:""`
	AudienceValue string `json:"audienceValue" required:"true" doc:"Audience value" example:""`
}
