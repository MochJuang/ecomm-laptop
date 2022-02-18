package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64       `gorm:"primaryKey:autoIncrement" json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Token     string       `json:"token"`
	Alamat    []*Alamat    `json:"alamat,omitempty"`
	Keranjang []*Keranjang `json:"keranjang,omitempty"`
	Transaksi []*Transaksi `json:"transaksi,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
