package schema

import "github.com/qlcchain/go-sonata-server/models"

type ProductPrice struct {
	ID *string `json:"id"`

	// A technical attribute to extend the class
	AtType string `json:"@type,omitempty"`

	// A narrative that explains in detail the semantics of this product price
	Description string `json:"description,omitempty"`

	// A short descriptive name such as "Subscription price"
	// Required: true
	Name *string `json:"name"`

	// price
	Price *models.Price `json:"price,omitempty" gorm:"embedded"`

	// price type
	PriceType *models.PriceType `json:"priceType,omitempty"`

	// recurring charge period
	RecurringChargePeriod models.ChargePeriod `json:"recurringChargePeriod,omitempty"`

	// Unit of Measure, if price depends on it (like Gb for example)
	UnitOfMeasure string `json:"unitOfMeasure,omitempty"`
}
