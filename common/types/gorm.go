package types

import (
	"time"

	"gorm.io/gorm"
)

type BaseGormModel struct {
	ID        int64 `gorm:"primaryKey; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type BaseGormModelResponse struct {
	ID        int64          `json:"id" required:"false"`
	CreatedAt time.Time      `json:"createdAt" required:"false"`
	UpdatedAt time.Time      `json:"updatedAt" required:"false"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" required:"false"`
}
