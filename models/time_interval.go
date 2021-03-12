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

// TimeInterval A time interval, e.g.  3 hours, or 5 days.
//
// swagger:model TimeInterval
type TimeInterval struct {

	// Amount
	// Required: true
	Amount *int32 `json:"amount"`

	// time unit
	// Required: true
	TimeUnit *TimeUnit `json:"timeUnit"`
}

// Validate validates this time interval
func (m *TimeInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeInterval) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TimeInterval) validateTimeUnit(formats strfmt.Registry) error {

	if err := validate.Required("timeUnit", "body", m.TimeUnit); err != nil {
		return err
	}

	if err := validate.Required("timeUnit", "body", m.TimeUnit); err != nil {
		return err
	}

	if m.TimeUnit != nil {
		if err := m.TimeUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeUnit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this time interval based on the context it is used
func (m *TimeInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimeUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeInterval) contextValidateTimeUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeUnit != nil {
		if err := m.TimeUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeUnit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeInterval) UnmarshalBinary(b []byte) error {
	var res TimeInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
