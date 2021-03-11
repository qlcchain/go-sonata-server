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

// ProductOrderEventType Product Order event type
//
// swagger:model ProductOrderEventType
type ProductOrderEventType string

func NewProductOrderEventType(value ProductOrderEventType) *ProductOrderEventType {
	v := value
	return &v
}

const (

	// ProductOrderEventTypeProductOrderCreationNotification captures enum value "productOrderCreationNotification"
	ProductOrderEventTypeProductOrderCreationNotification ProductOrderEventType = "productOrderCreationNotification"

	// ProductOrderEventTypeProductOrderAttributeValueChangeNotification captures enum value "productOrderAttributeValueChangeNotification"
	ProductOrderEventTypeProductOrderAttributeValueChangeNotification ProductOrderEventType = "productOrderAttributeValueChangeNotification"

	// ProductOrderEventTypeProductOrderStateChangeNotification captures enum value "productOrderStateChangeNotification"
	ProductOrderEventTypeProductOrderStateChangeNotification ProductOrderEventType = "productOrderStateChangeNotification"

	// ProductOrderEventTypeProductOrderInformationRequiredNotification captures enum value "productOrderInformationRequiredNotification"
	ProductOrderEventTypeProductOrderInformationRequiredNotification ProductOrderEventType = "productOrderInformationRequiredNotification"
)

// for schema
var productOrderEventTypeEnum []interface{}

func init() {
	var res []ProductOrderEventType
	if err := json.Unmarshal([]byte(`["productOrderCreationNotification","productOrderAttributeValueChangeNotification","productOrderStateChangeNotification","productOrderInformationRequiredNotification"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productOrderEventTypeEnum = append(productOrderEventTypeEnum, v)
	}
}

func (m ProductOrderEventType) validateProductOrderEventTypeEnum(path, location string, value ProductOrderEventType) error {
	if err := validate.EnumCase(path, location, value, productOrderEventTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this product order event type
func (m ProductOrderEventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProductOrderEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this product order event type based on context it is used
func (m ProductOrderEventType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
