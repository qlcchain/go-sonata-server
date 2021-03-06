// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// NewNotificationProductOfferingQualificationStateChangeNotificationParams creates a new NotificationProductOfferingQualificationStateChangeNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationProductOfferingQualificationStateChangeNotificationParams() *NotificationProductOfferingQualificationStateChangeNotificationParams {
	return &NotificationProductOfferingQualificationStateChangeNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationProductOfferingQualificationStateChangeNotificationParamsWithTimeout creates a new NotificationProductOfferingQualificationStateChangeNotificationParams object
// with the ability to set a timeout on a request.
func NewNotificationProductOfferingQualificationStateChangeNotificationParamsWithTimeout(timeout time.Duration) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	return &NotificationProductOfferingQualificationStateChangeNotificationParams{
		timeout: timeout,
	}
}

// NewNotificationProductOfferingQualificationStateChangeNotificationParamsWithContext creates a new NotificationProductOfferingQualificationStateChangeNotificationParams object
// with the ability to set a context for a request.
func NewNotificationProductOfferingQualificationStateChangeNotificationParamsWithContext(ctx context.Context) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	return &NotificationProductOfferingQualificationStateChangeNotificationParams{
		Context: ctx,
	}
}

// NewNotificationProductOfferingQualificationStateChangeNotificationParamsWithHTTPClient creates a new NotificationProductOfferingQualificationStateChangeNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationProductOfferingQualificationStateChangeNotificationParamsWithHTTPClient(client *http.Client) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	return &NotificationProductOfferingQualificationStateChangeNotificationParams{
		HTTPClient: client,
	}
}

/* NotificationProductOfferingQualificationStateChangeNotificationParams contains all the parameters to send to the API endpoint
   for the notification product offering qualification state change notification operation.

   Typically these are written to a http.Request.
*/
type NotificationProductOfferingQualificationStateChangeNotificationParams struct {

	// ProductOfferingQualificationStateChangeNotification.
	ProductOfferingQualificationStateChangeNotification *models.PoQEventContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notification product offering qualification state change notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) WithDefaults() *NotificationProductOfferingQualificationStateChangeNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notification product offering qualification state change notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) WithTimeout(timeout time.Duration) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) WithContext(ctx context.Context) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) WithHTTPClient(client *http.Client) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductOfferingQualificationStateChangeNotification adds the productOfferingQualificationStateChangeNotification to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) WithProductOfferingQualificationStateChangeNotification(productOfferingQualificationStateChangeNotification *models.PoQEventContainer) *NotificationProductOfferingQualificationStateChangeNotificationParams {
	o.SetProductOfferingQualificationStateChangeNotification(productOfferingQualificationStateChangeNotification)
	return o
}

// SetProductOfferingQualificationStateChangeNotification adds the productOfferingQualificationStateChangeNotification to the notification product offering qualification state change notification params
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) SetProductOfferingQualificationStateChangeNotification(productOfferingQualificationStateChangeNotification *models.PoQEventContainer) {
	o.ProductOfferingQualificationStateChangeNotification = productOfferingQualificationStateChangeNotification
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationProductOfferingQualificationStateChangeNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ProductOfferingQualificationStateChangeNotification != nil {
		if err := r.SetBodyParam(o.ProductOfferingQualificationStateChangeNotification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
