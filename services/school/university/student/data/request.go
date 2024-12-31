package data

type StudentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Student id" example:"1"`
}

type CreateStudentRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	LevelID  int64 `json:"levelID" required:"false" doc:"Level id" example:"1"`
}

type UpdateStudentRequest struct {
	UserID  int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	LevelID int64 `json:"levelID" required:"true" doc:"Level id" example:"1"`
}
