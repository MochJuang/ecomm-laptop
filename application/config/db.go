package config

import (
	"fmt"
	"os"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	//nanti kita isi modelnya di sini
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.Brand{},
		&model.Merk{},
		&model.Warna{},
		&model.Memory{},
		&model.Product{},
		&model.User{},
		&model.Alamat{},
		&model.Keranjang{},
		&model.Transaksi{},
		&model.Banner{},
	)
	// tx := db.Session(&gorm.Session{Logger: newLogger})
	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
