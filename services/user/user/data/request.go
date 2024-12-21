package data

import "time"

type UserID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"User id" example:"1"`
}
type UserRoleRequest struct {
	RoleID int64 `json:"roleID" required:"true" doc:"Role id" example:"1"`
}

type CreateUserWithEmailRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	RoleID      int64  `json:"roleID" required:"true" doc:"Role id" example:"1"`
	IsActivated bool   `json:"isActivated" required:"true" doc:"Is activated" example:"false"`
}

type CreateUserWithPhoneNumberRequest struct {
	PhoneNumber uint64 `json:"phoneNumber" required:"true" doc:"Phone number" example:"237690909090"`
	RoleID      int64  `json:"roleID" required:"true" doc:"Role id" example:"1"`
	IsActivated bool   `json:"isActivated" required:"true" doc:"Is activated" example:"false"`
}

type UpdateUserRequest struct {
	Email       string `json:"email" required:"true" minLength:"3" max:"100" doc:"Email" example:"example@domain.com"`
	PhoneNumber uint64 `json:"phoneNumber" required:"false" doc:"Phone number" example:"237690909090"`
	RoleID      int64  `json:"roleID" required:"true" doc:"Role id" example:"1"`
	IsActivated bool   `json:"isActivated" required:"true" doc:"Is activated" example:"false"`
}

type UpdateUserInfoRequest struct {
	Gender    string `json:"Gender" required:"false" enum:"m,f" doc:"Gender" example:"M"`
	Username  string `json:"username" required:"false" max:"30" doc:"User name" example:"meta_human"`
	FirstName string `json:"firstName" required:"false" max:"30" doc:"First name" example:"John"`
	LastName  string `json:"lastName" required:"false" max:"30" doc:"Last name" example:"Doe"`

	Birthday      *time.Time `json:"birthday" required:"false" doc:"Birthday date time"`
	BirthLocation string     `json:"birthLocation" required:"false" doc:"Birth location"`
	Address       string     `json:"address" required:"false" max:"30" doc:"Address" example:"No City"`
	Language      string     `json:"language" required:"false" min:"2" max:"2" doc:"Language code with 2 letter" example:"en"`
	Image         string     `json:"image" required:"false" doc:"Thumbnail"`
}
