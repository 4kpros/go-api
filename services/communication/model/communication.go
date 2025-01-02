package model

import (
	"api/common/types"
	"api/services/communication/data"
)

type Communication struct {
	types.BaseGormModel
	Subject       string `gorm:"default:null"`
	Message       string `gorm:"not null"`
	AudienceType  string `gorm:"not null"`
	AudienceValue string `gorm:"not null"`
}

func (item *Communication) ToResponse() *data.CommunicationResponse {
	if item == nil {
		return nil
	}
	resp := &data.CommunicationResponse{}
	resp.Subject = item.Subject
	resp.Message = item.Message
	resp.AudienceType = item.AudienceType
	resp.AudienceValue = item.AudienceValue

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Communication) []data.CommunicationResponse {
	resp := make([]data.CommunicationResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
