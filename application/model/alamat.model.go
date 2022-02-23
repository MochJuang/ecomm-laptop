package model

import (
	"time"

	"gorm.io/gorm"
)

type Alamat struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	UserId      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:id" json:"user"`
	ProvinceId  int    `json:"province_id"`
	CityId      int    `json:"city_id"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
