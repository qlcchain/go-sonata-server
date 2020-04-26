// Code generated by go-swagger; DO NOT EDIT.

package geographic_address_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
)

// NewGeographicAddressValidationCreateParams creates a new GeographicAddressValidationCreateParams object
// no default values defined in spec.
func NewGeographicAddressValidationCreateParams() GeographicAddressValidationCreateParams {

	return GeographicAddressValidationCreateParams{}
}

// GeographicAddressValidationCreateParams contains all the bound params for the geographic address validation create operation
// typically these are obtained from a http.Request
//
// swagger:parameters geographicAddressValidationCreate
type GeographicAddressValidationCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	AddressValidationRequest *models.GeographicAddressValidationCreate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGeographicAddressValidationCreateParams() beforehand.
func (o *GeographicAddressValidationCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.GeographicAddressValidationCreate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("addressValidationRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("addressValidationRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.AddressValidationRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("addressValidationRequest", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
