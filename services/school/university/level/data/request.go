package data

type LevelID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Level id" example:"1"`
}

type CreateLevelRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Level name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Level description" example:""`
}

type UpdateLevelRequest struct {
	Name        string `json:"name" required:"true" doc:"Level name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Level description" example:""`
}
