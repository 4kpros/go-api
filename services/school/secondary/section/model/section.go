package model

import (
	"api/common/types"
	"api/services/school/secondary/section/data"
)

type Section struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (section *Section) ToResponse() *data.SectionResponse {
	resp := &data.SectionResponse{}
	resp.SchoolID = section.SchoolID
	resp.Name = section.Name
	resp.Description = section.Description
	return resp
}

func ToResponseList(sectionList []Section) []data.SectionResponse {
	resp := make([]data.SectionResponse, len(sectionList))
	for index, section := range sectionList {
		resp[index] = *section.ToResponse()
	}
	return resp
}
