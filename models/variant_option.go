package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VariantOption struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	VariantID uuid.UUID
	Value     string
}
type ListVariantOption struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	VariantID uuid.UUID `gorm:"type:uuid"`
	Variant         Variant
	Value            string
}

// Auto generate UUID before insert
func (p *VariantOption) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
func (ListVariantOption) TableName() string {
	return "variant_options"
}
