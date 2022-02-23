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
	Brand     Brand      `gorm:"foreignkey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:id" json:"brand"`
	Product   []*Product `json:"product,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResultMerk struct {
	ID        uint64
	Merk      string
	Harga     uint64
	Warna     string
	Disk      string
	Ram       int
	IsSsd     bool
	CreatedAt time.Time
}

type FilterMerk struct {
	Search  string
	BrandId uint64
	Disk    string
	Ram     int
}

type DetailMerk struct {
	Detail       Merk
	VariasiWarna []DetailVariasiWarna
}
type DetailVariasiWarna struct {
	WarnaId       uint64
	Warna         string
	VariasiMemory []DetailVariasiMemory
}
type DetailVariasiMemory struct {
	MemoryId uint64
	Harga    string
	Disk     string
	Ram      int
	IsSsd    bool
}
