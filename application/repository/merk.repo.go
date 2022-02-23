package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MerkRepository interface {
	Insert(merk model.Merk) (model.Merk, error)
	FindAndCount(filter model.FilterMerk) (int, error)
	FindAndPaginate(filter model.FilterMerk, limit int, offset int) ([]model.ResultMerk, error)
	DetailMerk(id uint64) (model.DetailMerk, error)
}

type MerkConnection struct {
	connection *gorm.DB
}

func NewMerkRepository(db *gorm.DB) MerkRepository {
	return &MerkConnection{
		connection: db,
	}
}

func (db *MerkConnection) Insert(merk model.Merk) (model.Merk, error) {
	result := db.connection.Create(&merk)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return merk, errors.New(result.Error.Error())
	}

	return merk, nil
}

func (db *MerkConnection) FindAndCount(filter model.FilterMerk) (int, error) {

	var sql string = `
		SELECT
			COUNT(*) as totalData 
		FROM
			(
			SELECT
				COUNT(merks.id) 
			FROM
				merks
				INNER JOIN products ON products.merk_id = merks.id
				INNER JOIN memories ON products.memory_id = memories.id
				INNER JOIN warnas ON products.warna_id = warnas.id 
	`
	if filter.Search != "" || filter.BrandId != 0 || filter.Disk != "" || filter.Ram != 0 {
		sql += "WHERE"
	}

	if filter.Search != "" {
		sql += "  merks.merk LIKE '%" + filter.Search + "%'"
	}

	if filter.BrandId != 0 {
		sql += " AND  merks.brand_id = " + strconv.FormatUint(filter.BrandId, 10)
	}

	if filter.Disk != "" {
		sql += " AND  memories.disk LIKE '%" + filter.Disk + "%'"
	}

	if filter.Ram != 0 {
		sql += " AND  memories.ram = " + strconv.Itoa(filter.Ram)
	}

	sql += ` GROUP BY merks.id 
				) AS a`

	log.Println(sql)

	var totalData int
	db.connection.Raw(sql).Scan(&totalData)

	return totalData, nil
}

func (db *MerkConnection) FindAndPaginate(filter model.FilterMerk, limit int, offset int) ([]model.ResultMerk, error) {
	var result []model.ResultMerk

	var sql string = `
		SELECT
			merks.id,
			merks.merk,
			products.harga,
			warnas.warna,
			memories.disk,
			memories.ram,
			memories.is_ssd,
			merks.created_at	
		FROM
			merks
			INNER JOIN products ON products.merk_id = merks.id
			INNER JOIN memories ON products.memory_id = memories.id
			INNER JOIN warnas ON products.warna_id = warnas.id 
	`
	if filter.Search != "" || filter.BrandId != 0 || filter.Disk != "" || filter.Ram != 0 {
		sql += "WHERE"
	}

	if filter.Search != "" {
		sql += "  merks.merk LIKE '%" + filter.Search + "%'"
	}

	if filter.BrandId != 0 {
		if filter.Search != "" {
			sql += " AND"
		}
		sql += "  merks.brand_id = " + strconv.FormatUint(filter.BrandId, 10)
	}

	if filter.Disk != "" {
		if filter.Search != "" || filter.BrandId != 0 {
			sql += " AND"
		}
		sql += "  memories.disk LIKE '%" + filter.Disk + "%'"
	}

	if filter.Ram != 0 {
		if filter.Search != "" || filter.BrandId != 0 || filter.Disk != "" {
			sql += " AND"
		}
		sql += "  memories.ram = " + strconv.Itoa(filter.Ram)
	}

	sql += " GROUP BY merks.id LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa(offset)

	fmt.Println(sql)

	db.connection.Raw(sql).Scan(&result)

	return result, nil
}

func (db *MerkConnection) DetailMerk(id uint64) (model.DetailMerk, error) {
	var merk model.Merk

	db.connection.Where("id = ?", id).Find(&merk)

	var warnas []*model.Warna

	db.connection.Raw(`
		SELECT
			*
		FROM
			products a
		INNER JOIN warnas b ON a.warna_id = b.id
		WHERE
			merk_id = ` + strconv.FormatUint(id, 10) + `
		GROUP BY warna_id
	`).Scan(&warnas)

	var detailWarnas []model.DetailVariasiWarna
	for _, warna := range warnas {
		var memories []model.DetailVariasiMemory
		db.connection.Raw(`
			SELECT
				b.id as memory_id, a.harga, b.disk, b.ram, b.is_ssd 
			FROM
				products a
				INNER JOIN memories b ON a.memory_id = b.id 
			WHERE
				a.merk_id = ` + strconv.FormatUint(warna.ID, 10) + `
				AND a.warna_id = ` + strconv.FormatUint(id, 10) + `
		`).Scan(&memories)
		detailWarnas = append(detailWarnas, model.DetailVariasiWarna{
			WarnaId:       warna.ID,
			Warna:         warna.Warna,
			VariasiMemory: memories,
		})
	}

	return model.DetailMerk{
		Detail:       merk,
		VariasiWarna: detailWarnas,
	}, nil
}
