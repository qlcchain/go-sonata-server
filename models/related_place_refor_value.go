// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelatedPlaceReforValue Place defines the places where the products qualification must be done.
//
// swagger:discriminator RelatedPlaceReforValue @type
type RelatedPlaceReforValue interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The actual type of the target instance when needed for disambiguation.
	// Used when place is described by reference
	// @referredType could be valued to FormattedAddress, FieldedAddress, GeographicSite, GeographicLocation or ReferencedAddress
	AtReferredType() string
	SetAtReferredType(string)

	// When sub-classing, this defines the sub-class entity name.
	// Used when place is described by value (litterally)
	// Could be valued to FormattedAddress, FieldedAddress,  GeographicLocation or ReferencedAddress
	AtType() string
	SetAtType(string)

	// href to this place resource
	Href() string
	SetHref(string)

	// id of this place
	// Max Length: 45
	ID() string
	SetID(string)

	// Role of this place
	// Required: true
	Role() *string
	SetRole(*string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type relatedPlaceReforValue struct {
	atReferredTypeField string

	atTypeField string

	hrefField string

	idField string

	roleField *string
}

// AtReferredType gets the at referred type of this polymorphic type
func (m *relatedPlaceReforValue) AtReferredType() string {
	return m.atReferredTypeField
}

// SetAtReferredType sets the at referred type of this polymorphic type
func (m *relatedPlaceReforValue) SetAtReferredType(val string) {
	m.atReferredTypeField = val
}

// AtType gets the at type of this polymorphic type
func (m *relatedPlaceReforValue) AtType() string {
	return "RelatedPlaceReforValue"
}

// SetAtType sets the at type of this polymorphic type
func (m *relatedPlaceReforValue) SetAtType(val string) {
}

// Href gets the href of this polymorphic type
func (m *relatedPlaceReforValue) Href() string {
	return m.hrefField
}

// SetHref sets the href of this polymorphic type
func (m *relatedPlaceReforValue) SetHref(val string) {
	m.hrefField = val
}

// ID gets the id of this polymorphic type
func (m *relatedPlaceReforValue) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *relatedPlaceReforValue) SetID(val string) {
	m.idField = val
}

// Role gets the role of this polymorphic type
func (m *relatedPlaceReforValue) Role() *string {
	return m.roleField
}

// SetRole sets the role of this polymorphic type
func (m *relatedPlaceReforValue) SetRole(val *string) {
	m.roleField = val
}

// UnmarshalRelatedPlaceReforValueSlice unmarshals polymorphic slices of RelatedPlaceReforValue
func UnmarshalRelatedPlaceReforValueSlice(reader io.Reader, consumer runtime.Consumer) ([]RelatedPlaceReforValue, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []RelatedPlaceReforValue
	for _, element := range elements {
		obj, err := unmarshalRelatedPlaceReforValue(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalRelatedPlaceReforValue unmarshals polymorphic RelatedPlaceReforValue
func UnmarshalRelatedPlaceReforValue(reader io.Reader, consumer runtime.Consumer) (RelatedPlaceReforValue, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalRelatedPlaceReforValue(data, consumer)
}

func unmarshalRelatedPlaceReforValue(data []byte, consumer runtime.Consumer) (RelatedPlaceReforValue, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		AtType string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.AtType); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.AtType {
	case "RelatedPlaceReforValue":
		var result relatedPlaceReforValue
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)
}

// Validate validates this related place refor value
func (m *relatedPlaceReforValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *relatedPlaceReforValue) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID(), 45); err != nil {
		return err
	}

	return nil
}

func (m *relatedPlaceReforValue) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this related place refor value based on context it is used
func (m *relatedPlaceReforValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
