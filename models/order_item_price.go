// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderItemPrice Structure used to define a product price.
// An order item could have 0-* order item price.
//
// swagger:model OrderItemPrice
type OrderItemPrice struct {

	// Technical attribute to extend the class
	AtType string `json:"@type,omitempty"`

	// A narrative that explains in detail the semantics of yhis order item price
	Description string `json:"description,omitempty"`

	// Name of the product price
	// Required: true
	Name *string `json:"name"`

	// price
	// Required: true
	Price *Price `json:"price"`

	// price type
	// Required: true
	PriceType PriceType `json:"priceType"`

	// recurring charge period
	RecurringChargePeriod ChargePeriod `json:"recurringChargePeriod,omitempty"`
}

// Validate validates this order item price
func (m *OrderItemPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringChargePeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItemPrice) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OrderItemPrice) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if m.Price != nil {
		if err := m.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemPrice) validatePriceType(formats strfmt.Registry) error {

	if err := m.PriceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priceType")
		}
		return err
	}

	return nil
}

func (m *OrderItemPrice) validateRecurringChargePeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.RecurringChargePeriod) { // not required
		return nil
	}

	if err := m.RecurringChargePeriod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("recurringChargePeriod")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderItemPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderItemPrice) UnmarshalBinary(b []byte) error {
	var res OrderItemPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
