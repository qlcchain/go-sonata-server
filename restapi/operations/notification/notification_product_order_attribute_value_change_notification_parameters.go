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

// NewNotificationProductOrderAttributeValueChangeNotificationParams creates a new NotificationProductOrderAttributeValueChangeNotificationParams object
// no default values defined in spec.
func NewNotificationProductOrderAttributeValueChangeNotificationParams() NotificationProductOrderAttributeValueChangeNotificationParams {

	return NotificationProductOrderAttributeValueChangeNotificationParams{}
}

// NotificationProductOrderAttributeValueChangeNotificationParams contains all the bound params for the notification product order attribute value change notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters notificationProductOrderAttributeValueChangeNotification
type NotificationProductOrderAttributeValueChangeNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	ProductOrderAttributeValueChange *models.EventPlus
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNotificationProductOrderAttributeValueChangeNotificationParams() beforehand.
func (o *NotificationProductOrderAttributeValueChangeNotificationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.EventPlus
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("productOrderAttributeValueChange", "body"))
			} else {
				res = append(res, errors.NewParseError("productOrderAttributeValueChange", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ProductOrderAttributeValueChange = &body
			}
		}
	} else {
		res = append(res, errors.Required("productOrderAttributeValueChange", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
