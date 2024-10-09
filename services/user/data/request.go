package data

type UserId struct {
	Id int `json:"id" path:"id" doc:"User id" example:"29"`
}

type UserWithEmailRequest struct {
	Email string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	Role  int    `json:"role" required:"true" doc:"Role id" example:"1"`
}

type UserWithPhoneNumberRequest struct {
	PhoneNumber int `json:"phoneNumber" required:"true" doc:"Phone number" example:"690909090"`
	Role        int `json:"role" required:"true" doc:"Role id" example:"1"`
}

type UserRequest struct {
	Email       string `json:"email" doc:"Email" minLength:"2" maxLength:"20" example:"example@domain.com"`
	PhoneNumber int    `json:"phoneNumber" doc:"Phone number" minLength:"8" maxLength:"15" example:"690909090"`
	Role        int    `json:"role" doc:"Role id" example:"1"`
	Language    string `json:"language" doc:"Language with 2 letter" minLength:"2" maxLength:"2" example:"en"`
}

type UserInfoRequest struct {
	UserName  string `json:"userName" doc:"User name" minLength:"2" maxLength:"30" example:"meta_human"`
	FirstName string `json:"firstName" doc:"First name" minLength:"2" maxLength:"30" example:"John"`
	LastName  string `json:"lastName" doc:"Last name" minLength:"2" maxLength:"30" example:"Doe"`
	Address   string `json:"address" doc:"Address" minLength:"2" maxLength:"30" example:"Meta verse"`
}
