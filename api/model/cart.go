package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID          string         `json:"id" gorm:"type:varchar(36);primaryKey;not null"`
	KodeProduct string         `json:"kode_product" gorm:"type:varchar(36);not null"`
	NameProduct string         `json:"name_product" gorm:"type:varchar(255);not null"`
	Quantity    int64          `json:"quantity" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	return
}
