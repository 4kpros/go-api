package types

import (
	"time"
)

type BaseGormModel struct {
	ID        int64 `gorm:"primaryKey; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BaseGormModelResponse struct {
	ID        int64     `json:"id" required:"false"`
	CreatedAt time.Time `json:"createdAt" required:"false"`
	UpdatedAt time.Time `json:"updatedAt" required:"false"`
}
