// Code generated by go-swagger; DO NOT EDIT.

package geographic_site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// GeographicSiteFindHandlerFunc turns a function with the right signature into a geographic site find handler
type GeographicSiteFindHandlerFunc func(GeographicSiteFindParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GeographicSiteFindHandlerFunc) Handle(params GeographicSiteFindParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GeographicSiteFindHandler interface for that can handle valid geographic site find params
type GeographicSiteFindHandler interface {
	Handle(GeographicSiteFindParams, *models.Principal) middleware.Responder
}

// NewGeographicSiteFind creates a new http.Handler for the geographic site find operation
func NewGeographicSiteFind(ctx *middleware.Context, handler GeographicSiteFindHandler) *GeographicSiteFind {
	return &GeographicSiteFind{Context: ctx, Handler: handler}
}

/*GeographicSiteFind swagger:route GET /geographicSiteManagement/v3/geographicSite GeographicSite geographicSiteFind

retrieveGeographicSites

The Buyer requests that the Seller provides a list of Service Sites known to the Seller based on a set of Site/Address filter criteria.  For each Service Site returned, the Seller also provides a Service Site Identifier, which uniquely identifies this Service Site within the Seller.

*/
type GeographicSiteFind struct {
	Context *middleware.Context
	Handler GeographicSiteFindHandler
}

func (o *GeographicSiteFind) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGeographicSiteFindParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
