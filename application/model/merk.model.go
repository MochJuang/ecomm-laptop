package model

import (
	"time"

	"gorm.io/gorm"
)

type Merk struct {
	gorm.Model
	ID        uint64     `gorm:"primaryKey:autoIncrement" json:"id"`
	Merk      string     `json:"merk"`
	BrandId   uint64     `gorm:"not null" json:"-"`
	Brand     Brand      `gorm:"foreignkey:BrandId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"brand"`
	Product   []*Product `json:"product,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
