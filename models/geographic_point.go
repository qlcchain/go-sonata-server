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

// GeographicPoint A geometric point on earth, which can include a latitude, longitude, and elevation in “decimal degrees” format.
//
// swagger:model GeographicPoint
type GeographicPoint struct {

	// A unique identifier for the geographic point.
	ID string `json:"id,omitempty"`

	// The latitude expressed in decimal degrees format
	// Required: true
	Latitude *string `json:"latitude"`

	// The longitude expressed in decimal degrees format
	// Required: true
	Longitude *string `json:"longitude"`
}

// Validate validates this geographic point
func (m *GeographicPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicPoint) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *GeographicPoint) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this geographic point based on context it is used
func (m *GeographicPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeographicPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeographicPoint) UnmarshalBinary(b []byte) error {
	var res GeographicPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
