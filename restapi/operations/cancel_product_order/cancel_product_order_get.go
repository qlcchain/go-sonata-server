// Code generated by go-swagger; DO NOT EDIT.

package cancel_product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// CancelProductOrderGetHandlerFunc turns a function with the right signature into a cancel product order get handler
type CancelProductOrderGetHandlerFunc func(CancelProductOrderGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelProductOrderGetHandlerFunc) Handle(params CancelProductOrderGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CancelProductOrderGetHandler interface for that can handle valid cancel product order get params
type CancelProductOrderGetHandler interface {
	Handle(CancelProductOrderGetParams, *models.Principal) middleware.Responder
}

// NewCancelProductOrderGet creates a new http.Handler for the cancel product order get operation
func NewCancelProductOrderGet(ctx *middleware.Context, handler CancelProductOrderGetHandler) *CancelProductOrderGet {
	return &CancelProductOrderGet{Context: ctx, Handler: handler}
}

/*CancelProductOrderGet swagger:route GET /productOrderManagement/v3/cancelProductOrder/{CancelProductOrderId} CancelProductOrder cancelProductOrderGet

Get a product order cancellation request (by id)

This operation allows buyer to retrieve one product order cancellation request

*/
type CancelProductOrderGet struct {
	Context *middleware.Context
	Handler CancelProductOrderGetHandler
}

func (o *CancelProductOrderGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelProductOrderGetParams()

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
