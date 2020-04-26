// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// NewNotificationProductOrderCreationNotificationParams creates a new NotificationProductOrderCreationNotificationParams object
// no default values defined in spec.
func NewNotificationProductOrderCreationNotificationParams() NotificationProductOrderCreationNotificationParams {

	return NotificationProductOrderCreationNotificationParams{}
}

// NotificationProductOrderCreationNotificationParams contains all the bound params for the notification product order creation notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters notificationProductOrderCreationNotification
type NotificationProductOrderCreationNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	ProductOrderCreationNotification models.Event
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNotificationProductOrderCreationNotificationParams() beforehand.
func (o *NotificationProductOrderCreationNotificationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		body, err := models.UnmarshalEvent(r.Body, route.Consumer)
		if err != nil {
			res = append(res, err)
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ProductOrderCreationNotification = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
