package data

type PupilID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Pupil id" example:"1"`
}

type CreatePupilRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	ClassID  int64 `json:"classID" required:"false" doc:"Class id" example:"1"`
}

type UpdatePupilRequest struct {
	UserID  int64 `json:"userID" required:"true" doc:"User id" example:"1"`
	ClassID int64 `json:"classID" required:"true" doc:"Class id" example:"1"`
}
