// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QuoteItemCreate This structure is used to describe Quote item only in the POST operation for the request.
//
// swagger:model QuoteItem_Create
type QuoteItemCreate struct {

	// Link to the schema describing this REST resource
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the base (class) type of the quote Item.
	AtType string `json:"@type,omitempty"`

	// action
	// Required: true
	Action *ProductActionType `json:"action"`

	// Identifier of the quote item (generally it is a sequence number 01, 02, 03, ...).
	// Required: true
	ID *string `json:"id"`

	// note
	Note []*Note `json:"note"`

	// product
	Product *Product `json:"product,omitempty"`

	// product offering
	ProductOffering *ProductOfferingRef `json:"productOffering,omitempty"`

	// qualification
	Qualification *ProductOfferingQualificationRef `json:"qualification,omitempty"`

	// quote item relationship
	QuoteItemRelationship []*QuoteItemRelationship `json:"quoteItemRelationship"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// requested quote item term
	RequestedQuoteItemTerm *ItemTerm `json:"requestedQuoteItemTerm,omitempty"`
}

// Validate validates this quote item create
func (m *QuoteItemCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOffering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQualification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteItemRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedQuoteItemTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteItemCreate) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *QuoteItemCreate) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	for i := 0; i < len(m.Note); i++ {
		if swag.IsZero(m.Note[i]) { // not required
			continue
		}

		if m.Note[i] != nil {
			if err := m.Note[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItemCreate) validateProduct(formats strfmt.Registry) error {
	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) validateProductOffering(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductOffering) { // not required
		return nil
	}

	if m.ProductOffering != nil {
		if err := m.ProductOffering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productOffering")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) validateQualification(formats strfmt.Registry) error {
	if swag.IsZero(m.Qualification) { // not required
		return nil
	}

	if m.Qualification != nil {
		if err := m.Qualification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qualification")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) validateQuoteItemRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.QuoteItemRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.QuoteItemRelationship); i++ {
		if swag.IsZero(m.QuoteItemRelationship[i]) { // not required
			continue
		}

		if m.QuoteItemRelationship[i] != nil {
			if err := m.QuoteItemRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quoteItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItemCreate) validateRelatedParty(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedParty) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedParty); i++ {
		if swag.IsZero(m.RelatedParty[i]) { // not required
			continue
		}

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItemCreate) validateRequestedQuoteItemTerm(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedQuoteItemTerm) { // not required
		return nil
	}

	if m.RequestedQuoteItemTerm != nil {
		if err := m.RequestedQuoteItemTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedQuoteItemTerm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quote item create based on the context it is used
func (m *QuoteItemCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProduct(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductOffering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQualification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuoteItemRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedQuoteItemTerm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteItemCreate) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {
		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) contextValidateNote(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Note); i++ {

		if m.Note[i] != nil {
			if err := m.Note[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItemCreate) contextValidateProduct(ctx context.Context, formats strfmt.Registry) error {

	if m.Product != nil {
		if err := m.Product.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) contextValidateProductOffering(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductOffering != nil {
		if err := m.ProductOffering.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productOffering")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) contextValidateQualification(ctx context.Context, formats strfmt.Registry) error {

	if m.Qualification != nil {
		if err := m.Qualification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qualification")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteItemCreate) contextValidateQuoteItemRelationship(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QuoteItemRelationship); i++ {

		if m.QuoteItemRelationship[i] != nil {
			if err := m.QuoteItemRelationship[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quoteItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItemCreate) contextValidateRelatedParty(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedParty); i++ {

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteItemCreate) contextValidateRequestedQuoteItemTerm(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestedQuoteItemTerm != nil {
		if err := m.RequestedQuoteItemTerm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedQuoteItemTerm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuoteItemCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteItemCreate) UnmarshalBinary(b []byte) error {
	var res QuoteItemCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
