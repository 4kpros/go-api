package data

type SchoolId struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"School id" example:"1"`
}

type SchoolRequest struct {
	Name string `json:"name" required:"true" minLength:"2" maxLength:"30" doc:"School name" example:"client"`
	Type string `json:"type" required:"true" minLength:"2" maxLength:"20" enum:"secondary,university" doc:"School type" example:"secondary"`
}

type DirectorRequest struct {
	SchoolId int64 `json:"schoolId" required:"true" doc:"School id" example:"1"`
	UserId   int64 `json:"userId" required:"true" doc:"User id" example:"1"`
}
