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

func (domain *Domain) ToResponse() *data.DomainResponse {
	resp := &data.DomainResponse{}
	resp.ID = domain.ID
	resp.CreatedAt = domain.CreatedAt
	resp.UpdatedAt = domain.UpdatedAt
	resp.DeletedAt = domain.DeletedAt

	resp.DepartmentID = domain.DepartmentID
	resp.Name = domain.Name
	resp.Description = domain.Description
	return resp
}

func ToResponseList(domainList []Domain) []data.DomainResponse {
	resp := make([]data.DomainResponse, len(domainList))
	for index, domain := range domainList {
		resp[index] = *domain.ToResponse()
	}
	return resp
}
