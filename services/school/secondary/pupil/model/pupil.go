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

func (pupil *Pupil) ToResponse() *data.PupilResponse {
	resp := &data.PupilResponse{}
	resp.SchoolID = pupil.SchoolID
	resp.UserID = pupil.UserID
	resp.ClassID = pupil.ClassID
	return resp
}

func ToResponseList(pupilList []Pupil) []data.PupilResponse {
	resp := make([]data.PupilResponse, len(pupilList))
	for index, pupil := range pupilList {
		resp[index] = *pupil.ToResponse()
	}
	return resp
}
