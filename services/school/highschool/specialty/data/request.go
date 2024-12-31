package data

type SpecialtyID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Specialty id" example:"1"`
}

type SpecialtyRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	SectionID   int64  `json:"sectionID" required:"true" doc:"Section id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Specialty name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Specialty description" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
