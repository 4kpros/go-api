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
	resp := &data.CommunicationResponse{
		Subject:       item.Subject,
		Message:       item.Message,
		AudienceType:  item.AudienceType,
		AudienceValue: item.AudienceValue,
	}
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
