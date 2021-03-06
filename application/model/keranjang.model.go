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
	User            User            `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id" json:"user"`
	ProductId       uint64          `gorm:"not null" json:"-"`
	Product         Product         `gorm:"foreignkey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id" json:"product"`
	Qty             int             `json:"qty"`
	KodeCheckout    string          `json:"kode_checkout"`
	StatusKeranjang StatusKeranjang `json:"StatusKeranjang" sql:"type:ENUM('save', 'checkout', 'pay')"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type KeranjangAddRequest struct {
	UserId    uint64 `json:"user_id" xml:"userId" form:"userId" validate:"required"`
	ProductId uint64 `json:"product_id" xml:"productId" form:"productId" validate:"required"`
	Qty       int    `json:"qty" xml:"qty" form:"qty" validate:"required"`
}
