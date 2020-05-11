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

// NewNotificationProductOfferingQualificationCreationNotificationParams creates a new NotificationProductOfferingQualificationCreationNotificationParams object
// with the default values initialized.
func NewNotificationProductOfferingQualificationCreationNotificationParams() *NotificationProductOfferingQualificationCreationNotificationParams {
	var ()
	return &NotificationProductOfferingQualificationCreationNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationProductOfferingQualificationCreationNotificationParamsWithTimeout creates a new NotificationProductOfferingQualificationCreationNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotificationProductOfferingQualificationCreationNotificationParamsWithTimeout(timeout time.Duration) *NotificationProductOfferingQualificationCreationNotificationParams {
	var ()
	return &NotificationProductOfferingQualificationCreationNotificationParams{

		timeout: timeout,
	}
}

// NewNotificationProductOfferingQualificationCreationNotificationParamsWithContext creates a new NotificationProductOfferingQualificationCreationNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotificationProductOfferingQualificationCreationNotificationParamsWithContext(ctx context.Context) *NotificationProductOfferingQualificationCreationNotificationParams {
	var ()
	return &NotificationProductOfferingQualificationCreationNotificationParams{

		Context: ctx,
	}
}

// NewNotificationProductOfferingQualificationCreationNotificationParamsWithHTTPClient creates a new NotificationProductOfferingQualificationCreationNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotificationProductOfferingQualificationCreationNotificationParamsWithHTTPClient(client *http.Client) *NotificationProductOfferingQualificationCreationNotificationParams {
	var ()
	return &NotificationProductOfferingQualificationCreationNotificationParams{
		HTTPClient: client,
	}
}

/*NotificationProductOfferingQualificationCreationNotificationParams contains all the parameters to send to the API endpoint
for the notification product offering qualification creation notification operation typically these are written to a http.Request
*/
type NotificationProductOfferingQualificationCreationNotificationParams struct {

	/*ProductOfferingQualificationCreationNotification*/
	ProductOfferingQualificationCreationNotification *models.PoQEventContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) WithTimeout(timeout time.Duration) *NotificationProductOfferingQualificationCreationNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) WithContext(ctx context.Context) *NotificationProductOfferingQualificationCreationNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) WithHTTPClient(client *http.Client) *NotificationProductOfferingQualificationCreationNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductOfferingQualificationCreationNotification adds the productOfferingQualificationCreationNotification to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) WithProductOfferingQualificationCreationNotification(productOfferingQualificationCreationNotification *models.PoQEventContainer) *NotificationProductOfferingQualificationCreationNotificationParams {
	o.SetProductOfferingQualificationCreationNotification(productOfferingQualificationCreationNotification)
	return o
}

// SetProductOfferingQualificationCreationNotification adds the productOfferingQualificationCreationNotification to the notification product offering qualification creation notification params
func (o *NotificationProductOfferingQualificationCreationNotificationParams) SetProductOfferingQualificationCreationNotification(productOfferingQualificationCreationNotification *models.PoQEventContainer) {
	o.ProductOfferingQualificationCreationNotification = productOfferingQualificationCreationNotification
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationProductOfferingQualificationCreationNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProductOfferingQualificationCreationNotification != nil {
		if err := r.SetBodyParam(o.ProductOfferingQualificationCreationNotification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
