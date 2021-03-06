// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GeographicLocation A set of coordinates (typically including latitude and longitude) that describes a particular location on earth.
//
// swagger:model GeographicLocation
type GeographicLocation struct {

	// geographic point
	// Required: true
	GeographicPoint []*GeographicPoint `json:"geographicPoint"`

	// Unique Identifier of a GeographicLocation
	ID string `json:"id,omitempty"`

	// The spatial reference system used to determine the coordinates
	// Required: true
	SpatialRef *string `json:"spatialRef"`
}

// Validate validates this geographic location
func (m *GeographicLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeographicPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpatialRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicLocation) validateGeographicPoint(formats strfmt.Registry) error {

	if err := validate.Required("geographicPoint", "body", m.GeographicPoint); err != nil {
		return err
	}

	for i := 0; i < len(m.GeographicPoint); i++ {
		if swag.IsZero(m.GeographicPoint[i]) { // not required
			continue
		}

		if m.GeographicPoint[i] != nil {
			if err := m.GeographicPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("geographicPoint" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GeographicLocation) validateSpatialRef(formats strfmt.Registry) error {

	if err := validate.Required("spatialRef", "body", m.SpatialRef); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this geographic location based on the context it is used
func (m *GeographicLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeographicPoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicLocation) contextValidateGeographicPoint(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GeographicPoint); i++ {

		if m.GeographicPoint[i] != nil {
			if err := m.GeographicPoint[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("geographicPoint" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeographicLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeographicLocation) UnmarshalBinary(b []byte) error {
	var res GeographicLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
