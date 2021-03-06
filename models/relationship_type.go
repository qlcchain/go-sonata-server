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

// RelationshipType Relationship type to be used between POQ item or product
//
// swagger:model RelationshipType
type RelationshipType string

func NewRelationshipType(value RelationshipType) *RelationshipType {
	v := value
	return &v
}

const (

	// RelationshipTypeReliesOn captures enum value "reliesOn"
	RelationshipTypeReliesOn RelationshipType = "reliesOn"

	// RelationshipTypeBundled captures enum value "bundled"
	RelationshipTypeBundled RelationshipType = "bundled"

	// RelationshipTypeComesFrom captures enum value "comesFrom"
	RelationshipTypeComesFrom RelationshipType = "comesFrom"
)

// for schema
var relationshipTypeEnum []interface{}

func init() {
	var res []RelationshipType
	if err := json.Unmarshal([]byte(`["reliesOn","bundled","comesFrom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relationshipTypeEnum = append(relationshipTypeEnum, v)
	}
}

func (m RelationshipType) validateRelationshipTypeEnum(path, location string, value RelationshipType) error {
	if err := validate.EnumCase(path, location, value, relationshipTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this relationship type
func (m RelationshipType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRelationshipTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this relationship type based on context it is used
func (m RelationshipType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
