package model

import (
	"time"

	"gorm.io/gorm"
)

type Warna struct {
	gorm.Model
	ID        uint64     `gorm:"primaryKey:autoIncrement" json:"id"`
	Warna     string     `json:"warna"`
	Product   []*Product `json:"product,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
