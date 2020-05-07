package mock

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
)

func GeographicSiteGeographicSiteFindHandler(params gs.GeographicSiteFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gs.NewGeographicSiteFindBadRequest().WithPayload(payload)
	}

	return middleware.NotImplemented("operation geographic_site.GeographicSiteFind has not yet been implemented")
}

func GeographicSiteGeographicSiteGetHandler(params gs.GeographicSiteGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gs.NewGeographicSiteGetBadRequest().WithPayload(payload)
	}
	
	return middleware.NotImplemented("operation geographic_site.GeographicSiteGet has not yet been implemented")
}
