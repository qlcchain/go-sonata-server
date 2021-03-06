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

// ProductOrderItemStateType An enumeration of valid order item statuses.
//
// swagger:model ProductOrderItemStateType
type ProductOrderItemStateType string

func NewProductOrderItemStateType(value ProductOrderItemStateType) *ProductOrderItemStateType {
	v := value
	return &v
}

const (

	// ProductOrderItemStateTypeAcknowledged captures enum value "acknowledged"
	ProductOrderItemStateTypeAcknowledged ProductOrderItemStateType = "acknowledged"

	// ProductOrderItemStateTypeRejected captures enum value "rejected"
	ProductOrderItemStateTypeRejected ProductOrderItemStateType = "rejected"

	// ProductOrderItemStateTypeInProgress captures enum value "inProgress"
	ProductOrderItemStateTypeInProgress ProductOrderItemStateType = "inProgress"

	// ProductOrderItemStateTypePending captures enum value "pending"
	ProductOrderItemStateTypePending ProductOrderItemStateType = "pending"

	// ProductOrderItemStateTypeHeld captures enum value "held"
	ProductOrderItemStateTypeHeld ProductOrderItemStateType = "held"

	// ProductOrderItemStateTypeAssesssingCancellation captures enum value "assesssingCancellation"
	ProductOrderItemStateTypeAssesssingCancellation ProductOrderItemStateType = "assesssingCancellation"

	// ProductOrderItemStateTypePendingCancellation captures enum value "pendingCancellation"
	ProductOrderItemStateTypePendingCancellation ProductOrderItemStateType = "pendingCancellation"

	// ProductOrderItemStateTypeCancelled captures enum value "cancelled"
	ProductOrderItemStateTypeCancelled ProductOrderItemStateType = "cancelled"

	// ProductOrderItemStateTypeInProgressDotConfigured captures enum value "inProgress.configured"
	ProductOrderItemStateTypeInProgressDotConfigured ProductOrderItemStateType = "inProgress.configured"

	// ProductOrderItemStateTypeFailed captures enum value "failed"
	ProductOrderItemStateTypeFailed ProductOrderItemStateType = "failed"

	// ProductOrderItemStateTypeCompleted captures enum value "completed"
	ProductOrderItemStateTypeCompleted ProductOrderItemStateType = "completed"
)

// for schema
var productOrderItemStateTypeEnum []interface{}

func init() {
	var res []ProductOrderItemStateType
	if err := json.Unmarshal([]byte(`["acknowledged","rejected","inProgress","pending","held","assesssingCancellation","pendingCancellation","cancelled","inProgress.configured","failed","completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productOrderItemStateTypeEnum = append(productOrderItemStateTypeEnum, v)
	}
}

func (m ProductOrderItemStateType) validateProductOrderItemStateTypeEnum(path, location string, value ProductOrderItemStateType) error {
	if err := validate.EnumCase(path, location, value, productOrderItemStateTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this product order item state type
func (m ProductOrderItemStateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProductOrderItemStateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this product order item state type based on context it is used
func (m ProductOrderItemStateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
