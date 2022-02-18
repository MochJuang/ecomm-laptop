package model

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type StatusKeranjang string

const (
	Save     StatusKeranjang = "save"
	Checkout StatusKeranjang = "checkout"
	Pay      StatusKeranjang = "pay"
)

func (e *StatusKeranjang) Scan(value interface{}) error {
	*e = StatusKeranjang(value.([]byte))
	return nil
}

func (e StatusKeranjang) Value() (driver.Value, error) {
	return string(e), nil
}

type Keranjang struct {
	gorm.Model
	ID              uint64          `gorm:"primaryKey:autoIncrement" json:"id"`
	UserId          uint64          `gorm:"not null" json:"-"`
	User            User            `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	ProductId       uint64          `gorm:"not null" json:"-"`
	Product         Product         `gorm:"foreignkey:ProductId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
	Qty             int             `json:"qty"`
	KodeCheckout    string          `json:"kode_checkout"`
	StatusKeranjang StatusKeranjang `json:"StatusKeranjang" sql:"type:ENUM('save', 'checkout', 'pay')"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
