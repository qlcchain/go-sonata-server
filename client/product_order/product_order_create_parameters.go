// Code generated by go-swagger; DO NOT EDIT.

package product_order

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

// NewProductOrderCreateParams creates a new ProductOrderCreateParams object
// with the default values initialized.
func NewProductOrderCreateParams() *ProductOrderCreateParams {
	var ()
	return &ProductOrderCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProductOrderCreateParamsWithTimeout creates a new ProductOrderCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProductOrderCreateParamsWithTimeout(timeout time.Duration) *ProductOrderCreateParams {
	var ()
	return &ProductOrderCreateParams{

		timeout: timeout,
	}
}

// NewProductOrderCreateParamsWithContext creates a new ProductOrderCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewProductOrderCreateParamsWithContext(ctx context.Context) *ProductOrderCreateParams {
	var ()
	return &ProductOrderCreateParams{

		Context: ctx,
	}
}

// NewProductOrderCreateParamsWithHTTPClient creates a new ProductOrderCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProductOrderCreateParamsWithHTTPClient(client *http.Client) *ProductOrderCreateParams {
	var ()
	return &ProductOrderCreateParams{
		HTTPClient: client,
	}
}

/*ProductOrderCreateParams contains all the parameters to send to the API endpoint
for the product order create operation typically these are written to a http.Request
*/
type ProductOrderCreateParams struct {

	/*ProductOrder*/
	ProductOrder *models.ProductOrderCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the product order create params
func (o *ProductOrderCreateParams) WithTimeout(timeout time.Duration) *ProductOrderCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product order create params
func (o *ProductOrderCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product order create params
func (o *ProductOrderCreateParams) WithContext(ctx context.Context) *ProductOrderCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product order create params
func (o *ProductOrderCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product order create params
func (o *ProductOrderCreateParams) WithHTTPClient(client *http.Client) *ProductOrderCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product order create params
func (o *ProductOrderCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductOrder adds the productOrder to the product order create params
func (o *ProductOrderCreateParams) WithProductOrder(productOrder *models.ProductOrderCreate) *ProductOrderCreateParams {
	o.SetProductOrder(productOrder)
	return o
}

// SetProductOrder adds the productOrder to the product order create params
func (o *ProductOrderCreateParams) SetProductOrder(productOrder *models.ProductOrderCreate) {
	o.ProductOrder = productOrder
}

// WriteToRequest writes these params to a swagger request
func (o *ProductOrderCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProductOrder != nil {
		if err := r.SetBodyParam(o.ProductOrder); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}