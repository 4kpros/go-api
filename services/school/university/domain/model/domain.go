package model

import (
	"api/common/types"
	"api/services/school/university/domain/data"
)

type Domain struct {
	types.BaseGormModel
	SchoolID     int64  `gorm:"not null"`
	DepartmentID int64  `gorm:"not null"`
	Name         string `gorm:"not null"`
	Description  string `gorm:"not null"`
}

func (item *Domain) ToResponse() *data.DomainResponse {
	resp := &data.DomainResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	resp.DeletedAt = item.DeletedAt

	resp.DepartmentID = item.DepartmentID
	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []Domain) []data.DomainResponse {
	resp := make([]data.DomainResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
