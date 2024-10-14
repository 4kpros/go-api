package data

type UserId struct {
	Id int64 `json:"id" path:"id" required:"true" doc:"User id" example:"1"`
}

type CreateUserWithEmailRequest struct {
	Email  string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	RoleId int64  `json:"roleId" required:"true" doc:"RoleId id" example:"1"`
}

type CreateUserWithPhoneNumberRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	RoleId      int64  `json:"roleId" required:"true" doc:"RoleId id" example:"1"`
}

type UpdateUserInfoRequest struct {
	UserName  string `json:"userName" required:"false" minLength:"2" maxLength:"30" doc:"User name" example:"meta_human"`
	FirstName string `json:"firstName" required:"false" minLength:"2" maxLength:"30" doc:"First name" example:"John"`
	LastName  string `json:"lastName" required:"false" minLength:"2" maxLength:"30" doc:"Last name" example:"Doe"`
	Address   string `json:"address" required:"false" minLength:"2" maxLength:"30" doc:"Address" example:"No City"`
}

type UpdateUserRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" maxLength:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	RoleId      int64  `json:"roleId" required:"true" doc:"RoleId id" example:"1"`
	Language    string `json:"language" required:"false" minLength:"2" maxLength:"2" doc:"Language code with 2 letter" example:"en"`
}
