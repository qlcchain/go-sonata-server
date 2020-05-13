package schema

type SubUnit struct {
	ID string `json:"id,omitempty"`
	// The discriminator used for the subunit, often just a simple number but may also be a range.
	// Required: true
	SubUnitIdentifier *string `json:"subUnitIdentifier"`

	// The type of subunit e.g.BERTH, FLAT, PIER, SUITE, SHOP, TOWER, UNIT, WHARF.
	// Required: true
	SubUnitType *string `json:"subUnitType"`
}
