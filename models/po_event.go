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

// PoEvent Event class is used to describe information structure used for notification.
//
// swagger:model PoEvent
type PoEvent struct {

	// event
	// Required: true
	Event *ProductOrderEvent `json:"event"`

	// event Id
	// Required: true
	EventID *string `json:"eventId"`

	// event time
	// Required: true
	// Format: date-time
	EventTime *strfmt.DateTime `json:"eventTime"`

	// event type
	// Required: true
	EventType *ProductOrderEventType `json:"eventType"`
}

// Validate validates this po event
func (m *PoEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoEvent) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *PoEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *PoEvent) validateEventTime(formats strfmt.Registry) error {

	if err := validate.Required("eventTime", "body", m.EventTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PoEvent) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("eventType", "body", m.EventType); err != nil {
		return err
	}

	if err := validate.Required("eventType", "body", m.EventType); err != nil {
		return err
	}

	if m.EventType != nil {
		if err := m.EventType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this po event based on the context it is used
func (m *PoEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoEvent) contextValidateEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.Event != nil {
		if err := m.Event.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *PoEvent) contextValidateEventType(ctx context.Context, formats strfmt.Registry) error {

	if m.EventType != nil {
		if err := m.EventType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoEvent) UnmarshalBinary(b []byte) error {
	var res PoEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
