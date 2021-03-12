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

// ProductOfferingQualificationItem An individual article included in a POQ that describes a Product of a particular type (Product Offering) being delivered to a specific geographical location.  The objective is to determine if it is feasible for the Seller to deliver this item as described and for the Seller to inform the Buyer of the estimated time interval to complete this delivery.
//
// swagger:model ProductOfferingQualificationItem
type ProductOfferingQualificationItem struct {

	// When sub-classing, this defines the sub-class entity name
	AtType string `json:"@type,omitempty"`

	// action
	// Required: true
	Action *ProductActionType `json:"action"`

	// alternate product proposal
	AlternateProductProposal []*AlternateProductProposal `json:"alternateProductProposal"`

	// eligible product offering
	EligibleProductOffering []*ProductOfferingRef `json:"eligibleProductOffering"`

	// Date until seller is guaranted the qualification result.
	// Format: date-time
	GuaranteedUntilDate strfmt.DateTime `json:"guaranteedUntilDate,omitempty"`

	// Id of this POQ item
	// Required: true
	ID *string `json:"id"`

	// installation interval
	InstallationInterval *TimeInterval `json:"installationInterval,omitempty"`

	// product
	Product *Product `json:"product,omitempty"`

	// product offering
	ProductOffering *ProductOfferingRef `json:"productOffering,omitempty"`

	// product offering qualification item relationship
	ProductOfferingQualificationItemRelationship []*ProductOfferingQualificationItemRelationship `json:"productOfferingQualificationItemRelationship"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// A description of the reason a particular color is being provided. This may include a specific standard reason codes and descriptions.
	ServiceConfidenceReason string `json:"serviceConfidenceReason,omitempty"`

	// serviceability confidence
	ServiceabilityConfidence ServiceabilityColor `json:"serviceabilityConfidence,omitempty"`

	// state
	// Required: true
	State *ProductOfferingQualificationItemStateType `json:"state"`

	// state change
	StateChange []*StateChange `json:"stateChange"`

	// termination error
	TerminationError []*TerminationError `json:"terminationError"`
}

// Validate validates this product offering qualification item
func (m *ProductOfferingQualificationItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlternateProductProposal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEligibleProductOffering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuaranteedUntilDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallationInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOffering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOfferingQualificationItemRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceabilityConfidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductOfferingQualificationItem) validateAction(formats strfmt.Registry) error {

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

func (m *ProductOfferingQualificationItem) validateAlternateProductProposal(formats strfmt.Registry) error {
	if swag.IsZero(m.AlternateProductProposal) { // not required
		return nil
	}

	for i := 0; i < len(m.AlternateProductProposal); i++ {
		if swag.IsZero(m.AlternateProductProposal[i]) { // not required
			continue
		}

		if m.AlternateProductProposal[i] != nil {
			if err := m.AlternateProductProposal[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternateProductProposal" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateEligibleProductOffering(formats strfmt.Registry) error {
	if swag.IsZero(m.EligibleProductOffering) { // not required
		return nil
	}

	for i := 0; i < len(m.EligibleProductOffering); i++ {
		if swag.IsZero(m.EligibleProductOffering[i]) { // not required
			continue
		}

		if m.EligibleProductOffering[i] != nil {
			if err := m.EligibleProductOffering[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eligibleProductOffering" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateGuaranteedUntilDate(formats strfmt.Registry) error {
	if swag.IsZero(m.GuaranteedUntilDate) { // not required
		return nil
	}

	if err := validate.FormatOf("guaranteedUntilDate", "body", "date-time", m.GuaranteedUntilDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateInstallationInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallationInterval) { // not required
		return nil
	}

	if m.InstallationInterval != nil {
		if err := m.InstallationInterval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installationInterval")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateProduct(formats strfmt.Registry) error {
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

func (m *ProductOfferingQualificationItem) validateProductOffering(formats strfmt.Registry) error {
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

func (m *ProductOfferingQualificationItem) validateProductOfferingQualificationItemRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductOfferingQualificationItemRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductOfferingQualificationItemRelationship); i++ {
		if swag.IsZero(m.ProductOfferingQualificationItemRelationship[i]) { // not required
			continue
		}

		if m.ProductOfferingQualificationItemRelationship[i] != nil {
			if err := m.ProductOfferingQualificationItemRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productOfferingQualificationItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateRelatedParty(formats strfmt.Registry) error {
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

func (m *ProductOfferingQualificationItem) validateServiceabilityConfidence(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceabilityConfidence) { // not required
		return nil
	}

	if err := m.ServiceabilityConfidence.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceabilityConfidence")
		}
		return err
	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateStateChange(formats strfmt.Registry) error {
	if swag.IsZero(m.StateChange) { // not required
		return nil
	}

	for i := 0; i < len(m.StateChange); i++ {
		if swag.IsZero(m.StateChange[i]) { // not required
			continue
		}

		if m.StateChange[i] != nil {
			if err := m.StateChange[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stateChange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) validateTerminationError(formats strfmt.Registry) error {
	if swag.IsZero(m.TerminationError) { // not required
		return nil
	}

	for i := 0; i < len(m.TerminationError); i++ {
		if swag.IsZero(m.TerminationError[i]) { // not required
			continue
		}

		if m.TerminationError[i] != nil {
			if err := m.TerminationError[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("terminationError" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this product offering qualification item based on the context it is used
func (m *ProductOfferingQualificationItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAlternateProductProposal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEligibleProductOffering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallationInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProduct(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductOffering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductOfferingQualificationItemRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceabilityConfidence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateChange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerminationError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProductOfferingQualificationItem) contextValidateAlternateProductProposal(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AlternateProductProposal); i++ {

		if m.AlternateProductProposal[i] != nil {
			if err := m.AlternateProductProposal[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternateProductProposal" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateEligibleProductOffering(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EligibleProductOffering); i++ {

		if m.EligibleProductOffering[i] != nil {
			if err := m.EligibleProductOffering[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eligibleProductOffering" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateInstallationInterval(ctx context.Context, formats strfmt.Registry) error {

	if m.InstallationInterval != nil {
		if err := m.InstallationInterval.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installationInterval")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateProduct(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProductOfferingQualificationItem) contextValidateProductOffering(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProductOfferingQualificationItem) contextValidateProductOfferingQualificationItemRelationship(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProductOfferingQualificationItemRelationship); i++ {

		if m.ProductOfferingQualificationItemRelationship[i] != nil {
			if err := m.ProductOfferingQualificationItemRelationship[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productOfferingQualificationItemRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateRelatedParty(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProductOfferingQualificationItem) contextValidateServiceabilityConfidence(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServiceabilityConfidence.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceabilityConfidence")
		}
		return err
	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateStateChange(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StateChange); i++ {

		if m.StateChange[i] != nil {
			if err := m.StateChange[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stateChange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductOfferingQualificationItem) contextValidateTerminationError(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TerminationError); i++ {

		if m.TerminationError[i] != nil {
			if err := m.TerminationError[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("terminationError" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductOfferingQualificationItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductOfferingQualificationItem) UnmarshalBinary(b []byte) error {
	var res ProductOfferingQualificationItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
