package schema

import "github.com/qlcchain/go-sonata-server/models"

type ProductSpecificationRef struct {

	// describing
	Describing *models.Describing `json:"describing,omitempty" gorm:"embedded"`

	// A unique identifier of the product spec.
	// Required: true
	ID *string `json:"id"`
}
