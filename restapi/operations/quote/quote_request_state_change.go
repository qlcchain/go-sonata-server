// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// QuoteRequestStateChangeHandlerFunc turns a function with the right signature into a quote request state change handler
type QuoteRequestStateChangeHandlerFunc func(QuoteRequestStateChangeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn QuoteRequestStateChangeHandlerFunc) Handle(params QuoteRequestStateChangeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// QuoteRequestStateChangeHandler interface for that can handle valid quote request state change params
type QuoteRequestStateChangeHandler interface {
	Handle(QuoteRequestStateChangeParams, *models.Principal) middleware.Responder
}

// NewQuoteRequestStateChange creates a new http.Handler for the quote request state change operation
func NewQuoteRequestStateChange(ctx *middleware.Context, handler QuoteRequestStateChangeHandler) *QuoteRequestStateChange {
	return &QuoteRequestStateChange{Context: ctx, Handler: handler}
}

/* QuoteRequestStateChange swagger:route POST /quoteManagement/v2/quote/requestStateChange Quote quoteRequestStateChange

request a quote state change

Request from buyer to cancel or reject a quote.
When seller receive cancel request, seller will shift quote state to CANCELLED (no change on order item state)
When seller receive reject request, seller will shift quote state to REJECTED (no change on order item state)

*/
type QuoteRequestStateChange struct {
	Context *middleware.Context
	Handler QuoteRequestStateChangeHandler
}

func (o *QuoteRequestStateChange) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewQuoteRequestStateChangeParams()
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
