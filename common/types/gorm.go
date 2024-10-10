package types

import (
	"time"

	"gorm.io/gorm"
)

type BaseGormModel struct {
	ID        uint           `gorm:"primaryKey; not null" json:"id" required:"false"`
	CreatedAt time.Time      `json:"createdAt" required:"false"`
	UpdatedAt time.Time      `json:"updatedAt" required:"false"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" required:"false"`
}
