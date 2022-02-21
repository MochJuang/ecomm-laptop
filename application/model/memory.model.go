package model

import (
	"time"

	"gorm.io/gorm"
)

type Memory struct {
	gorm.Model
	ID        uint64     `gorm:"primaryKey:autoIncrement" json:"id"`
	Disk      string     `json:"disk"`
	Ram       int        `json:"ram"`
	IsSsd     bool       `json:"is_ssd"`
	Product   []*Product `json:"product,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
