package schema

import (
	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type QuotePrice struct {
	ID string `json:"id,omitempty"`
	// Indicates the base (class) type of the quote price
	AtType string `json:"@type,omitempty"`

	// Description of the quote/quote item price.
	Description string `json:"description,omitempty"`

	// Name of the quote /quote item price
	// Required: true
	Name *string `json:"name"`

	// price
	// Required: true
	Price *models.Price `json:"price" gorm:"embedded"`

	// price type
	// Required: true
	PriceType models.PriceType `json:"priceType"`

	// recurring charge period
	RecurringChargePeriod models.ChargePeriod `json:"recurringChargePeriod,omitempty"`
}

func (m *QuotePrice) To() *models.QuotePrice {
	if m == nil {
		return nil
	}
	to := &models.QuotePrice{}
	_ = util.Convert(m, to)
	return to
}

func ToQuotePrice(from []*QuotePrice) []*models.QuotePrice {
	if from == nil {
		return nil
	}
	var to []*models.QuotePrice
	_ = util.Convert(from, to)
	return to
}
