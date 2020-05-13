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

	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
)

func GeographicSiteGeographicSiteFindHandler(params gs.GeographicSiteFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gs.NewGeographicSiteFindUnauthorized().WithPayload(payload)
	}

	var sites []*schema.GeographicSite
	//TODO: query from db
	tx := Store.Set(db.AutoPreLoad, true)
	if err := tx.Where("", params).Find(&sites).Error; err == nil {
		return gs.NewGeographicSiteFindOK()
	} else if err == gorm.ErrRecordNotFound {
		return gs.NewGeographicSiteFindNotFound()
	} else {
		return gs.NewGeographicSiteFindInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func GeographicSiteGeographicSiteGetHandler(params gs.GeographicSiteGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gs.NewGeographicSiteGetUnauthorized().WithPayload(payload)
	}

	if site, err := db.GetGeographicSite(Store, params.SiteID); err == nil {
		return gs.NewGeographicSiteGetOK().WithPayload(site)
	} else if err == gorm.ErrRecordNotFound {
		return gs.NewGeographicSiteGetNotFound()
	} else {
		return gs.NewGeographicSiteGetInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}
