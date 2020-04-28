package mock

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
)

func GeographicSiteGeographicSiteFindHandler(params gs.GeographicSiteFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation geographic_site.GeographicSiteFind has not yet been implemented")
}

func GeographicSiteGeographicSiteGetHandler(params gs.GeographicSiteGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation geographic_site.GeographicSiteGet has not yet been implemented")
}
