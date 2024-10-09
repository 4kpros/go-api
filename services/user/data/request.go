package data

type UserId struct {
	Id int `json:"id" path:"id" required:"true" doc:"User id" example:"29"`
}

type UserWithEmailRequest struct {
	Email string `json:"email" required:"true" doc:"Email" minLength:"3" maxLength:"30" example:"example@domain.com"`
	Role  int    `json:"role" required:"true" doc:"Role id" example:"1"`
}

type UserWithPhoneNumberRequest struct {
	PhoneNumber int `json:"phoneNumber" required:"true" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Role        int `json:"role" required:"true" doc:"Role id" example:"1"`
}

type UserRequest struct {
	Email       string `json:"email" required:"false" doc:"Email" minLength:"2" maxLength:"30" example:"example@domain.com"`
	PhoneNumber int    `json:"phoneNumber" required:"false" doc:"Phone number" minLength:"11" maxLength:"25" example:"237690909090"`
	Role        int    `json:"role" required:"true" doc:"Role id" example:"1"`
	Language    string `json:"language" required:"false" doc:"Language with 2 letter" minLength:"2" maxLength:"2" example:"en"`
}

type UserInfoRequest struct {
	UserName  string `json:"userName" required:"false" doc:"User name" minLength:"2" maxLength:"30" example:"meta_human"`
	FirstName string `json:"firstName" required:"false" doc:"First name" minLength:"2" maxLength:"30" example:"John"`
	LastName  string `json:"lastName" required:"false" doc:"Last name" minLength:"2" maxLength:"30" example:"Doe"`
	Address   string `json:"address" required:"false" doc:"Address" minLength:"2" maxLength:"30" example:"Meta verse"`
}
