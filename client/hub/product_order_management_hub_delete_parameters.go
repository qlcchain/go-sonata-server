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

// NewProductOrderManagementHubDeleteParams creates a new ProductOrderManagementHubDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductOrderManagementHubDeleteParams() *ProductOrderManagementHubDeleteParams {
	return &ProductOrderManagementHubDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductOrderManagementHubDeleteParamsWithTimeout creates a new ProductOrderManagementHubDeleteParams object
// with the ability to set a timeout on a request.
func NewProductOrderManagementHubDeleteParamsWithTimeout(timeout time.Duration) *ProductOrderManagementHubDeleteParams {
	return &ProductOrderManagementHubDeleteParams{
		timeout: timeout,
	}
}

// NewProductOrderManagementHubDeleteParamsWithContext creates a new ProductOrderManagementHubDeleteParams object
// with the ability to set a context for a request.
func NewProductOrderManagementHubDeleteParamsWithContext(ctx context.Context) *ProductOrderManagementHubDeleteParams {
	return &ProductOrderManagementHubDeleteParams{
		Context: ctx,
	}
}

// NewProductOrderManagementHubDeleteParamsWithHTTPClient creates a new ProductOrderManagementHubDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductOrderManagementHubDeleteParamsWithHTTPClient(client *http.Client) *ProductOrderManagementHubDeleteParams {
	return &ProductOrderManagementHubDeleteParams{
		HTTPClient: client,
	}
}

/* ProductOrderManagementHubDeleteParams contains all the parameters to send to the API endpoint
   for the product order management hub delete operation.

   Typically these are written to a http.Request.
*/
type ProductOrderManagementHubDeleteParams struct {

	// HubID.
	HubID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the product order management hub delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductOrderManagementHubDeleteParams) WithDefaults() *ProductOrderManagementHubDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the product order management hub delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductOrderManagementHubDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) WithTimeout(timeout time.Duration) *ProductOrderManagementHubDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) WithContext(ctx context.Context) *ProductOrderManagementHubDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) WithHTTPClient(client *http.Client) *ProductOrderManagementHubDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHubID adds the hubID to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) WithHubID(hubID string) *ProductOrderManagementHubDeleteParams {
	o.SetHubID(hubID)
	return o
}

// SetHubID adds the hubId to the product order management hub delete params
func (o *ProductOrderManagementHubDeleteParams) SetHubID(hubID string) {
	o.HubID = hubID
}

// WriteToRequest writes these params to a swagger request
func (o *ProductOrderManagementHubDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
