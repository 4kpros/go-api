package data

type ClassID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Class id" example:"1"`
}

type CreateClassRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Class name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Class description" example:""`
}

type UpdateClassRequest struct {
	Name        string `json:"name" required:"true" doc:"Class name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Class description" example:""`
}
