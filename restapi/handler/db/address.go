package db

type FieldedAddressModel struct {
	// City that the address is in
	City string `json:"city,omitempty"`

	// Country that the address is in
	Country string `json:"country,omitempty"`

	// geographic sub adress
	//GeographicSubAdress []*GeographicSubAddress `json:"geographicSubAdress"`

	// Unique identifier of the address
	ID string `json:"id,omitempty"`

	// "An area of defined or undefined boundaries within a local authority or other legislatively defined area, usually rural or semi-rural in nature." [ANZLIC-STREET], or a suburb "a bounded locality within a city, town or shire principally of urban character " [ANZLICSTREET].
	Locality string `json:"locality,omitempty"`

	// The four-digit extension on an American postal code, what comes after the hyphen when specified.
	PostCodeExtension string `json:"postCodeExtension,omitempty"`

	// Descriptor for a postal delivery area, used to speed and simplify the delivery of mail (also known as zipcode)
	Postcode string `json:"postcode,omitempty"`

	// The State or Province that the address is in
	StateOrProvince string `json:"stateOrProvince,omitempty"`

	// Name of the street or other street type
	StreetName string `json:"streetName,omitempty"`

	// Number identifying a specific property on a public street. It may be combined with streetNrLast for ranged addresses
	StreetNr string `json:"streetNr,omitempty"`

	// Last number in a range of street numbers allocated to a property
	StreetNrLast string `json:"streetNrLast,omitempty"`

	// Last street number suffix for a ranged address
	StreetNrLastSuffix string `json:"streetNrLastSuffix,omitempty"`

	// The first street number suffix
	StreetNrSuffix string `json:"streetNrSuffix,omitempty"`

	// A modifier denoting a relative direction
	StreetSuffix string `json:"streetSuffix,omitempty"`

	// Alley, avenue, boulevard, brae, crescent, drive, highway, lane, terrace, parade, place, tarn, way, wharf
	StreetType string `json:"streetType,omitempty"`
}

type GeographicAddressModel struct {
	// Technical attribute to extend this class
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Technical attribute to extend this class
	AtType string `json:"@type,omitempty"`

	// This attribute specifies if a Buyer must use one of the known existing Service Sites at this location for any Products delivered to this Address.  For example, if a particular building owner mandated that all interconnects be done in a shared Public Meet-Me-Room, this attribute would be set to False for that Address.
	AllowsNewSite bool `json:"allowsNewSite,omitempty"`

	//FieldedAddressID string
	// fielded address
	FieldedAddress *FieldedAddressModel `json:"fieldedAddress,omitempty" gorm:"foreignkey:ID"`

	// geographic location
	//GeographicLocationID string
	//GeographicLocation   *models.GeographicLocation `json:"geographicLocation,omitempty" gorm:"foreignkey:ID"`

	// This attribute specifies if that Address contains Service Sites that are public such as Meet-Me-Rooms at an interconnect location or a shared telecom room in the basement of a multi-tenant building.
	HasPublicSite bool `json:"hasPublicSite,omitempty"`

	// Unique identifier of the address
	ID string `json:"id,omitempty"`

	// referenced address
	//ReferencedAddress *models.ReferencedAddress `json:"referencedAddress,omitempty"`
}
