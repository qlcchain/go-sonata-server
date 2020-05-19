/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"time"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/util"

	po "github.com/qlcchain/go-sonata-server/restapi/operations/product_order"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
)

func GetProductOrder(id string) (*models.ProductOrder, error) {
	from := &schema.ProductOrder{}
	tx := Store.Set(AutoPreLoad, true)
	if err := tx.Where("id=?", id).First(from).Error; err != nil {
		return nil, err
	}
	return from.To(), nil
}

func GetProductOrderByParams(param *po.ProductOrderFindParams) ([]*models.ProductOrderSummary, error) {
	var r []*models.ProductOrderSummary
	tx := Store.Set(AutoPreLoad, true)

	var ids []*string
	filter := make(map[*string]interface{}, 0)
	//BuyerID *string
	if v, b := handler.VerifyField(param.BuyerID); b {
		var p []*schema.ProductOrder
		if err := tx.Model(&schema.RelatedParty{
			ID: swag.String(v),
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
	//SellerID *string
	if v, b := handler.VerifyField(param.SellerID); b {
		var p []*schema.ProductOrder
		if err := tx.Model(&schema.RelatedParty{
			ID: swag.String(v),
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

	//SiteCompanyName *string
	//SiteCustomerName *string
	//SiteName *string
	//State *string
	var p []*schema.Product
	if err := Store.Model(&schema.GeographicSite{
		SiteCompanyName:  swag.StringValue(param.SiteCompanyName),
		SiteCustomerName: swag.StringValue(param.SiteCustomerName),
		SiteName:         swag.StringValue(param.SiteName),
		Status:           models.Status(swag.StringValue(param.State)),
	}).Related(&p, "ID").Error; err == nil {
		for _, r := range p {
			if _, ok := filter[r.ID]; !ok {
				ids = append(ids, r.ID)
				filter[r.ID] = struct{}{}
			}
		}
	}

	//BuyerRequestedDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.BuyerRequestedDateGt); b {
		tx = tx.Where("buyer_request_date>=?", v)
	}

	//BuyerRequestedDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.BuyerRequestedDateLt); b {
		tx = tx.Where("buyer_request_date<=?", v)
	}

	//CompletionDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.CompletionDateGt); b {
		tx = tx.Where("completion_date<=?", v)
	}
	//CompletionDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.CompletionDateLt); b {
		tx = tx.Where("completion_date<=?", v)
	}
	//ExpectedCompletionDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.ExpectedCompletionDateGt); b {
		tx = tx.Where("expected_completion_date>=?", v)
	}
	//ExpectedCompletionDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.ExpectedCompletionDateLt); b {
		tx = tx.Where("expected_completion_date<=", v)
	}
	//ExternalID *string
	if v, b := handler.VerifyField(param.ExternalID); b {
		tx = tx.Where("external_id=?", v)
	}
	//Limit *string
	if v, b := handler.VerifyField(param.Limit); b {
		tx = tx.Limit(v)
	}
	//Offset *string
	if v, b := handler.VerifyField(param.Offset); b {
		tx = tx.Offset(v)
	}
	//OrderCancellationDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.OrderCancellationDateGt); b {
		tx = tx.Where("cancellation_date>=?", v)
	}
	//OrderCancellationDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.OrderCancellationDateLt); b {
		tx = tx.Where("cancellation_date<=?", v)
	}
	//OrderDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.OrderDateGt); b {
		tx = tx.Where("order_date>=?", v)
	}
	//OrderDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.OrderDateLt); b {
		tx = tx.Where("order_date<=?", v)
	}
	//ProjectID *string
	if v, b := handler.VerifyField(param.ProjectID); b {
		tx = tx.Where("project_id=?", v)
	}
	//RequestedCompletionDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.RequestedCompletionDateGt); b {
		tx = tx.Where("requested_completion_date>=?", v)
	}
	//RequestedCompletionDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.RequestedCompletionDateLt); b {
		tx = tx.Where("requested_completion_date<=?", v)
	}
	//RequestedStartDateGt *strfmt.DateTime
	if v, b := handler.VerifyField(param.RequestedStartDateGt); b {
		tx = tx.Where("requested_start_date>=?", v)
	}
	//RequestedStartDateLt *strfmt.DateTime
	if v, b := handler.VerifyField(param.RequestedStartDateLt); b {
		tx = tx.Where("requested_start_date<=?", v)
	}

	if len(ids) > 0 {
		tx = tx.Where("id IN (?)", ids)
	}

	var pos []*schema.ProductOrder
	if err := tx.Find(&pos).Error; err == nil {
		for _, order := range pos {
			r = append(r, &models.ProductOrderSummary{
				ExternalID: order.ExternalID,
				ID:         order.ID,
				OrderDate:  order.OrderDate,
				State:      order.State,
			})
		}
		return r, nil
	} else {
		return nil, err
	}
}

func ProductOrderCreateToProductOrder(param *models.ProductOrderCreate) (*schema.ProductOrder, error) {
	to := &schema.ProductOrder{}
	if err := util.Convert(param, to); err == nil {
		to.ID = util.NewOrDefaultPtr(to.ID)
		// TODO: how to calculate???
		to.ExpectedCompletionDate = strfmt.DateTime(time.Now().AddDate(0, 0, 5))
		to.State = models.ProductOrderStateTypeAcknowledged
		now := strfmt.DateTime(time.Now())
		to.OrderDate = &now
		to.StateChange = []*schema.StateChange{
			{
				ID:           util.NewIDPtr(),
				ChangeDate:   now,
				ChangeReason: "new",
				State:        models.ProductOfferingQualificationStateTypeInProgress,
			},
		}
		for _, note := range to.Note {
			note.ID = util.NewIDPtr()
		}
		for _, item := range to.OrderItem {
			item.ID = util.NewIDPtr()
			item.State = models.ProductOrderItemStateTypeAcknowledged
			item.StateChange = []*schema.StateChange{
				{
					ID:           util.NewIDPtr(),
					ChangeDate:   now,
					ChangeReason: "new",
					State:        models.ProductOfferingQualificationStateTypeInProgress,
				},
			}
			for _, price := range item.OrderItemPrice {
				price.ID = util.NewIDPtr()
			}
			product := item.Product
			for _, price := range product.ProductPrice {
				price.ID = util.NewIDPtr()
			}
			for _, relationship := range product.ProductRelationship {
				relationship.ID = util.NewIDPtr()
			}
			for _, term := range product.ProductTerm {
				term.ID = util.NewIDPtr()
			}
			for _, status := range product.StatusChange {
				status.ID = util.NewIDPtr()
			}
		}
		return to, nil
	} else {
		return nil, err
	}
}
