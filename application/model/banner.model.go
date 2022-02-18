package model

import (
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
