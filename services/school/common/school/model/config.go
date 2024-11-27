package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolConfig struct {
	types.BaseGormModel
	EmailDomain string `gorm:"default:null"`
}

func (schoolConfig *SchoolConfig) ToResponse() *data.SchoolConfigResponse {
	schoolConfigResp := &data.SchoolConfigResponse{
		EmailDomain: schoolConfig.EmailDomain,
	}
	return schoolConfigResp
}
