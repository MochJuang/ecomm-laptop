package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey:autoIncrement" json:"id"`

	MerkId    uint64 `gorm:"not null" json:"-"`
	Merk      Merk   `gorm:"foreignkey:MerkId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"merk"`
	MemoryId  uint64 `gorm:"not null" json:"-"`
	Memory    Memory `gorm:"foreignkey:MemoryId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"memory"`
	WarnaId   uint64 `gorm:"not null" json:"-"`
	Warna     Warna  `gorm:"foreignkey:WarnaId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"warna"`
	Stock     uint16 `json:"stock"`
	Harga     uint64 `json:"harga"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
