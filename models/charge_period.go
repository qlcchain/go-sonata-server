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

// ChargePeriod charge period
//
// swagger:model ChargePeriod
type ChargePeriod string

func NewChargePeriod(value ChargePeriod) *ChargePeriod {
	v := value
	return &v
}

const (

	// ChargePeriodDay captures enum value "day"
	ChargePeriodDay ChargePeriod = "day"

	// ChargePeriodWeek captures enum value "week"
	ChargePeriodWeek ChargePeriod = "week"

	// ChargePeriodMonth captures enum value "month"
	ChargePeriodMonth ChargePeriod = "month"

	// ChargePeriodYear captures enum value "year"
	ChargePeriodYear ChargePeriod = "year"
)

// for schema
var chargePeriodEnum []interface{}

func init() {
	var res []ChargePeriod
	if err := json.Unmarshal([]byte(`["day","week","month","year"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chargePeriodEnum = append(chargePeriodEnum, v)
	}
}

func (m ChargePeriod) validateChargePeriodEnum(path, location string, value ChargePeriod) error {
	if err := validate.EnumCase(path, location, value, chargePeriodEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this charge period
func (m ChargePeriod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateChargePeriodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this charge period based on context it is used
func (m ChargePeriod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
