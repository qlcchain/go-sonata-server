// Code generated by go-swagger; DO NOT EDIT.

package geographic_address_validation

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

// NewGeographicAddressValidationCreateParams creates a new GeographicAddressValidationCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeographicAddressValidationCreateParams() *GeographicAddressValidationCreateParams {
	return &GeographicAddressValidationCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeographicAddressValidationCreateParamsWithTimeout creates a new GeographicAddressValidationCreateParams object
// with the ability to set a timeout on a request.
func NewGeographicAddressValidationCreateParamsWithTimeout(timeout time.Duration) *GeographicAddressValidationCreateParams {
	return &GeographicAddressValidationCreateParams{
		timeout: timeout,
	}
}

// NewGeographicAddressValidationCreateParamsWithContext creates a new GeographicAddressValidationCreateParams object
// with the ability to set a context for a request.
func NewGeographicAddressValidationCreateParamsWithContext(ctx context.Context) *GeographicAddressValidationCreateParams {
	return &GeographicAddressValidationCreateParams{
		Context: ctx,
	}
}

// NewGeographicAddressValidationCreateParamsWithHTTPClient creates a new GeographicAddressValidationCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeographicAddressValidationCreateParamsWithHTTPClient(client *http.Client) *GeographicAddressValidationCreateParams {
	return &GeographicAddressValidationCreateParams{
		HTTPClient: client,
	}
}

/* GeographicAddressValidationCreateParams contains all the parameters to send to the API endpoint
   for the geographic address validation create operation.

   Typically these are written to a http.Request.
*/
type GeographicAddressValidationCreateParams struct {

	// AddressValidationRequest.
	AddressValidationRequest *models.GeographicAddressValidationCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the geographic address validation create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeographicAddressValidationCreateParams) WithDefaults() *GeographicAddressValidationCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the geographic address validation create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeographicAddressValidationCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) WithTimeout(timeout time.Duration) *GeographicAddressValidationCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) WithContext(ctx context.Context) *GeographicAddressValidationCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) WithHTTPClient(client *http.Client) *GeographicAddressValidationCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressValidationRequest adds the addressValidationRequest to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) WithAddressValidationRequest(addressValidationRequest *models.GeographicAddressValidationCreate) *GeographicAddressValidationCreateParams {
	o.SetAddressValidationRequest(addressValidationRequest)
	return o
}

// SetAddressValidationRequest adds the addressValidationRequest to the geographic address validation create params
func (o *GeographicAddressValidationCreateParams) SetAddressValidationRequest(addressValidationRequest *models.GeographicAddressValidationCreate) {
	o.AddressValidationRequest = addressValidationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GeographicAddressValidationCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AddressValidationRequest != nil {
		if err := r.SetBodyParam(o.AddressValidationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
