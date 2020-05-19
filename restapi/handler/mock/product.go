/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"
)

func ProductProductFindHandler(params product.ProductFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return product.NewProductFindUnauthorized().WithPayload(payload)
	}

	var summaryList []*models.ProductSummary
	if products, err := db.GetProductsByParams(&params); err == nil {
		for _, p := range products {
			item := &models.ProductSummary{
				BuyerProductID:  swag.StringValue(params.BuyerID),
				Href:            p.Href,
				ID:              p.ID,
				ProductOffering: p.ProductOffering,
				StartDate:       *p.StartDate,
				Status:          p.Status,
			}
			if p.ProductSpecification != nil {
				item.ProductSpecification = &models.ProductSpecificationSummary{
					ID: p.ProductSpecification.ID,
				}
			}
			summaryList = append(summaryList, item)
		}
		return product.NewProductFindOK().WithPayload(summaryList)
	} else if err == gorm.ErrRecordNotFound {
		return product.NewProductFindNotFound()
	} else {
		return product.NewProductFindServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func ProductProductGetHandler(params product.ProductGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return product.NewProductGetUnauthorized().WithPayload(payload)
	}

	if p, err := db.GetProduct(params.ProductID); err == nil {
		return product.NewProductGetOK().WithPayload(p)
	} else if err == gorm.ErrRecordNotFound {
		return product.NewProductGetNotFound()
	} else {
		return product.NewProductFindServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}
