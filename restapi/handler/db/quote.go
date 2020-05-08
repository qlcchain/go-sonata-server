package db

import (
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func GetQuote(db *gorm.DB, id string) (models.Quote, error) {
	q := &schema.Quote{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).First(q).Error; err != nil {
		return nil, err
	} else {
		return q, nil
	}
}
