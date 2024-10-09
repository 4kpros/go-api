package data

type RoleId struct {
	Id int `json:"id" path:"id" doc:"Role id" example:"29"`
}

type RoleRequest struct {
	Name        string `json:"name" required:"true" doc:"Role name" minLength:"2" maxLength:"20" example:"Client"`
	Description string `json:"description" doc:"Role description" minLength:"3" maxLength:"80" example:"Client role used to allow users to access your services"`
}
