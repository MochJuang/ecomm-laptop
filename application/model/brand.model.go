package model

import (
	"time"

	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID        uint64  `gorm:"primaryKey:autoIncrement" json:"id"`
	Brand     string  `json:"brand"`
	Image     string  `json:"image"`
	Merk      []*Merk `json:"merk,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
