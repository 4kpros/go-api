package data

type SchoolID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"School id" example:"1"`
}

type SchoolRequest struct {
	Name string `json:"name" required:"true" minLength:"2" maxLength:"50" doc:"School name" example:"Harvard University"`
	Type string `json:"type" required:"true" minLength:"2" maxLength:"20" enum:"secondary,university" doc:"School type" example:"secondary"`
}

type AddDirectorRequestPath struct {
	SchoolID int64 `json:"schoolID" path:"id" required:"true" doc:"School id" example:"1"`
}
type AddDirectorRequestBody struct {
	UserID int64 `json:"userID" required:"true" doc:"User id" example:"1"`
}

type DeleteDirectorRequest struct {
	SchoolID int64 `json:"schoolID" path:"schoolID" required:"true" doc:"School id" example:"1"`
	UserID   int64 `json:"userID" path:"userID" required:"true" doc:"User id" example:"1"`
}
