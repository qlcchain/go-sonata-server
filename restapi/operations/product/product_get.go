// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProductGetHandlerFunc turns a function with the right signature into a product get handler
type ProductGetHandlerFunc func(ProductGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProductGetHandlerFunc) Handle(params ProductGetParams) middleware.Responder {
	return fn(params)
}

// ProductGetHandler interface for that can handle valid product get params
type ProductGetHandler interface {
	Handle(ProductGetParams) middleware.Responder
}

// NewProductGet creates a new http.Handler for the product get operation
func NewProductGet(ctx *middleware.Context, handler ProductGetHandler) *ProductGet {
	return &ProductGet{Context: ctx, Handler: handler}
}

/*ProductGet swagger:route GET /productInventoryManagement/v3/product/{ProductId} Product productGet

productGet (by id) - retrieve one product with all information

The Buyer requests the details associated with a single Product based on a Seller Product Identifier.

*/
type ProductGet struct {
	Context *middleware.Context
	Handler ProductGetHandler
}

func (o *ProductGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProductGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
