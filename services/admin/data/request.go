package data

type CreateAdminRequest struct {
	Email    string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	Password string `json:"password" required:"true" minLength:"8" maxLength:"30" doc:"Base64 encoded password" example:""`
	Token    string `json:"token" required:"true" doc:"Token" minLength:"3" example:""`
}
