package schema

type ProductRef struct {

	// Targeted Buyer product id - Informative
	BuyerProductID string `json:"buyerProductId,omitempty"`

	// Reference of the product
	Href string `json:"href,omitempty"`

	// Unique identifier of the product
	// Required: true
	ID *string `json:"id"`
}
