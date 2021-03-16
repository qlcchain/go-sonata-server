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

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

var (
	Store *gorm.DB
)

const (
	AutoPreLoad = "gorm:auto_preload"
)

func CreateTables() error {
	if Store != nil {
		return Store.AutoMigrate(
			// user management
			&schema.User{}, &schema.HubSubscriber{},
			// GeographicAddress
			&schema.GeographicAddress{}, &schema.FieldedAddress{}, &schema.FormattedAddress{},
			&schema.GeographicLocation{}, &schema.GeographicSubAddress{}, &schema.GeographicPoint{},
			&schema.SubUnit{}, &models.ReferencedAddress{},
			// GeographicSite
			&schema.GeographicSite{}, &schema.FormattedAddress{}, &schema.FieldedAddress{},
			&schema.GeographicSubAddress{}, &schema.SubUnit{}, &schema.GeographicLocation{}, &schema.GeographicPoint{},
			&models.ReferencedAddress{}, &schema.GeographicSite{},
			// Product
			&schema.Product{}, &models.Agreement{}, &schema.ProductRelationship{}, &models.ProductRef{}, &schema.ProductSpecificationRef{},
			&schema.ProductTerm{}, &models.ProductOfferingRef{}, &models.ProductOfferingQualificationItemRelationship{}, &schema.ProductOrderRef{},
			&schema.ProductPrice{}, &models.TerminationError{}, &models.ProductOfferingQualificationRef{},
			// Order
			&schema.ProductOrder{}, &schema.OrderItem{}, &schema.OrderItemPrice{}, &schema.OrderMessage{}, &models.OrderItemRelationShip{},
			&models.QuoteRef{}, &schema.QualificationRef{},
			// POQ
			&schema.ProductOfferingQualification{}, &schema.ProductOfferingQualificationItem{}, &schema.StatusChange{}, &schema.StateChange{},
			&models.BillingAccountRef{},
			&schema.RelatedParty{}, &schema.AlternateProductProposal{},
			// Quote
			&schema.Quote{}, &models.AgreementRef{}, &schema.Note{}, &schema.QuoteItem{}, &models.QuoteItemRelationship{},
		).Error
	} else {
		return errors.New("store is nil")
	}
}

func GetTx() *gorm.DB {
	return Store.Set(AutoPreLoad, true)
}
