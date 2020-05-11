// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// QuoteEvent Event class is used to describe information structure used for notification.
//
// swagger:discriminator QuoteEvent eventId
type QuoteEvent interface {
	runtime.Validatable

	Event() QuoteSummaryView
	SetEvent(QuoteSummaryView)

	// Id of the event
	// Required: true
	EventID() string
	SetEventID(string)

	// Datetime when the event occurred
	// Required: true
	// Format: date-time
	EventTime() *strfmt.DateTime
	SetEventTime(*strfmt.DateTime)

	// event type
	// Required: true
	EventType() QuoteEventType
	SetEventType(QuoteEventType)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type quoteEvent struct {
	eventField QuoteSummaryView

	eventIdField string

	eventTimeField *strfmt.DateTime

	eventTypeField QuoteEventType
}

// Event gets the event of this polymorphic type
func (m *quoteEvent) Event() QuoteSummaryView {
	return m.eventField
}

// SetEvent sets the event of this polymorphic type
func (m *quoteEvent) SetEvent(val QuoteSummaryView) {
	m.eventField = val
}

// EventID gets the event Id of this polymorphic type
func (m *quoteEvent) EventID() string {
	return "QuoteEvent"
}

// SetEventID sets the event Id of this polymorphic type
func (m *quoteEvent) SetEventID(val string) {
}

// EventTime gets the event time of this polymorphic type
func (m *quoteEvent) EventTime() *strfmt.DateTime {
	return m.eventTimeField
}

// SetEventTime sets the event time of this polymorphic type
func (m *quoteEvent) SetEventTime(val *strfmt.DateTime) {
	m.eventTimeField = val
}

// EventType gets the event type of this polymorphic type
func (m *quoteEvent) EventType() QuoteEventType {
	return m.eventTypeField
}

// SetEventType sets the event type of this polymorphic type
func (m *quoteEvent) SetEventType(val QuoteEventType) {
	m.eventTypeField = val
}

// UnmarshalQuoteEventSlice unmarshals polymorphic slices of QuoteEvent
func UnmarshalQuoteEventSlice(reader io.Reader, consumer runtime.Consumer) ([]QuoteEvent, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []QuoteEvent
	for _, element := range elements {
		obj, err := unmarshalQuoteEvent(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalQuoteEvent unmarshals polymorphic QuoteEvent
func UnmarshalQuoteEvent(reader io.Reader, consumer runtime.Consumer) (QuoteEvent, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalQuoteEvent(data, consumer)
}

func unmarshalQuoteEvent(data []byte, consumer runtime.Consumer) (QuoteEvent, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the eventId property.
	var getType struct {
		EventID string `json:"eventId"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("eventId", "body", getType.EventID); err != nil {
		return nil, err
	}

	// The value of eventId is used to determine which type to create and unmarshal the data into
	switch getType.EventID {
	case "QuoteEvent":
		var result quoteEvent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "QuoteEventPlus":
		var result QuoteEventPlus
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid eventId value: %q", getType.EventID)
}

// Validate validates this quote event
func (m *quoteEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
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

func (m *quoteEvent) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event()); err != nil {
		return err
	}

	if err := m.Event().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("event")
		}
		return err
	}

	return nil
}

func (m *quoteEvent) validateEventTime(formats strfmt.Registry) error {

	if err := validate.Required("eventTime", "body", m.EventTime()); err != nil {
		return err
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *quoteEvent) validateEventType(formats strfmt.Registry) error {

	if err := m.EventType().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eventType")
		}
		return err
	}

	return nil
}
