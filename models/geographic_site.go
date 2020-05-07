// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GeographicSite Designated as Service Site in TS.
// A fixed physical location at which a Product can be installed.  Its location can be described either with geocodes (Lat/Long information) or by association with an Address or Address Reference.  This association may include a Sub-address describing where within that Address or Address Reference, this particular Service Site is located.
//
// swagger:model GeographicSite
type GeographicSite struct {

	// Technical attribute to extend this class
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Technical attribute to extend this class
	AtType string `json:"@type,omitempty"`

	// Additional site information
	AdditionnalSiteInformation string `json:"additionnalSiteInformation,omitempty"`

	// A textual description of the Service Site.
	Description string `json:"description,omitempty"`

	// fielded address
	FieldedAddress *FieldedAddress `json:"fieldedAddress,omitempty" gorm:"foreignkey:ID"`

	// formatted adress
	FormattedAdress *FormattedAddress `json:"formattedAdress,omitempty" gorm:"foreignkey:ID"`

	// geographic location
	GeographicLocation *GeographicLocation `json:"geographicLocation,omitempty" gorm:"foreignkey:ID"`

	// Identifier of the Service Site unique within the Seller.
	ID string `json:"id,omitempty"`

	// referenced address
	ReferencedAddress *ReferencedAddress `json:"referencedAddress,omitempty" gorm:"foreignkey:ID"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty" gorm:"foreignkey:ID"`

	// The name of the company that is the administrative authority (e.g. controls access) for this Service Site. (For example, the building owner)
	SiteCompanyName string `json:"siteCompanyName,omitempty"`

	// The name of the company that is the administrative authority for the space within this Service Site. (For example, the company leasing space in a multi-tenant building).
	SiteCustomerName string `json:"siteCustomerName,omitempty"`

	// A name commonly used by people to refer to this Service Site.
	SiteName string `json:"siteName,omitempty"`

	// This defines whether a Service Site is public or private.  “PUBLIC” means that the existence of this Service Site is public information.  A meet-me-room in a hosted data center facility (where all interconnects between parties take place) is an example of a public Service Site.  A shared facility in the basement of a multi-tenant business building where all interconnects between parties take place is another example of a public Service Site.  “PRIVATE” means that the existence of this Service Site is on a need-to-know basis.  A wiring closet set up inside a customer facility just to connect two parties is an example of a private Service Site. For “PRIVATE” sites, the Seller does not return any information regarding the existence of this Service Site unless it has been established that this Buyer is authorized to obtain this information.
	SiteType string `json:"siteType,omitempty"`

	// status
	Status Status `json:"status,omitempty"`
}

// Validate validates this geographic site
func (m *GeographicSite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldedAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormattedAdress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeographicLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferencedAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeographicSite) validateFieldedAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldedAddress) { // not required
		return nil
	}

	if m.FieldedAddress != nil {
		if err := m.FieldedAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldedAddress")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicSite) validateFormattedAdress(formats strfmt.Registry) error {

	if swag.IsZero(m.FormattedAdress) { // not required
		return nil
	}

	if m.FormattedAdress != nil {
		if err := m.FormattedAdress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formattedAdress")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicSite) validateGeographicLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.GeographicLocation) { // not required
		return nil
	}

	if m.GeographicLocation != nil {
		if err := m.GeographicLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geographicLocation")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicSite) validateReferencedAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferencedAddress) { // not required
		return nil
	}

	if m.ReferencedAddress != nil {
		if err := m.ReferencedAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("referencedAddress")
			}
			return err
		}
	}

	return nil
}

func (m *GeographicSite) validateRelatedParty(formats strfmt.Registry) error {

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

func (m *GeographicSite) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeographicSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeographicSite) UnmarshalBinary(b []byte) error {
	var res GeographicSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
