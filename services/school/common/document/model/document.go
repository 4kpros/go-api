package model

import (
	"api/common/types"
	"api/services/school/common/document/data"
)

type Document struct {
	types.BaseGormModel
	SchoolID       int64  `gorm:"not null"`
	YearID         int64  `gorm:"not null"`
	SubjectID      int64  `gorm:"default:null"`
	TeachingUnitID int64  `gorm:"default:null"`
	Type           string `gorm:"not null"`
	URL            string `gorm:"default:null"`
	Name           string `gorm:"not null"`
	Description    string `gorm:"default:null"`
}

func (item *Document) ToResponse() *data.DocumentResponse {
	resp := &data.DocumentResponse{
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		SubjectID:      item.SubjectID,
		TeachingUnitID: item.TeachingUnitID,
		Type:           item.Type,
		URL:            item.URL,
		Description:    item.Description,
	}

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt
	return resp
}

func ToResponseList(itemList []Document) []data.DocumentResponse {
	resp := make([]data.DocumentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
