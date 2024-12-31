package data

type TestID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Test id" example:"1"`
}

type CreateTestRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	SubjectID   int64  `json:"subjectID" required:"true" doc:"Subject id" example:"1"`
	Type        string `json:"type" required:"true" enum:"Written,Practical" doc:"Type" example:"CC"`
	Percentage  int    `json:"percentage" required:"true" doc:"Percentage" example:"100"`
	Description string `json:"description" required:"false" doc:"Description" example:""`
}

type UpdateTestRequest struct {
	Type        string `json:"type" required:"true" enum:"Written,Practical" doc:"Type" example:"CC"`
	Percentage  int    `json:"percentage" required:"true" doc:"Percentage" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`
}
