// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NotificationProductOfferingQualificationStateChangeNotificationHandlerFunc turns a function with the right signature into a notification product offering qualification state change notification handler
type NotificationProductOfferingQualificationStateChangeNotificationHandlerFunc func(NotificationProductOfferingQualificationStateChangeNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NotificationProductOfferingQualificationStateChangeNotificationHandlerFunc) Handle(params NotificationProductOfferingQualificationStateChangeNotificationParams) middleware.Responder {
	return fn(params)
}

// NotificationProductOfferingQualificationStateChangeNotificationHandler interface for that can handle valid notification product offering qualification state change notification params
type NotificationProductOfferingQualificationStateChangeNotificationHandler interface {
	Handle(NotificationProductOfferingQualificationStateChangeNotificationParams) middleware.Responder
}

// NewNotificationProductOfferingQualificationStateChangeNotification creates a new http.Handler for the notification product offering qualification state change notification operation
func NewNotificationProductOfferingQualificationStateChangeNotification(ctx *middleware.Context, handler NotificationProductOfferingQualificationStateChangeNotificationHandler) *NotificationProductOfferingQualificationStateChangeNotification {
	return &NotificationProductOfferingQualificationStateChangeNotification{Context: ctx, Handler: handler}
}

/* NotificationProductOfferingQualificationStateChangeNotification swagger:route POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationStateChangeNotification Notification notificationProductOfferingQualificationStateChangeNotification

Product Offering Qualification State Change Notification structure

Product Offering Qualification State Change Notification structure definition

*/
type NotificationProductOfferingQualificationStateChangeNotification struct {
	Context *middleware.Context
	Handler NotificationProductOfferingQualificationStateChangeNotificationHandler
}

func (o *NotificationProductOfferingQualificationStateChangeNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNotificationProductOfferingQualificationStateChangeNotificationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
