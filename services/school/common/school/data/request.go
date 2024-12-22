package data

import (
	"time"
)

type SchoolID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"School id" example:"1"`
}

type SchoolRequest struct {
	Name   string               `json:"name" required:"true" minLength:"2" maxLength:"50" doc:"School name" example:"Harvard University"`
	Type   string               `json:"type" required:"true" minLength:"2" maxLength:"20" enum:"secondary,university" doc:"School type" example:"secondary"`
	Info   *SchoolInfoRequest   `json:"info" required:"true" doc:"Information"`
	Config *SchoolConfigRequest `json:"config" required:"true" doc:"Configuration"`
}

type SchoolInfoRequest struct {
	FullName    string `json:"fullName" required:"true" doc:"School name" example:""`
	Description string `json:"description" required:"false" doc:"Description" example:""`
	Slogan      string `json:"slogan" required:"false" doc:"Slogan" example:""`

	PhoneNumber1 int64 `json:"phoneNumber1" required:"false" doc:"Phone number 1" example:""`
	PhoneNumber2 int64 `json:"phoneNumber2" required:"false" doc:"Phone number 2" example:""`
	PhoneNumber3 int64 `json:"phoneNumber3" required:"false" doc:"Phone number 3" example:""`

	Email1 string `json:"email1" required:"false" doc:"Email 1" example:""`
	Email2 string `json:"email2" required:"false" doc:"Email 2" example:""`
	Email3 string `json:"email3" required:"false" doc:"Email 3" example:""`

	Founder   string     `json:"founder" required:"false" doc:"Founder name" example:""`
	FoundedAt *time.Time `json:"foundedAt" required:"false" doc:"Founded date time" example:""`

	Address           string  `json:"address" required:"false" doc:"Address" example:""`
	LocationLongitude float64 `json:"locationLongitude" required:"false" doc:"Location longitude" example:""`
	LocationLatitude  float64 `json:"locationLatitude" required:"false" doc:"Location latitude" example:""`

	Logo string `json:"logo" required:"false" doc:"School logo" example:""`

	Image1 string `json:"image1" required:"false" doc:"Image 1" example:""`
	Image2 string `json:"image2" required:"false" doc:"Image 2" example:""`
	Image3 string `json:"image3" required:"false" doc:"Image 3" example:""`
	Image4 string `json:"image4" required:"false" doc:"Image 4" example:""`
}

type SchoolConfigRequest struct {
	EmailDomain string `json:"emailDomain" required:"false" doc:"Email domain" example:""`
}
