package data

type RoleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Role id" example:"1"`
}

type RoleRequest struct {
	Name        string `json:"name" required:"true" minLength:"2" maxLength:"30" doc:"Role name" example:"client"`
	Feature     string `json:"feature" required:"true" minLength:"2" maxLength:"30" doc:"Feature name" example:"feature-admin"`
	Description string `json:"description" required:"false" minLength:"2" maxLength:"200" doc:"Role description" example:"Client role used to allow users to access your services"`
}
