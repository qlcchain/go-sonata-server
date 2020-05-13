package schema

import (
	"strings"

	"github.com/qlcchain/go-sonata-server/util"

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
	if r == nil {
		return nil
	}

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
	if m == nil {
		return nil
	}
	return &RelatedParty{
		AtReferredType:  m.AtReferredType,
		EmailAddress:    m.EmailAddress,
		ID:              util.NewOrDefaultPtr(m.ID),
		Name:            m.Name,
		Number:          m.Number,
		NumberExtension: m.NumberExtension,
		Role:            m.Role,
		Roles:           strings.Join(m.Role, ","),
	}
}

func ToRelatedParty(party []*RelatedParty) []*models.RelatedParty {
	var r []*models.RelatedParty
	for _, p := range party {
		r = append(r, p.To())
	}
	return r
}

func FromRelatedParty(party []*models.RelatedParty) []*RelatedParty {
	var to []*RelatedParty
	for _, p := range party {
		r := &RelatedParty{
			ID: util.NewOrDefaultPtr(p.ID),
		}
		to = append(to, r.From(p))
	}
	return to
}
