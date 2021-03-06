// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FormattedAddress A type of Address that has discrete fields for each type of boundary or identifier with the exception of street and more specific location details, which are combined into a maximum of two strings based on local postal addressing conventions
//
// swagger:model FormattedAddress
type FormattedAddress struct {

	// The first address line in a formatted address
	// Required: true
	AddrLine1 *string `json:"addrLine1"`

	// The second address line in a formatted address
	AddrLine2 string `json:"addrLine2,omitempty"`

	// City that the address is in
	City string `json:"city,omitempty"`

	// Country that the address is in
	Country string `json:"country,omitempty"`

	// Unique identifier of the address
	ID string `json:"id,omitempty"`

	// "An area of defined or undefined boundaries within a local authority or other legislatively defined area, usually rural or semi-rural in nature." [ANZLIC-STREET], or a suburb "a bounded locality within a city, town or shire principally of urban character " [ANZLICSTREET].
	Locality string `json:"locality,omitempty"`

	// The four-digit extension on an American postal code, what comes after the hyphen when specified.
	PostCodeExtension string `json:"postCodeExtension,omitempty"`

	// Descriptor for a postal delivery area, used to speed and simplify the delivery of mail (also known as zipcode)
	Postcode string `json:"postcode,omitempty"`

	// The State or Province that the address is in
	StateOrProvince string `json:"stateOrProvince,omitempty"`
}

// Validate validates this formatted address
func (m *FormattedAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddrLine1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormattedAddress) validateAddrLine1(formats strfmt.Registry) error {

	if err := validate.Required("addrLine1", "body", m.AddrLine1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this formatted address based on context it is used
func (m *FormattedAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FormattedAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormattedAddress) UnmarshalBinary(b []byte) error {
	var res FormattedAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
