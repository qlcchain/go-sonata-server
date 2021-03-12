// Code generated by go-swagger; DO NOT EDIT.

package geographic_site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// GeographicSiteGetHandlerFunc turns a function with the right signature into a geographic site get handler
type GeographicSiteGetHandlerFunc func(GeographicSiteGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GeographicSiteGetHandlerFunc) Handle(params GeographicSiteGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GeographicSiteGetHandler interface for that can handle valid geographic site get params
type GeographicSiteGetHandler interface {
	Handle(GeographicSiteGetParams, *models.Principal) middleware.Responder
}

// NewGeographicSiteGet creates a new http.Handler for the geographic site get operation
func NewGeographicSiteGet(ctx *middleware.Context, handler GeographicSiteGetHandler) *GeographicSiteGet {
	return &GeographicSiteGet{Context: ctx, Handler: handler}
}

/* GeographicSiteGet swagger:route GET /geographicSiteManagement/v3/geographicSite/{SiteId} GeographicSite geographicSiteGet

retrieveSite

The Buyer requests the full details for a single Service Site based on a Service Site identifier that was previously provided by the Seller.

*/
type GeographicSiteGet struct {
	Context *middleware.Context
	Handler GeographicSiteGetHandler
}

func (o *GeographicSiteGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGeographicSiteGetParams()
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
