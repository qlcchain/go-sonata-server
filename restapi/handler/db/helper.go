package db

import (
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/models"
)

func ListGeographicAddress(db *gorm.DB) ([]models.GeographicAddress, error) {
	var addresses []models.GeographicAddress
	if err := db.Set(AutoPreLoad, true).
		Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

func GetGeographicAddress(db *gorm.DB, id string) (*models.GeographicAddress, error) {
	address := &models.GeographicAddress{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).
		First(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}
