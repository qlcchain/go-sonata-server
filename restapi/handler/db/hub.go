package db

import (
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/schema"
)

func FindSubscriber(db *gorm.DB, id, attype string) (*schema.HubSubscriber, error) {
	var r = &schema.HubSubscriber{}
	err := db.Where("id=? AND type=?", id, attype).First(r).Error
	return r, err
}

func ListSubscribers(db *gorm.DB, attype string) ([]*schema.HubSubscriber, error) {
	var r []*schema.HubSubscriber
	err := db.Where("type=?", attype).Find(&r).Error
	return r, err
}

func AddSubscriber(db *gorm.DB, subscriber *schema.HubSubscriber) error {
	return db.Create(subscriber).Error
}

func DeleteSubscriber(db *gorm.DB, id string) error {
	return db.Where("id=?", id).Delete(&schema.HubSubscriber{}).Error
}
