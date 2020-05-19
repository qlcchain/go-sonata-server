package schema

type QualificationRef struct {

	// Technical attribute to extend the API
	AtReferredType string `json:"@referredType,omitempty"`

	// Hyperlink to the qualification previously done for this item
	Href string `json:"href,omitempty"`

	// id of the qualification previously done for this item
	ID *string `json:"id,omitempty"`

	// item id of the qualification previously done for this item
	QualificationItem string `json:"qualificationItem,omitempty"`
}
