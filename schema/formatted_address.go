package schema

import (
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type FormattedAddress struct {

	// The first address line in a formatted address
	// Required: true
	AddrLine1 *string `json:"addrLine1"`

	// The second address line in a formatted address
	AddrLine2 string `json:"addrLine2,omitempty"`

	// City that the address is in
	City string `json:"city,omitempty"`

	// Country that the address is in
	Country string `json:"country,omitempty"`

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
}

func (m *FormattedAddress) To() *models.FormattedAddress {
	if m == nil {
		return nil
	}
	to := &models.FormattedAddress{}
	if err := util.Convert(m, to); err != nil {
		log.Error(err)
	}
	return to
}

func (m *FormattedAddress) From(f *models.FormattedAddress) *FormattedAddress {
	if f == nil {
		return nil
	}
	r := &FormattedAddress{}
	if err := util.Convert(f, r); err != nil {
		log.Error(err)
	}
	return r
}
