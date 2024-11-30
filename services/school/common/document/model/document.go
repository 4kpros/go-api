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

func (document *Document) ToResponse() *data.DocumentResponse {
	resp := &data.DocumentResponse{
		SchoolID:       document.SchoolID,
		YearID:         document.YearID,
		SubjectID:      document.SubjectID,
		TeachingUnitID: document.TeachingUnitID,
		Type:           document.Type,
		URL:            document.URL,
		Description:    document.Description,
	}

	resp.ID = document.ID
	resp.CreatedAt = document.CreatedAt
	resp.UpdatedAt = document.UpdatedAt
	resp.DeletedAt = document.DeletedAt
	return resp
}

func ToResponseList(documentList []Document) []data.DocumentResponse {
	resp := make([]data.DocumentResponse, len(documentList))
	for index, document := range documentList {
		resp[index] = *document.ToResponse()
	}
	return resp
}
