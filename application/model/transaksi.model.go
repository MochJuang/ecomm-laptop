package model

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type StatusTransaksi string

const (
	Wait   StatusTransaksi = "wait"
	Send   StatusTransaksi = "send"
	Finish StatusTransaksi = "finish"
	Cancel StatusTransaksi = "cancel"
)

func (e *StatusTransaksi) Scan(value interface{}) error {
	*e = StatusTransaksi(value.([]byte))
	return nil
}

func (e StatusTransaksi) Value() (driver.Value, error) {
	return string(e), nil
}

type Transaksi struct {
	gorm.Model
	ID              uint64          `gorm:"primaryKey:autoIncrement" json:"id"`
	UserId          uint64          `gorm:"not null" json:"-"`
	User            User            `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	KodeTransaksi   string          `json:"kode_transaksi"`
	AlamatId        uint64          `gorm:"not null" json:"-"`
	Alamat          Alamat          `gorm:"foreignkey:AlamatId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"alamat"`
	Total           uint64          `json:"total"`
	StatusTransaksi StatusKeranjang `json:"StatusKeranjang" sql:"type:ENUM('wait', 'send', 'finish', 'cancel')"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
