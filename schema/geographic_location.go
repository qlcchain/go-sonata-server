package schema

type GeographicLocation struct {

	// geographic point
	// Required: true
	GeographicPoint []*GeographicPoint `json:"geographicPoint" gorm:"foreignkey:ID"`

	// Unique Identifier of a GeographicLocation
	ID string `json:"id,omitempty"`

	// The spatial reference system used to determine the coordinates
	// Required: true
	SpatialRef *string `json:"spatialRef"`
}
