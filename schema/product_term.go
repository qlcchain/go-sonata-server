package schema

import "github.com/qlcchain/go-sonata-server/models"

type ProductTerm struct {
	ID *string `json:"id"`

	// Description of the commitment
	Description string `json:"description,omitempty"`

	// duration
	Duration *models.Quantity `json:"duration,omitempty" gorm:"embedded"`

	// Name of the commitment
	Name string `json:"name,omitempty"`

	// valid for
	ValidFor *models.TimePeriod `json:"validFor,omitempty" gorm:"embedded"`
}
