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

// NewQuoteManagementHubDeleteParams creates a new QuoteManagementHubDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuoteManagementHubDeleteParams() *QuoteManagementHubDeleteParams {
	return &QuoteManagementHubDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteManagementHubDeleteParamsWithTimeout creates a new QuoteManagementHubDeleteParams object
// with the ability to set a timeout on a request.
func NewQuoteManagementHubDeleteParamsWithTimeout(timeout time.Duration) *QuoteManagementHubDeleteParams {
	return &QuoteManagementHubDeleteParams{
		timeout: timeout,
	}
}

// NewQuoteManagementHubDeleteParamsWithContext creates a new QuoteManagementHubDeleteParams object
// with the ability to set a context for a request.
func NewQuoteManagementHubDeleteParamsWithContext(ctx context.Context) *QuoteManagementHubDeleteParams {
	return &QuoteManagementHubDeleteParams{
		Context: ctx,
	}
}

// NewQuoteManagementHubDeleteParamsWithHTTPClient creates a new QuoteManagementHubDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuoteManagementHubDeleteParamsWithHTTPClient(client *http.Client) *QuoteManagementHubDeleteParams {
	return &QuoteManagementHubDeleteParams{
		HTTPClient: client,
	}
}

/* QuoteManagementHubDeleteParams contains all the parameters to send to the API endpoint
   for the quote management hub delete operation.

   Typically these are written to a http.Request.
*/
type QuoteManagementHubDeleteParams struct {

	// HubID.
	HubID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quote management hub delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteManagementHubDeleteParams) WithDefaults() *QuoteManagementHubDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quote management hub delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteManagementHubDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) WithTimeout(timeout time.Duration) *QuoteManagementHubDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) WithContext(ctx context.Context) *QuoteManagementHubDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) WithHTTPClient(client *http.Client) *QuoteManagementHubDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHubID adds the hubID to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) WithHubID(hubID string) *QuoteManagementHubDeleteParams {
	o.SetHubID(hubID)
	return o
}

// SetHubID adds the hubId to the quote management hub delete params
func (o *QuoteManagementHubDeleteParams) SetHubID(hubID string) {
	o.HubID = hubID
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteManagementHubDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param HubId
	if err := r.SetPathParam("HubId", o.HubID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
