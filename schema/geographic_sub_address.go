package schema

import "github.com/qlcchain/go-sonata-server/models"

type GeographicSubAddress struct {

	// Technical attribute to extend this class
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Technical attribute to extend this class
	AtType string `json:"@type,omitempty"`

	// Allows for buildings that have well-known names
	BuildingName string `json:"buildingName,omitempty"`

	// Unique Identifier of the subAddress
	ID string `json:"id,omitempty"`

	// Used where a level type may be repeated e.g. BASEMENT 1, BASEMENT 2
	LevelNumber string `json:"levelNumber,omitempty"`

	// Describes level types within a building
	LevelType string `json:"levelType,omitempty"`

	// Private streets internal to a property (e.g. a university) may have internal names that are not recorded by the land title office
	PrivateStreetName string `json:"privateStreetName,omitempty"`

	// Private streets numbers internal to a private street
	PrivateStreetNumber string `json:"privateStreetNumber,omitempty"`

	// sub unit
	SubUnit []*models.SubUnit `json:"subUnit" gorm:"foreignkey:SubUnitIdentifier"`
}
