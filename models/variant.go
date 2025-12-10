package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid"`
	Name      string
}
type ListVariant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid"`
	Product   Product
	Name      string
}

// Auto generate UUID before insert
func (p *Variant) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
func (ListVariant) TableName() string{
	return  "variants"
}
