// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProductOfferingQualificationManagementHubPostHandlerFunc turns a function with the right signature into a product offering qualification management hub post handler
type ProductOfferingQualificationManagementHubPostHandlerFunc func(ProductOfferingQualificationManagementHubPostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProductOfferingQualificationManagementHubPostHandlerFunc) Handle(params ProductOfferingQualificationManagementHubPostParams) middleware.Responder {
	return fn(params)
}

// ProductOfferingQualificationManagementHubPostHandler interface for that can handle valid product offering qualification management hub post params
type ProductOfferingQualificationManagementHubPostHandler interface {
	Handle(ProductOfferingQualificationManagementHubPostParams) middleware.Responder
}

// NewProductOfferingQualificationManagementHubPost creates a new http.Handler for the product offering qualification management hub post operation
func NewProductOfferingQualificationManagementHubPost(ctx *middleware.Context, handler ProductOfferingQualificationManagementHubPostHandler) *ProductOfferingQualificationManagementHubPost {
	return &ProductOfferingQualificationManagementHubPost{Context: ctx, Handler: handler}
}

/*ProductOfferingQualificationManagementHubPost swagger:route POST /productOfferingQualificationManagement/v3/hub Hub productOfferingQualificationManagementHubPost

hubCreate

A request initiated by the Buyer to instruct the Seller to send notifications of POQ state changes in the event the Seller uses the Deferred Response pattern to respond to a Create Product Offering Qualifica-tion request.

*/
type ProductOfferingQualificationManagementHubPost struct {
	Context *middleware.Context
	Handler ProductOfferingQualificationManagementHubPostHandler
}

func (o *ProductOfferingQualificationManagementHubPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProductOfferingQualificationManagementHubPostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
