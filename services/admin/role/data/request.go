package data

type RoleId struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Role id" example:"1"`
}

type RoleRequest struct {
	Name        string `json:"name" required:"true" minLength:"2" maxLength:"30" doc:"Role name" example:"client"`
	Description string `json:"description" required:"false" minLength:"2" maxLength:"200" doc:"Role description" example:"Client role used to allow users to access your services"`
}
