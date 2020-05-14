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

	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func GetProduct(db *gorm.DB, id string) (*models.Product, error) {
	r := &schema.Product{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).First(r).Error; err != nil {
		return nil, err
	} else {
		return r.To(), nil
	}
}

func GetProductsByParams(db *gorm.DB, params *product.ProductFindParams) ([]*schema.Product, error) {
	tx := db.Set(AutoPreLoad, true)
	var r []*schema.Product
	var ids []*string
	filter := make(map[*string]interface{}, 0)

	// BuyerID
	if _, b := handler.VerifyField(params.BuyerID); b {
		var p []*schema.Product
		if err := tx.Model(&schema.RelatedParty{
			ID: util.NewOrDefaultPtr(params.BuyerID),
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// Status
	if v, b := handler.VerifyField(params.Status); b {
		tx = tx.Where("status=?", v)
	}
	// ProductSpecificationID
	if _, b := handler.VerifyField(params.ProductSpecificationID); b {
		var p []*schema.Product
		if err := tx.Model(&schema.ProductSpecificationRef{
			ID: params.ProductSpecificationID,
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// ProductOfferingID
	if _, b := handler.VerifyField(params.ProductOfferingID); b {
		var p []*schema.Product
		if err := tx.Model(&models.ProductOrderRef{
			ID: params.ProductOfferingID,
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// BuyerProductID
	if v, b := handler.VerifyField(params.BuyerProductID); b {
		tx = tx.Where("buyer_product_id=?", v)
	}
	// GeographicalSiteID
	if _, b := handler.VerifyField(params.GeographicalSiteID); b {
		var p []*schema.Product
		if err := tx.Model(&schema.GeographicSite{
			ID: swag.StringValue(params.GeographicalSiteID),
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// RelatedProductID
	if _, b := handler.VerifyField(params.RelatedProductID); b {
		var p []*schema.Product
		if err := tx.Model(&schema.ProductRelationship{
			ID: params.RelatedProductID,
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// BillingAccountID
	if _, b := handler.VerifyField(params.BillingAccountID); b {
		var p []*schema.Product
		if err := tx.Model(&models.BillingAccountRef{
			ID: params.BillingAccountID,
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// ProductOrderID
	if _, b := handler.VerifyField(params.ProductOrderID); b {
		var p []*schema.Product
		if err := tx.Model(&schema.ProductOrderRef{
			ID: util.NewOrDefaultPtr(params.ProductOrderID),
		}).Related(&p, "ID").Error; err == nil {
			for _, r := range p {
				if _, ok := filter[r.ID]; !ok {
					ids = append(ids, r.ID)
					filter[r.ID] = struct{}{}
				}
			}
		} else {
			return nil, err
		}
	}
	// StartDateLt
	if v, b := handler.VerifyField(params.StartDateLt); b {
		tx = tx.Where("start_date<=?", v)
	}
	// StartDateGt
	if v, b := handler.VerifyField(params.StartDateGt); b {
		tx = tx.Where("start_date>=?", v)
	}
	// LastUpdateDateGt
	if v, b := handler.VerifyField(params.LastUpdateDateGt); b {
		tx = tx.Where("last_update_date>=?", v)
	}
	// LastUpdateDateLt
	if v, b := handler.VerifyField(params.LastUpdateDateLt); b {
		tx = tx.Where("last_update_date<=?", v)
	}
	// Offset
	if v, b := handler.VerifyField(params.Offset); b {
		tx = tx.Offset(v)
	}
	// Limit
	if v, b := handler.VerifyField(params.Limit); b {
		tx = tx.Limit(v)
	}

	if len(ids) > 0 {
		tx = tx.Where("id IN (?)", ids)
	}

	err := tx.Find(&r).Error
	return r, err
}
