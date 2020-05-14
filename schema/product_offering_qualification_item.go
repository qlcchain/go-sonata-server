package schema

import (
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

type ProductOfferingQualificationItem struct {

	// When sub-classing, this defines the sub-class entity name
	AtType string `json:"@type,omitempty"`

	// action
	// Required: true
	Action models.ProductActionType `json:"action"`

	// alternate product proposal
	AlternateProductProposal []*AlternateProductProposal `json:"alternateProductProposal" gorm:"foreignkey:ID"`

	// eligible product offering
	EligibleProductOffering []*models.ProductOfferingRef `json:"eligibleProductOffering" gorm:"foreignkey:ID"`

	// Date until seller is guaranted the qualification result.
	// Format: date-time
	GuaranteedUntilDate strfmt.DateTime `json:"guaranteedUntilDate,omitempty"`

	// Id of this POQ item
	// Required: true
	ID *string `json:"id"`

	// installation interval
	InstallationInterval *models.TimeInterval `json:"installationInterval,omitempty" gorm:"embedded"`

	// product
	Product *Product `json:"product,omitempty" gorm:"foreignkey:ID"`

	// product offering
	ProductOffering *models.ProductOfferingRef `json:"productOffering,omitempty" gorm:"foreignkey:ID"`

	// product offering qualification item relationship
	ProductOfferingQualificationItemRelationship []*models.ProductOfferingQualificationItemRelationship `json:"productOfferingQualificationItemRelationship" gorm:"foreignkey:ID"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty" gorm:"many2many:poq_item_related_party;foreignkey:ID"`

	// A description of the reason a particular color is being provided. This may include a specific standard reason codes and descriptions.
	ServiceConfidenceReason string `json:"serviceConfidenceReason,omitempty"`

	// serviceability confidence
	ServiceabilityConfidence models.ServiceabilityColor `json:"serviceabilityConfidence,omitempty"`

	// state
	// Required: true
	State models.ProductOfferingQualificationItemStateType `json:"state"`

	// state change
	StateChange []*StateChange `json:"stateChange" gorm:"foreignkey:ID"`

	// termination error
	TerminationError []*models.TerminationError `json:"terminationError" gorm:"foreignkey:ID"`
}
