// Code generated by go-swagger; DO NOT EDIT.

package geographic_address_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// GeographicAddressValidationCreateHandlerFunc turns a function with the right signature into a geographic address validation create handler
type GeographicAddressValidationCreateHandlerFunc func(GeographicAddressValidationCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GeographicAddressValidationCreateHandlerFunc) Handle(params GeographicAddressValidationCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GeographicAddressValidationCreateHandler interface for that can handle valid geographic address validation create params
type GeographicAddressValidationCreateHandler interface {
	Handle(GeographicAddressValidationCreateParams, *models.Principal) middleware.Responder
}

// NewGeographicAddressValidationCreate creates a new http.Handler for the geographic address validation create operation
func NewGeographicAddressValidationCreate(ctx *middleware.Context, handler GeographicAddressValidationCreateHandler) *GeographicAddressValidationCreate {
	return &GeographicAddressValidationCreate{Context: ctx, Handler: handler}
}

/* GeographicAddressValidationCreate swagger:route POST /geographicAddressManagement/v3/geographicAddressValidation GeographicAddressValidation geographicAddressValidationCreate

validate a Geographic Address

The Buyer sends Address information known to the Buyer to the Seller.  The Seller re-sponds with a list of Addresses known to the Seller that likely match the Address infor-mation sent by the Buyer.  For each Address returned, the Seller generally also provides an Address Identifier, which uniquely identifies this Address within the Seller.

*/
type GeographicAddressValidationCreate struct {
	Context *middleware.Context
	Handler GeographicAddressValidationCreateHandler
}

func (o *GeographicAddressValidationCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGeographicAddressValidationCreateParams()
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
