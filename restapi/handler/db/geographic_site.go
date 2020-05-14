/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func GetGeographicSite(db *gorm.DB, id string) (*models.GeographicSite, error) {
	address := &schema.GeographicSite{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).First(address).Error; err != nil {
		return nil, err
	} else {
		return address.To(), nil
	}
}

func GetGeographicSiteByParams(db *gorm.DB, params *gs.GeographicSiteFindParams) ([]*schema.GeographicSite, error) {
	var r []*schema.GeographicSite

	tx := db.Set(AutoPreLoad, true)
	filter := make(map[string]interface{}, 0)

	// query by best match
	var sites []*schema.GeographicSite
	if err := tx.Where(&schema.GeographicSite{
		SiteCompanyName:  swag.StringValue(params.SiteCompanyName),
		SiteCustomerName: swag.StringValue(params.SiteCustomerName),
		SiteName:         swag.StringValue(params.SiteName),
		SiteType:         swag.StringValue(params.SiteType),
		Status:           models.Status(swag.StringValue(params.Status)),
	}).Find(&sites).Error; err == nil {
		for _, site := range sites {
			if _, ok := filter[site.ID]; !ok {
				r = append(r, site)
				filter[site.ID] = struct{}{}
			}
		}
	}
	if params.GeographicAddressCity == nil && params.GeographicAddressCountry == nil && params.GeographicAddressID == nil &&
		params.GeographicAddressPostcode == nil && params.GeographicAddressStreetName == nil && params.GeographicAddressStreetNr == nil {
		log.Warnln("all geographic address are emptry, ignore...")
	} else {
		var fa []*schema.FieldedAddress
		if err := tx.Find(&fa, &schema.FieldedAddress{
			City:       swag.StringValue(params.GeographicAddressCity),
			Country:    swag.StringValue(params.GeographicAddressCountry),
			ID:         swag.StringValue(params.GeographicAddressID),
			Postcode:   swag.StringValue(params.GeographicAddressPostcode),
			StreetName: swag.StringValue(params.GeographicAddressStreetName),
			StreetNr:   swag.StringValue(params.GeographicAddressStreetNr),
		}).Error; err == nil {
			var ids []string
			for _, f := range fa {
				ids = append(ids, f.ID)
			}
			var sites []*schema.GeographicSite
			if err := tx.Where("id IN (?)", ids).Find(&sites).Error; err == nil {
				for _, site := range sites {
					if _, ok := filter[site.ID]; !ok {
						r = append(r, site)
						filter[site.ID] = struct{}{}
					}
				}
			}
		}
	}

	if len(r) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return r, nil
}
