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

// NewNotificationQuoteInformationRequiredNotificationParams creates a new NotificationQuoteInformationRequiredNotificationParams object
// with the default values initialized.
func NewNotificationQuoteInformationRequiredNotificationParams() *NotificationQuoteInformationRequiredNotificationParams {
	var ()
	return &NotificationQuoteInformationRequiredNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationQuoteInformationRequiredNotificationParamsWithTimeout creates a new NotificationQuoteInformationRequiredNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationQuoteInformationRequiredNotificationParamsWithTimeout(timeout time.Duration) *NotificationQuoteInformationRequiredNotificationParams {
	var ()
	return &NotificationQuoteInformationRequiredNotificationParams{

		timeout: timeout,
	}
}

// NewNotificationQuoteInformationRequiredNotificationParamsWithContext creates a new NotificationQuoteInformationRequiredNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationQuoteInformationRequiredNotificationParamsWithContext(ctx context.Context) *NotificationQuoteInformationRequiredNotificationParams {
	var ()
	return &NotificationQuoteInformationRequiredNotificationParams{

		Context: ctx,
	}
}

// NewNotificationQuoteInformationRequiredNotificationParamsWithHTTPClient creates a new NotificationQuoteInformationRequiredNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationQuoteInformationRequiredNotificationParamsWithHTTPClient(client *http.Client) *NotificationQuoteInformationRequiredNotificationParams {
	var ()
	return &NotificationQuoteInformationRequiredNotificationParams{
		HTTPClient: client,
	}
}

/*NotificationQuoteInformationRequiredNotificationParams contains all the parameters to send to the API endpoint
for the notification quote information required notification operation typically these are written to a http.Request
*/
type NotificationQuoteInformationRequiredNotificationParams struct {

	/*QuoteInformationRequiredNotification*/
	QuoteInformationRequiredNotification []*models.QuoteEventPlus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) WithTimeout(timeout time.Duration) *NotificationQuoteInformationRequiredNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) WithContext(ctx context.Context) *NotificationQuoteInformationRequiredNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) WithHTTPClient(client *http.Client) *NotificationQuoteInformationRequiredNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuoteInformationRequiredNotification adds the quoteInformationRequiredNotification to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) WithQuoteInformationRequiredNotification(quoteInformationRequiredNotification []*models.QuoteEventPlus) *NotificationQuoteInformationRequiredNotificationParams {
	o.SetQuoteInformationRequiredNotification(quoteInformationRequiredNotification)
	return o
}

// SetQuoteInformationRequiredNotification adds the quoteInformationRequiredNotification to the notification quote information required notification params
func (o *NotificationQuoteInformationRequiredNotificationParams) SetQuoteInformationRequiredNotification(quoteInformationRequiredNotification []*models.QuoteEventPlus) {
	o.QuoteInformationRequiredNotification = quoteInformationRequiredNotification
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationQuoteInformationRequiredNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.QuoteInformationRequiredNotification != nil {
		if err := r.SetBodyParam(o.QuoteInformationRequiredNotification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
