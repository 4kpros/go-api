package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolConfig struct {
	types.BaseGormModel
	EmailDomain string `gorm:"default:null"`
}

func (item *SchoolConfig) ToResponse() *data.SchoolConfigResponse {
	resp := &data.SchoolConfigResponse{}
	if item == nil {
		return resp
	}
	resp.EmailDomain = item.EmailDomain
	return resp
}

func FromConfigRequest(item *data.SchoolConfigRequest) *SchoolConfig {
	resp := &SchoolConfig{
		EmailDomain: item.EmailDomain,
	}
	return resp
}
