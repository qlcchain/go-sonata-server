package mock

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"
)

func ProductProductFindHandler(params product.ProductFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return product.NewProductFindBadRequest().WithPayload(payload)
	}

	var summaryList []*models.ProductSummary
	var products []schema.Product
	// TODO: query from db
	if err := Store.Set(db.AutoPreLoad, true).Where("", params).Find(&products).Error; err == nil {
		for _, p := range products {
			summaryList = append(summaryList, &models.ProductSummary{
				BuyerProductID:  *params.BuyerID,
				Href:            p.Href,
				ID:              p.ID,
				ProductOffering: p.ProductOffering,
				ProductSpecification: &models.ProductSpecificationSummary{
					ID: p.ProductSpecification.ID,
				},
				StartDate: *p.StartDate,
				Status:    p.Status,
			})
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
		return product.NewProductGetBadRequest().WithPayload(payload)
	}

	if p, err := db.GetProduct(Store, params.ProductID); err == nil {
		return product.NewProductGetOK().WithPayload(p)
	} else if err == gorm.ErrRecordNotFound {
		return product.NewProductGetNotFound()
	} else {
		return product.NewProductFindServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}
