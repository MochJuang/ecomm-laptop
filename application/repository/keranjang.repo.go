package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type KeranjangRepository interface {
	Insert(keranjang model.Keranjang) (model.Keranjang, error)
	Find(userId uint64) ([]model.Keranjang, error)
}

type keranjangConnection struct {
	connection *gorm.DB
}

func NewKeranjangRepository(db *gorm.DB) KeranjangRepository {
	return &keranjangConnection{
		connection: db,
	}
}

func (db *keranjangConnection) Insert(keranjang model.Keranjang) (model.Keranjang, error) {
	var isExist model.Keranjang

	res := db.connection.Where("user_id=? AND product_id=?", keranjang.UserId, keranjang.ProductId).First(&isExist)

	if res.RowsAffected > 0 {
		isExist.Qty = keranjang.Qty + isExist.Qty
		db.connection.Save(&isExist)
	} else {
		result := db.connection.Create(&keranjang)
		if result.Error != nil {
			logrus.Info(result.Error.Error())
			return keranjang, errors.New(result.Error.Error())
		}

		if result.RowsAffected == 0 {
			return model.Keranjang{}, errors.New("Cant add cart")
		}
	}

	return keranjang, nil
}

func (db *keranjangConnection) Find(userId uint64) ([]model.Keranjang, error) {
	var keranjangs []model.Keranjang
	result := db.connection.Preload("User").Preload("Product").Preload("Product.Merk").Preload("Product.Memory").Preload("Product.Warna").Where("user_id=?", userId).Find(&keranjangs)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return keranjangs, errors.New(result.Error.Error())
	}
	return keranjangs, nil
}
