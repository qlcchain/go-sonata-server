/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func ListGeographicAddress(db *gorm.DB) ([]*models.GeographicAddress, error) {
	var addresses []schema.GeographicAddress
	if err := db.Set(AutoPreLoad, true).Find(&addresses).Error; err != nil {
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

func GetGeographicAddressByIds(db *gorm.DB, ids []string) ([]*schema.GeographicAddress, error) {
	if len(ids) == 0 {
		return nil, errors.New("invalid ids")
	}
	tx := db.Set(AutoPreLoad, true)
	var address []*schema.GeographicAddress
	err := tx.Where("id IN (?)", ids).Find(&address).Error
	return address, err
}

func GetGeographicAddress(db *gorm.DB, id string) (*models.GeographicAddress, error) {
	address := &schema.GeographicAddress{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).First(address).Error; err != nil {
		return nil, err
	}
	to := &models.GeographicAddress{}
	if err := util.Convert(address, to); err == nil {
		return to, nil
	} else {
		return nil, err
	}
}

func GetGeographicAddressByFieldedAddress(db *gorm.DB, req *models.FieldedAddressRequest) ([]*schema.GeographicAddress, error) {
	var r []*schema.FieldedAddress
	tx := db.Set(AutoPreLoad, true)
	tx = tx.Where(&schema.FieldedAddress{
		City:               req.City,
		Country:            req.Country,
		Locality:           req.Locality,
		PostCodeExtension:  req.PostCodeExtension,
		Postcode:           req.Postcode,
		StateOrProvince:    req.StateOrProvince,
		StreetName:         req.StreetName,
		StreetNr:           req.StreetNr,
		StreetNrLast:       req.StreetNrLast,
		StreetNrLastSuffix: req.StreetNrLastSuffix,
		StreetNrSuffix:     req.StreetNrSuffix,
		StreetSuffix:       req.StreetSuffix,
		StreetType:         req.StreetType,
	})
	if req.GeographicSubAddress != nil {
		var sub []*schema.GeographicSubAddress
		if err := db.Find(&sub, req.GeographicSubAddress).Error; err == nil {
			var ids []string
			for _, address := range sub {
				ids = append(ids, address.ID)
			}
			tx = tx.Where("id IN (?)", ids)
		}
	}
	if err := tx.Find(&r).Error; err == nil {
		var ids []string
		for _, fieldedAddress := range r {
			ids = append(ids, fieldedAddress.ID)
		}
		return GetGeographicAddressByIds(db, ids)
	} else {
		return nil, err
	}
}

func GetGeographicAddressByFormattedAddress(db *gorm.DB, param *models.FormattedAddressRequest) ([]*schema.GeographicAddress, error) {
	fa := request2FormattedAddress(param)
	var r []*schema.FormattedAddress
	if err := db.Where(fa).Find(&r).Error; err == nil {
		var ids []string
		for _, formattedAddress := range r {
			ids = append(ids, formattedAddress.ID)
		}
		return GetGeographicAddressByIds(db, ids)
	} else {
		return nil, err
	}
}

func request2FormattedAddress(req *models.FormattedAddressRequest) *schema.FormattedAddress {
	if req == nil {
		return nil
	}

	to := &schema.FormattedAddress{}
	_ = util.Convert(req, to)
	return to
}
