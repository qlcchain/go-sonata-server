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

// NewNotificationQuoteAttributeValueChangeNotificationParams creates a new NotificationQuoteAttributeValueChangeNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationQuoteAttributeValueChangeNotificationParams() *NotificationQuoteAttributeValueChangeNotificationParams {
	return &NotificationQuoteAttributeValueChangeNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationQuoteAttributeValueChangeNotificationParamsWithTimeout creates a new NotificationQuoteAttributeValueChangeNotificationParams object
// with the ability to set a timeout on a request.
func NewNotificationQuoteAttributeValueChangeNotificationParamsWithTimeout(timeout time.Duration) *NotificationQuoteAttributeValueChangeNotificationParams {
	return &NotificationQuoteAttributeValueChangeNotificationParams{
		timeout: timeout,
	}
}

// NewNotificationQuoteAttributeValueChangeNotificationParamsWithContext creates a new NotificationQuoteAttributeValueChangeNotificationParams object
// with the ability to set a context for a request.
func NewNotificationQuoteAttributeValueChangeNotificationParamsWithContext(ctx context.Context) *NotificationQuoteAttributeValueChangeNotificationParams {
	return &NotificationQuoteAttributeValueChangeNotificationParams{
		Context: ctx,
	}
}

// NewNotificationQuoteAttributeValueChangeNotificationParamsWithHTTPClient creates a new NotificationQuoteAttributeValueChangeNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationQuoteAttributeValueChangeNotificationParamsWithHTTPClient(client *http.Client) *NotificationQuoteAttributeValueChangeNotificationParams {
	return &NotificationQuoteAttributeValueChangeNotificationParams{
		HTTPClient: client,
	}
}

/* NotificationQuoteAttributeValueChangeNotificationParams contains all the parameters to send to the API endpoint
   for the notification quote attribute value change notification operation.

   Typically these are written to a http.Request.
*/
type NotificationQuoteAttributeValueChangeNotificationParams struct {

	// QuoteAttributeValueChangeNotification.
	QuoteAttributeValueChangeNotification []*models.QuoteEventPlus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notification quote attribute value change notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationQuoteAttributeValueChangeNotificationParams) WithDefaults() *NotificationQuoteAttributeValueChangeNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notification quote attribute value change notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationQuoteAttributeValueChangeNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) WithTimeout(timeout time.Duration) *NotificationQuoteAttributeValueChangeNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) WithContext(ctx context.Context) *NotificationQuoteAttributeValueChangeNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) WithHTTPClient(client *http.Client) *NotificationQuoteAttributeValueChangeNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteAttributeValueChangeNotification adds the quoteAttributeValueChangeNotification to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) WithQuoteAttributeValueChangeNotification(quoteAttributeValueChangeNotification []*models.QuoteEventPlus) *NotificationQuoteAttributeValueChangeNotificationParams {
	o.SetQuoteAttributeValueChangeNotification(quoteAttributeValueChangeNotification)
	return o
}

// SetQuoteAttributeValueChangeNotification adds the quoteAttributeValueChangeNotification to the notification quote attribute value change notification params
func (o *NotificationQuoteAttributeValueChangeNotificationParams) SetQuoteAttributeValueChangeNotification(quoteAttributeValueChangeNotification []*models.QuoteEventPlus) {
	o.QuoteAttributeValueChangeNotification = quoteAttributeValueChangeNotification
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationQuoteAttributeValueChangeNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.QuoteAttributeValueChangeNotification != nil {
		if err := r.SetBodyParam(o.QuoteAttributeValueChangeNotification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
