package data

type DirectorID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Director id" example:"1"`
}

type DirectorRequest struct {
	UserID   int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
}
