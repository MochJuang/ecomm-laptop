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

type LoginRequest struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}
type RegisterRequest struct {
	Name            string `json:"name" form:"name" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required"`
	Password        string `json:"password" form:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required"`
}

type UserToken struct {
	Token string
	Name  string
}
