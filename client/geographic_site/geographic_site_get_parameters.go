// Code generated by go-swagger; DO NOT EDIT.

package geographic_site

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

// NewGeographicSiteGetParams creates a new GeographicSiteGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeographicSiteGetParams() *GeographicSiteGetParams {
	return &GeographicSiteGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeographicSiteGetParamsWithTimeout creates a new GeographicSiteGetParams object
// with the ability to set a timeout on a request.
func NewGeographicSiteGetParamsWithTimeout(timeout time.Duration) *GeographicSiteGetParams {
	return &GeographicSiteGetParams{
		timeout: timeout,
	}
}

// NewGeographicSiteGetParamsWithContext creates a new GeographicSiteGetParams object
// with the ability to set a context for a request.
func NewGeographicSiteGetParamsWithContext(ctx context.Context) *GeographicSiteGetParams {
	return &GeographicSiteGetParams{
		Context: ctx,
	}
}

// NewGeographicSiteGetParamsWithHTTPClient creates a new GeographicSiteGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeographicSiteGetParamsWithHTTPClient(client *http.Client) *GeographicSiteGetParams {
	return &GeographicSiteGetParams{
		HTTPClient: client,
	}
}

/* GeographicSiteGetParams contains all the parameters to send to the API endpoint
   for the geographic site get operation.

   Typically these are written to a http.Request.
*/
type GeographicSiteGetParams struct {

	// SiteID.
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the geographic site get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeographicSiteGetParams) WithDefaults() *GeographicSiteGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the geographic site get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeographicSiteGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the geographic site get params
func (o *GeographicSiteGetParams) WithTimeout(timeout time.Duration) *GeographicSiteGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the geographic site get params
func (o *GeographicSiteGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the geographic site get params
func (o *GeographicSiteGetParams) WithContext(ctx context.Context) *GeographicSiteGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the geographic site get params
func (o *GeographicSiteGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the geographic site get params
func (o *GeographicSiteGetParams) WithHTTPClient(client *http.Client) *GeographicSiteGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the geographic site get params
func (o *GeographicSiteGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the geographic site get params
func (o *GeographicSiteGetParams) WithSiteID(siteID string) *GeographicSiteGetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the geographic site get params
func (o *GeographicSiteGetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GeographicSiteGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SiteId
	if err := r.SetPathParam("SiteId", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
