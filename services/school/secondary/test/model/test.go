package model

import (
	"api/common/types"
	"api/services/school/secondary/test/data"
)

type Test struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	SubjectID   int64  `gorm:"not null"`
	Type        string `gorm:"not null"`
	Percentage  int    `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (test *Test) ToResponse() *data.TestResponse {
	resp := &data.TestResponse{}
	resp.SchoolID = test.SchoolID
	resp.SubjectID = test.SubjectID
	resp.Type = test.Type
	resp.Percentage = test.Percentage
	resp.Description = test.Description
	return resp
}

func ToResponseList(testList []Test) []data.TestResponse {
	resp := make([]data.TestResponse, len(testList))
	for index, test := range testList {
		resp[index] = *test.ToResponse()
	}
	return resp
}
