// Code generated by go-swagger; DO NOT EDIT.

package product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProductOrderCreateHandlerFunc turns a function with the right signature into a product order create handler
type ProductOrderCreateHandlerFunc func(ProductOrderCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProductOrderCreateHandlerFunc) Handle(params ProductOrderCreateParams) middleware.Responder {
	return fn(params)
}

// ProductOrderCreateHandler interface for that can handle valid product order create params
type ProductOrderCreateHandler interface {
	Handle(ProductOrderCreateParams) middleware.Responder
}

// NewProductOrderCreate creates a new http.Handler for the product order create operation
func NewProductOrderCreate(ctx *middleware.Context, handler ProductOrderCreateHandler) *ProductOrderCreate {
	return &ProductOrderCreate{Context: ctx, Handler: handler}
}

/*ProductOrderCreate swagger:route POST /productOrderManagement/v3/productOrder ProductOrder productOrderCreate

Create a product order

This operation is used to create an order. Depending on the order activity, one can "INSTALL", "CHANGE", or "DISCONNECT" an associated product.

*/
type ProductOrderCreate struct {
	Context *middleware.Context
	Handler ProductOrderCreateHandler
}

func (o *ProductOrderCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProductOrderCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
