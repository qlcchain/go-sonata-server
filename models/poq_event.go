// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PoqEvent The product offering qualification passed as the "event" in a notification.
//
// swagger:model PoqEvent
type PoqEvent struct {

	// Link to the POQ
	Href string `json:"href,omitempty"`

	// The Serviceability Request's unique identifier.
	ID string `json:"id,omitempty"`
}

// Validate validates this poq event
func (m *PoqEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this poq event based on context it is used
func (m *PoqEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PoqEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoqEvent) UnmarshalBinary(b []byte) error {
	var res PoqEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
