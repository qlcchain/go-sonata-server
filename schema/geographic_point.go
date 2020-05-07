package schema

type GeographicPoint struct {

	// A unique identifier for the geographic point.
	ID string `json:"id,omitempty"`

	// The latitude expressed in decimal degrees format
	// Required: true
	Latitude *string `json:"latitude"`

	// The longitude expressed in decimal degrees format
	// Required: true
	Longitude *string `json:"longitude"`
}
