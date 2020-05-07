package schema

import (
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

type Product struct {

	// the Base Type of the product if specialization.
	// This is a technical attribute to extend this class.
	AtBaseType string `json:"@baseType,omitempty"`

	// A pointer to a file describing extension attributes (if used).
	// This is a technical attribute to extend this class.
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// The type of product if specialization
	AtType string `json:"@type,omitempty"`

	// agreement
	Agreement []*models.Agreement `json:"agreement" gorm:"foreignkey:ID"`

	// billing account
	BillingAccount []*models.BillingAccountRef `json:"billingAccount" gorm:"foreignkey:ID"`

	// This identifier is optionally provided during the Product ordering and stored for informative purpose in the Seller inventory.
	BuyerProductID string `json:"buyerProductId,omitempty"`

	// Reference of the product (link)
	Href string `json:"href,omitempty"`

	// Unique identifier of the product in the product domain.
	// Required: true
	ID *string `json:"id"`

	// Latest date when the product has been updated
	// Format: date-time
	LastUpdateDate strfmt.DateTime `json:"lastUpdateDate,omitempty"`

	// product offering
	ProductOffering *models.ProductOfferingRef `json:"productOffering,omitempty"`

	// product order
	ProductOrder []*models.ProductOrderRef `json:"productOrder" gorm:"foreignkey:ID"`

	// product price
	ProductPrice []*ProductPrice `json:"productPrice" gorm:"foreignkey:ID"`

	// product relationship
	ProductRelationship []*ProductRelationship `json:"productRelationship" gorm:"foreignkey:ID"`

	// product specification
	ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty" gorm:"foreignkey:ID"`

	// product term
	ProductTerm []*ProductTerm `json:"productTerm" gorm:"foreignkey:ID"`

	// related party
	RelatedParty []*models.RelatedParty `json:"relatedParty" gorm:"foreignkey:ID"`

	// site
	Site []*GeographicSite `json:"site" gorm:"foreignkey:ID"`

	// Start date is when the product is active for the first time (when the install in the product order has been processed).
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate"`

	// status
	// Required: true
	Status models.ProductStatus `json:"status"`

	// status change
	StatusChange []*StatusChange `json:"statusChange" gorm:"foreignkey:ID"`

	// Termination date (commercial) is when the product has been terminated (when the disconnect in the product order has been processed).
	// Format: date-time
	TerminationDate strfmt.DateTime `json:"terminationDate,omitempty"`
}
