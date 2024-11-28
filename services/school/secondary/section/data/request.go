package data

type SectionID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Section id" example:"1"`
}

type CreateSectionRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Section name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Section description" example:""`
}

type UpdateSectionRequest struct {
	Name        string `json:"name" required:"true" doc:"Section name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Section description" example:""`
}
