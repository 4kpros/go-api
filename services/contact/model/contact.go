package model

import (
	"api/common/types"
	"api/services/contact/data"
)

type Contact struct {
	types.BaseGormModel
	Subject string `gorm:"default:null"`
	Email   string `gorm:"not null"`
	Message string `gorm:"not null"`
}

func (item *Contact) ToResponse() *data.ContactResponse {
	resp := &data.ContactResponse{
		Subject: item.Subject,
		Email:   item.Email,
		Message: item.Message,
	}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Contact) []data.ContactResponse {
	resp := make([]data.ContactResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
