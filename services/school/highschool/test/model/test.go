package model

import (
	"api/common/types"
	"api/services/school/highschool/test/data"
)

type Test struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	SubjectID   int64  `gorm:"not null"`
	Type        string `gorm:"not null"`
	Percentage  int    `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (item *Test) ToResponse() *data.TestResponse {
	resp := &data.TestResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SchoolID = item.SchoolID
	resp.SubjectID = item.SubjectID
	resp.Type = item.Type
	resp.Percentage = item.Percentage
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Test) []data.TestResponse {
	resp := make([]data.TestResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
