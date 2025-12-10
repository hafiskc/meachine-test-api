package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID    int64     `gorm:"unique"`
	ProductCode  string    `gorm:"unique"`
	ProductName  string
	ProductImage string
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time
	CreatedUser  uuid.UUID `gorm:"type:uuid"`
	IsFavourite  bool
	Active       bool
	HSNCode      string
	TotalStock   float64
}

// Auto generate UUID before insert
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
