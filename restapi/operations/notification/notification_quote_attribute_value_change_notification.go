// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NotificationQuoteAttributeValueChangeNotificationHandlerFunc turns a function with the right signature into a notification quote attribute value change notification handler
type NotificationQuoteAttributeValueChangeNotificationHandlerFunc func(NotificationQuoteAttributeValueChangeNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NotificationQuoteAttributeValueChangeNotificationHandlerFunc) Handle(params NotificationQuoteAttributeValueChangeNotificationParams) middleware.Responder {
	return fn(params)
}

// NotificationQuoteAttributeValueChangeNotificationHandler interface for that can handle valid notification quote attribute value change notification params
type NotificationQuoteAttributeValueChangeNotificationHandler interface {
	Handle(NotificationQuoteAttributeValueChangeNotificationParams) middleware.Responder
}

// NewNotificationQuoteAttributeValueChangeNotification creates a new http.Handler for the notification quote attribute value change notification operation
func NewNotificationQuoteAttributeValueChangeNotification(ctx *middleware.Context, handler NotificationQuoteAttributeValueChangeNotificationHandler) *NotificationQuoteAttributeValueChangeNotification {
	return &NotificationQuoteAttributeValueChangeNotification{Context: ctx, Handler: handler}
}

/*NotificationQuoteAttributeValueChangeNotification swagger:route POST /quoteNotification/v1/quoteNotification/v1/notification/quoteAttributeValueChangeNotification Notification notificationQuoteAttributeValueChangeNotification

Quote attribute value change notification structure

Quote attribute value change notification structure description.
Attribute resourcePatch allows to target quote but also quoteItem - example: resourcePath":"/quote/42/quoteItem/3" is the item #3 of quote #42
Attribute fieldPath allows to target attribute with value changed.

Specific business errors for current operation will be encapsulated in

HTTP Response 422 Unprocessable entity


*/
type NotificationQuoteAttributeValueChangeNotification struct {
	Context *middleware.Context
	Handler NotificationQuoteAttributeValueChangeNotificationHandler
}

func (o *NotificationQuoteAttributeValueChangeNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNotificationQuoteAttributeValueChangeNotificationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
