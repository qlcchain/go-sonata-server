package schema

type ProductOrderRef struct {

	// Hyperlink to the productOrder
	Href string `json:"href,omitempty"`

	// Identifier of the productOrder(provided by the seller)
	// Required: true
	ID *string `json:"id"`

	// Identifier of the line item (generally it is a sequence number 01, 02, 03, ...)
	// Required: true
	OrderItemID *string `json:"orderItemId"`
}
