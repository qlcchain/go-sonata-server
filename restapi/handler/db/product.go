package db

import (
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func GetProduct(db *gorm.DB, id string) (*models.Product, error) {
	product := &schema.Product{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).
		First(product).Error; err != nil {
		return nil, err
	} else {
		to := &models.Product{}
		if err := util.Convert(product, to); err == nil {
			return to, nil
		} else {
			return nil, err
		}
	}
}
