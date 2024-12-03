package model

import (
	"api/common/types"
	"api/services/school/secondary/pupil/data"
)

type Pupil struct {
	types.BaseGormModel
	SchoolID int64 `gorm:"not null"`
	UserID   int64 `gorm:"not null"`
	ClassID  int64 `gorm:"default:null"`
}

func (item *Pupil) ToResponse() *data.PupilResponse {
	resp := &data.PupilResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.SchoolID = item.SchoolID
	resp.UserID = item.UserID
	resp.ClassID = item.ClassID
	return resp
}

func ToResponseList(itemList []Pupil) []data.PupilResponse {
	resp := make([]data.PupilResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
