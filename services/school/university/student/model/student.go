package model

import (
	"api/common/types"
	"api/services/school/university/student/data"
)

type Student struct {
	types.BaseGormModel
	SchoolID int64 `gorm:"not null"`
	UserID   int64 `gorm:"not null"`
	LevelID  int64 `gorm:"default:null"`
}

func (item *Student) ToResponse() *data.StudentResponse {
	resp := &data.StudentResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.SchoolID = item.SchoolID
	resp.UserID = item.UserID
	resp.LevelID = item.LevelID
	return resp
}

func ToResponseList(itemList []Student) []data.StudentResponse {
	resp := make([]data.StudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
