// Code generated by go-swagger; DO NOT EDIT.

package hub

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
)

// NewQuoteManagementHubFindParams creates a new QuoteManagementHubFindParams object
// with the default values initialized.
func NewQuoteManagementHubFindParams() *QuoteManagementHubFindParams {
	var ()
	return &QuoteManagementHubFindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteManagementHubFindParamsWithTimeout creates a new QuoteManagementHubFindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQuoteManagementHubFindParamsWithTimeout(timeout time.Duration) *QuoteManagementHubFindParams {
	var ()
	return &QuoteManagementHubFindParams{

		timeout: timeout,
	}
}

// NewQuoteManagementHubFindParamsWithContext creates a new QuoteManagementHubFindParams object
// with the default values initialized, and the ability to set a context for a request
func NewQuoteManagementHubFindParamsWithContext(ctx context.Context) *QuoteManagementHubFindParams {
	var ()
	return &QuoteManagementHubFindParams{

		Context: ctx,
	}
}

// NewQuoteManagementHubFindParamsWithHTTPClient creates a new QuoteManagementHubFindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQuoteManagementHubFindParamsWithHTTPClient(client *http.Client) *QuoteManagementHubFindParams {
	var ()
	return &QuoteManagementHubFindParams{
		HTTPClient: client,
	}
}

/*QuoteManagementHubFindParams contains all the parameters to send to the API endpoint
for the quote management hub find operation typically these are written to a http.Request
*/
type QuoteManagementHubFindParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the quote management hub find params
func (o *QuoteManagementHubFindParams) WithTimeout(timeout time.Duration) *QuoteManagementHubFindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote management hub find params
func (o *QuoteManagementHubFindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote management hub find params
func (o *QuoteManagementHubFindParams) WithContext(ctx context.Context) *QuoteManagementHubFindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote management hub find params
func (o *QuoteManagementHubFindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote management hub find params
func (o *QuoteManagementHubFindParams) WithHTTPClient(client *http.Client) *QuoteManagementHubFindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote management hub find params
func (o *QuoteManagementHubFindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the quote management hub find params
func (o *QuoteManagementHubFindParams) WithFields(fields *string) *QuoteManagementHubFindParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the quote management hub find params
func (o *QuoteManagementHubFindParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteManagementHubFindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// header param fields
		if err := r.SetHeaderParam("fields", *o.Fields); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}