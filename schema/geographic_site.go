package schema

import (
	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type GeographicSite struct {

	// Technical attribute to extend this class
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Technical attribute to extend this class
	AtType string `json:"@type,omitempty"`

	// Additional site information
	AdditionalSiteInformation string `json:"additionalSiteInformation,omitempty"`

	// A textual description of the Service Site.
	Description string `json:"description,omitempty"`

	// fielded address
	FieldedAddress *FieldedAddress `json:"fieldedAddress,omitempty" gorm:"foreignkey:ID"`

	// formatted address
	FormattedAddress *FormattedAddress `json:"formattedAddress,omitempty" gorm:"foreignkey:ID"`

	// geographic location
	GeographicLocation *GeographicLocation `json:"geographicLocation,omitempty" gorm:"foreignkey:ID"`

	// Identifier of the Service Site unique within the Seller.
	ID string `json:"id,omitempty"`

	// referenced address
	ReferencedAddress *models.ReferencedAddress `json:"referencedAddress,omitempty" gorm:"foreignkey:ID"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty" gorm:"foreignkey:ID"`

	// The name of the company that is the administrative authority (e.g. controls access) for this Service Site. (For example, the building owner)
	SiteCompanyName string `json:"siteCompanyName,omitempty"`

	// The name of the company that is the administrative authority for the space within this Service Site. (For example, the company leasing space in a multi-tenant building).
	SiteCustomerName string `json:"siteCustomerName,omitempty"`

	// A name commonly used by people to refer to this Service Site.
	SiteName string `json:"siteName,omitempty"`

	// This defines whether a Service Site is public or private.  “PUBLIC” means that the existence of this Service Site is public information.  A meet-me-room in a hosted data center facility (where all interconnects between parties take place) is an example of a public Service Site.  A shared facility in the basement of a multi-tenant business building where all interconnects between parties take place is another example of a public Service Site.  “PRIVATE” means that the existence of this Service Site is on a need-to-know basis.  A wiring closet set up inside a customer facility just to connect two parties is an example of a private Service Site. For “PRIVATE” sites, the Seller does not return any information regarding the existence of this Service Site unless it has been established that this Buyer is authorized to obtain this information.
	SiteType string `json:"siteType,omitempty"`

	// status
	Status models.Status `json:"status,omitempty"`
}

func (m *GeographicSite) To() *models.GeographicSite {
	if m == nil {
		return nil
	}
	return &models.GeographicSite{
		AtSchemaLocation:           m.AtSchemaLocation,
		AtType:                     m.AtType,
		AdditionnalSiteInformation: m.AdditionalSiteInformation,
		Description:                m.Description,
		FieldedAddress:             m.FieldedAddress.To(),
		FormattedAdress:            m.FormattedAddress.To(),
		GeographicLocation:         m.GeographicLocation.To(),
		ID:                         m.ID,
		ReferencedAddress:          m.ReferencedAddress,
		RelatedParty:               ToRelatedParty(m.RelatedParty),
		SiteCompanyName:            m.SiteCompanyName,
		SiteCustomerName:           m.SiteCustomerName,
		SiteName:                   m.SiteName,
		SiteType:                   m.SiteType,
		Status:                     m.Status,
	}
}

func (m *GeographicSite) From(site *models.GeographicSite) *GeographicSite {
	if site == nil {
		return nil
	}

	r := &GeographicSite{
		AtSchemaLocation:          site.AtSchemaLocation,
		AtType:                    site.AtType,
		AdditionalSiteInformation: site.AdditionnalSiteInformation,
		Description:               site.Description,
		ID:                        util.NewOrDefault(site.ID),
		ReferencedAddress:         site.ReferencedAddress,
		RelatedParty:              FromRelatedParty(site.RelatedParty),
		SiteCompanyName:           site.SiteCompanyName,
		SiteCustomerName:          site.SiteCustomerName,
		SiteName:                  site.SiteName,
		SiteType:                  site.SiteType,
		Status:                    site.Status,
	}
	r.GeographicLocation.From(site.GeographicLocation)
	r.FieldedAddress.From(site.FieldedAddress)
	r.FormattedAddress.From(site.FormattedAdress)
	return r
}

func ToGeographicSite(sites []*GeographicSite) []*models.GeographicSite {
	var to []*models.GeographicSite
	for _, s := range sites {
		to = append(to, s.To())
	}
	return to
}

func FromGeographicSite(sites []*models.GeographicSite) []*GeographicSite {
	if sites == nil {
		return nil
	}
	var to []*GeographicSite
	for _, s := range sites {
		r := &GeographicSite{}
		to = append(to, r.From(s))
	}
	return to
}
