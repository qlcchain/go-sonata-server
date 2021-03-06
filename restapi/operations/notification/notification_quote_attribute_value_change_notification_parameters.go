// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// NewNotificationQuoteAttributeValueChangeNotificationParams creates a new NotificationQuoteAttributeValueChangeNotificationParams object
//
// There are no default values defined in the spec.
func NewNotificationQuoteAttributeValueChangeNotificationParams() NotificationQuoteAttributeValueChangeNotificationParams {

	return NotificationQuoteAttributeValueChangeNotificationParams{}
}

// NotificationQuoteAttributeValueChangeNotificationParams contains all the bound params for the notification quote attribute value change notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters notificationQuoteAttributeValueChangeNotification
type NotificationQuoteAttributeValueChangeNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	QuoteAttributeValueChangeNotification []*models.QuoteEventPlus
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNotificationQuoteAttributeValueChangeNotificationParams() beforehand.
func (o *NotificationQuoteAttributeValueChangeNotificationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.QuoteEventPlus
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("quoteAttributeValueChangeNotification", "body", ""))
			} else {
				res = append(res, errors.NewParseError("quoteAttributeValueChangeNotification", "body", "", err))
			}
		} else {

			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.QuoteAttributeValueChangeNotification = body
			}
		}
	} else {
		res = append(res, errors.Required("quoteAttributeValueChangeNotification", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
