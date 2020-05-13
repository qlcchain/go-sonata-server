package schema

import (
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

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

func ToGeographicPoint(from []*GeographicPoint) []*models.GeographicPoint {
	var to []*models.GeographicPoint
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	return to
}

func FromGeographicPoint(from []*models.GeographicPoint) []*GeographicPoint {
	var to []*GeographicPoint
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	for _, point := range to {
		point.ID = util.NewOrDefault(point.ID)
	}
	return to
}
