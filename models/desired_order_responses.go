// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DesiredOrderResponses The kind of responses that the buyer expects to receive from the seller.
//
// swagger:model DesiredOrderResponses
type DesiredOrderResponses string

func NewDesiredOrderResponses(value DesiredOrderResponses) *DesiredOrderResponses {
	v := value
	return &v
}

const (

	// DesiredOrderResponsesConfirmationAndEngineeringDesign captures enum value "confirmationAndEngineeringDesign"
	DesiredOrderResponsesConfirmationAndEngineeringDesign DesiredOrderResponses = "confirmationAndEngineeringDesign"

	// DesiredOrderResponsesConfirmationOnly captures enum value "confirmationOnly"
	DesiredOrderResponsesConfirmationOnly DesiredOrderResponses = "confirmationOnly"

	// DesiredOrderResponsesNone captures enum value "none"
	DesiredOrderResponsesNone DesiredOrderResponses = "none"
)

// for schema
var desiredOrderResponsesEnum []interface{}

func init() {
	var res []DesiredOrderResponses
	if err := json.Unmarshal([]byte(`["confirmationAndEngineeringDesign","confirmationOnly","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		desiredOrderResponsesEnum = append(desiredOrderResponsesEnum, v)
	}
}

func (m DesiredOrderResponses) validateDesiredOrderResponsesEnum(path, location string, value DesiredOrderResponses) error {
	if err := validate.EnumCase(path, location, value, desiredOrderResponsesEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this desired order responses
func (m DesiredOrderResponses) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDesiredOrderResponsesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this desired order responses based on context it is used
func (m DesiredOrderResponses) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
