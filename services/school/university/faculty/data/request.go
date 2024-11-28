package data

type FacultyID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Faculty id" example:"1"`
}

type CreateFacultyRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School ID" example:"1"`
	Name        string `json:"name" required:"true" doc:"Faculty name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Faculty description" example:""`
}

type UpdateFacultyRequest struct {
	Name        string `json:"name" required:"true" doc:"Faculty name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Faculty description" example:""`
}
