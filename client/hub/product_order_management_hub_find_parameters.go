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

// NewProductOrderManagementHubFindParams creates a new ProductOrderManagementHubFindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductOrderManagementHubFindParams() *ProductOrderManagementHubFindParams {
	return &ProductOrderManagementHubFindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductOrderManagementHubFindParamsWithTimeout creates a new ProductOrderManagementHubFindParams object
// with the ability to set a timeout on a request.
func NewProductOrderManagementHubFindParamsWithTimeout(timeout time.Duration) *ProductOrderManagementHubFindParams {
	return &ProductOrderManagementHubFindParams{
		timeout: timeout,
	}
}

// NewProductOrderManagementHubFindParamsWithContext creates a new ProductOrderManagementHubFindParams object
// with the ability to set a context for a request.
func NewProductOrderManagementHubFindParamsWithContext(ctx context.Context) *ProductOrderManagementHubFindParams {
	return &ProductOrderManagementHubFindParams{
		Context: ctx,
	}
}

// NewProductOrderManagementHubFindParamsWithHTTPClient creates a new ProductOrderManagementHubFindParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductOrderManagementHubFindParamsWithHTTPClient(client *http.Client) *ProductOrderManagementHubFindParams {
	return &ProductOrderManagementHubFindParams{
		HTTPClient: client,
	}
}

/* ProductOrderManagementHubFindParams contains all the parameters to send to the API endpoint
   for the product order management hub find operation.

   Typically these are written to a http.Request.
*/
type ProductOrderManagementHubFindParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the product order management hub find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductOrderManagementHubFindParams) WithDefaults() *ProductOrderManagementHubFindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the product order management hub find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductOrderManagementHubFindParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the product order management hub find params
func (o *ProductOrderManagementHubFindParams) WithTimeout(timeout time.Duration) *ProductOrderManagementHubFindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product order management hub find params
func (o *ProductOrderManagementHubFindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product order management hub find params
func (o *ProductOrderManagementHubFindParams) WithContext(ctx context.Context) *ProductOrderManagementHubFindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product order management hub find params
func (o *ProductOrderManagementHubFindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product order management hub find params
func (o *ProductOrderManagementHubFindParams) WithHTTPClient(client *http.Client) *ProductOrderManagementHubFindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product order management hub find params
func (o *ProductOrderManagementHubFindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ProductOrderManagementHubFindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
