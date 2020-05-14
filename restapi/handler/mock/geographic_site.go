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

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
)

func GeographicSiteGeographicSiteFindHandler(params gs.GeographicSiteFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gs.NewGeographicSiteFindUnauthorized().WithPayload(payload)
	}

	if sites, err := db.GetGeographicSiteByParams(Store, &params); err == nil {
		var payload []*models.GeographicSiteFindResp
		for _, site := range sites {
			payload = append(payload, &models.GeographicSiteFindResp{
				GeographicAddress: &models.GeographicAddressFindResp{
					FieldedAddress:   site.FieldedAddress.To(),
					FormattedAddress: site.FormattedAddress.To(),
				},
				ID:               site.ID,
				SiteCompanyName:  site.SiteCompanyName,
				SiteContactName:  site.ContractName(),
				SiteCustomerName: site.SiteCustomerName,
				SiteDescription:  site.Description,
				SiteName:         site.SiteName,
				SiteType:         site.SiteType,
				Status:           site.Status,
			})
		}
		return gs.NewGeographicSiteFindOK().WithPayload(payload)
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
