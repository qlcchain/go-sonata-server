package schema

import (
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type ProductPrice struct {
	ID *string `json:"id"`

	// A technical attribute to extend the class
	AtType string `json:"@type,omitempty" gorm:"column:type"`

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

func ToProductPrice(from []*ProductPrice) []*models.ProductPrice {
	var to []*models.ProductPrice
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	return to
}

func FromProductPrice(from []*models.ProductPrice) []*ProductPrice {
	var to []*ProductPrice
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	for _, price := range to {
		price.ID = util.NewOrDefaultPtr(price.ID)
	}
	return to
}
