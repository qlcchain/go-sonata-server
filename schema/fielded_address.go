package schema

type FieldedAddress struct {

	// City that the address is in
	City string `json:"city,omitempty"`

	// Country that the address is in
	Country string `json:"country,omitempty"`

	// geographic sub address
	GeographicSubAddress []*GeographicSubAddress `json:"geographicSubAddress" gorm:"foreignkey:ID"`

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
