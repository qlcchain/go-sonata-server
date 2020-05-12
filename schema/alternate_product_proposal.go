package schema

import "github.com/qlcchain/go-sonata-server/models"

type AlternateProductProposal struct {

	// Technical attribute to extend the class.
	AtType string `json:"@type,omitempty"`

	// eligible product offering
	EligibleProductOffering []*models.ProductOfferingRef `json:"eligibleProductOffering" gorm:"foreignkey:ID"`

	// Identifier of the Product Offering Qualification alternate proposal
	// Required: true
	ID *string `json:"id"`

	// installation interval
	InstallationInterval *models.TimeInterval `json:"installationInterval,omitempty" gorm:"embedded"`

	// product specification
	ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty" gorm:"foreignkey:ID"`
}
