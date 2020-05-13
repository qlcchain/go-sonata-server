package schema

import (
	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

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

func (m *GeographicLocation) To() *models.GeographicLocation {
	if m == nil {
		return nil
	}
	return &models.GeographicLocation{
		GeographicPoint: ToGeographicPoint(m.GeographicPoint),
		ID:              m.ID,
		SpatialRef:      m.SpatialRef,
	}
}

func (m *GeographicLocation) From(f *models.GeographicLocation) *GeographicLocation {
	if f == nil {
		return nil
	}

	return &GeographicLocation{
		GeographicPoint: FromGeographicPoint(f.GeographicPoint),
		ID:              util.NewOrDefault(f.ID),
		SpatialRef:      m.SpatialRef,
	}
}
