// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// ProductOfferingQualificationManagementHubGetHandlerFunc turns a function with the right signature into a product offering qualification management hub get handler
type ProductOfferingQualificationManagementHubGetHandlerFunc func(ProductOfferingQualificationManagementHubGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ProductOfferingQualificationManagementHubGetHandlerFunc) Handle(params ProductOfferingQualificationManagementHubGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ProductOfferingQualificationManagementHubGetHandler interface for that can handle valid product offering qualification management hub get params
type ProductOfferingQualificationManagementHubGetHandler interface {
	Handle(ProductOfferingQualificationManagementHubGetParams, *models.Principal) middleware.Responder
}

// NewProductOfferingQualificationManagementHubGet creates a new http.Handler for the product offering qualification management hub get operation
func NewProductOfferingQualificationManagementHubGet(ctx *middleware.Context, handler ProductOfferingQualificationManagementHubGetHandler) *ProductOfferingQualificationManagementHubGet {
	return &ProductOfferingQualificationManagementHubGet{Context: ctx, Handler: handler}
}

/* ProductOfferingQualificationManagementHubGet swagger:route GET /productOfferingQualificationManagement/v3/hub Hub productOfferingQualificationManagementHubGet

hubFind

This operation retrieves a set of hubs.

*/
type ProductOfferingQualificationManagementHubGet struct {
	Context *middleware.Context
	Handler ProductOfferingQualificationManagementHubGetHandler
}

func (o *ProductOfferingQualificationManagementHubGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProductOfferingQualificationManagementHubGetParams()
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
