package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func ListGeographicAddress(db *gorm.DB) ([]*models.GeographicAddress, error) {
	var addresses []schema.GeographicAddress
	if err := db.Set(AutoPreLoad, true).
		Find(&addresses).Error; err != nil {
		return nil, err
	}

	var result []*models.GeographicAddress
	for _, a := range addresses {
		to := &models.GeographicAddress{}
		if err := util.Convert(a, to); err == nil {
			result = append(result, to)
		} else {
			logrus.Error(err)
		}
	}
	return result, nil
}

func GetGeographicAddress(db *gorm.DB, id string) (*models.GeographicAddress, error) {
	address := &schema.GeographicAddress{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).
		First(address).Error; err != nil {
		return nil, err
	}
	to := &models.GeographicAddress{}
	if err := util.Convert(address, to); err == nil {
		return to, nil
	} else {
		return nil, err
	}
}
