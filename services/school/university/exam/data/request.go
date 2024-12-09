package data

type ExamID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Exam id" example:"1"`
}

type CreateExamRequest struct {
	SchoolID       int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	TeachingUnitID int64  `json:"teachingUnitID" required:"true" doc:"Teaching unit id" example:"1"`
	Type           string `json:"type" required:"true" enum:"CC,TP,EE" doc:"Type" example:"CC"`
	Percentage     int    `json:"percentage" required:"true" doc:"Percentage" example:""`
	Description    string `json:"description" required:"false" doc:"Description" example:""`
}

type UpdateExamRequest struct {
	Type        string `json:"type" required:"true" enum:"CC,TP,EE" doc:"Type" example:"CC"`
	Percentage  int    `json:"percentage" required:"true" doc:"Percentage" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`
}
