package model

import (
	"api/common/types"
	"api/services/school/secondary/class/data"
)

type Class struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (class *Class) ToResponse() *data.ClassResponse {
	resp := &data.ClassResponse{}
	resp.ID = class.ID
	resp.CreatedAt = class.CreatedAt
	resp.UpdatedAt = class.UpdatedAt
	resp.DeletedAt = class.DeletedAt

	resp.Name = class.Name
	resp.Description = class.Description
	return resp
}

func ToResponseList(classList []Class) []data.ClassResponse {
	resp := make([]data.ClassResponse, len(classList))
	for index, class := range classList {
		resp[index] = *class.ToResponse()
	}
	return resp
}
