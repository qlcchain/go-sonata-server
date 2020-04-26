// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewProductGetParams creates a new ProductGetParams object
// no default values defined in spec.
func NewProductGetParams() ProductGetParams {

	return ProductGetParams{}
}

// ProductGetParams contains all the bound params for the product get operation
// typically these are obtained from a http.Request
//
// swagger:parameters productGet
type ProductGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ProductID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewProductGetParams() beforehand.
func (o *ProductGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rProductID, rhkProductID, _ := route.Params.GetOK("ProductId")
	if err := o.bindProductID(rProductID, rhkProductID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProductID binds and validates parameter ProductID from path.
func (o *ProductGetParams) bindProductID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProductID = raw

	return nil
}
