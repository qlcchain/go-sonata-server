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

// QuoteEventPlus quote event plus
//
// swagger:model QuoteEventPlus
type QuoteEventPlus struct {
	QuoteEvent

	// field path
	// Required: true
	FieldPath []string `json:"fieldPath"`

	// resource path
	// Required: true
	ResourcePath *string `json:"resourcePath"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *QuoteEventPlus) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 QuoteEvent
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.QuoteEvent = aO0

	// AO1
	var dataAO1 struct {
		FieldPath []string `json:"fieldPath"`

		ResourcePath *string `json:"resourcePath"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.FieldPath = dataAO1.FieldPath

	m.ResourcePath = dataAO1.ResourcePath

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m QuoteEventPlus) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.QuoteEvent)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		FieldPath []string `json:"fieldPath"`

		ResourcePath *string `json:"resourcePath"`
	}

	dataAO1.FieldPath = m.FieldPath

	dataAO1.ResourcePath = m.ResourcePath

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this quote event plus
func (m *QuoteEventPlus) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with QuoteEvent
	if err := m.QuoteEvent.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteEventPlus) validateFieldPath(formats strfmt.Registry) error {

	if err := validate.Required("fieldPath", "body", m.FieldPath); err != nil {
		return err
	}

	return nil
}

func (m *QuoteEventPlus) validateResourcePath(formats strfmt.Registry) error {

	if err := validate.Required("resourcePath", "body", m.ResourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this quote event plus based on the context it is used
func (m *QuoteEventPlus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with QuoteEvent
	if err := m.QuoteEvent.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *QuoteEventPlus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteEventPlus) UnmarshalBinary(b []byte) error {
	var res QuoteEventPlus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
