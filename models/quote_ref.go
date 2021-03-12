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

// QuoteRef Quote used before this order to define pricing terms
//
// swagger:model QuoteRef
type QuoteRef struct {

	// Technical attribute to extend the API
	AtReferredType string `json:"@referredType,omitempty"`

	// Hyperlink to access the quote (using MEF quote API)
	Href string `json:"href,omitempty"`

	// Unique identifier for the Quote that is generated by the Seller when the Quote is initially accepted via an API.
	// Required: true
	ID *string `json:"id"`

	// Identifier of the quote item
	QuoteItem string `json:"quoteItem,omitempty"`
}

// Validate validates this quote ref
func (m *QuoteRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quote ref based on context it is used
func (m *QuoteRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuoteRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteRef) UnmarshalBinary(b []byte) error {
	var res QuoteRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
