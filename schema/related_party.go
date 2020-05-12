package schema

import (
	"strings"

	"github.com/qlcchain/go-sonata-server/models"
)

type RelatedParty struct {

	// Allow to specify the party type like Organization or Individual
	AtReferredType string `json:"@referredType,omitempty"`

	// The email address of the related party.
	EmailAddress string `json:"emailAddress,omitempty"`

	// An identifier of the related party.
	// Required: true
	ID *string `json:"id"`

	// The name of the related party, e.g. "Jean Pontus".
	// Required: true
	Name *string `json:"name"`

	// The telephone number of the related party.
	Number string `json:"number,omitempty"`

	// Phone number Extension.
	NumberExtension string `json:"numberExtension,omitempty"`

	// Role played by this party for this Site as UNISiteContact for example.
	// Required: true
	Role  []string `json:"role" gorm:"-"`
	Roles string   `json:"-" gorm:"column:roles"`
}

func (r *RelatedParty) To() *models.RelatedParty {
	var role []string
	if r.Roles != "" {
		role = strings.Split(r.Roles, ",")
	}
	return &models.RelatedParty{
		AtReferredType:  r.AtReferredType,
		EmailAddress:    r.EmailAddress,
		ID:              r.ID,
		Name:            r.Name,
		Number:          r.Number,
		NumberExtension: r.NumberExtension,
		Role:            role,
	}
}

func (r *RelatedParty) From(m *models.RelatedParty) *RelatedParty {
	r.AtReferredType = m.AtReferredType
	r.EmailAddress = m.EmailAddress
	r.ID = m.ID
	r.Name = m.Name
	r.Number = m.Number
	r.NumberExtension = m.NumberExtension
	r.Role = m.Role
	r.Roles = strings.Join(m.Role, ",")
	return r
}
