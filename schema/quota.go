package schema

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type Quote struct {

	// Indicates the base (class) type of the quote.
	AtBaseType string `json:"@baseType,omitempty"`

	// Link to the schema describing the REST resource.
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the (class) type of the quote.
	AtType string `json:"@type,omitempty"`

	// agreement
	Agreement []*models.AgreementRef `json:"agreement" gorm:"foreignkey:ID"`

	// Description of the quote
	Description string `json:"description,omitempty"`

	// Date when the quoted was Cancelled or Rejected or Accepted
	// Format: date-time
	EffectiveQuoteCompletionDate strfmt.DateTime `json:"effectiveQuoteCompletionDate,omitempty"`

	// This is the date wished by the requester to have the requested quote item(s) delivered
	// Format: date
	ExpectedFulfillmentStartDate strfmt.Date `json:"expectedFulfillmentStartDate,omitempty"`

	// This is the date filled by the seller to indicate expected quote completion date.
	// Format: date
	ExpectedQuoteCompletionDate strfmt.Date `json:"expectedQuoteCompletionDate,omitempty"`

	// ID given by the consumer and only understandable by him (to facilitate his searches afterwards)
	ExternalID string `json:"externalId,omitempty"`

	// Hyperlink to access the quote
	Href string `json:"href,omitempty"`

	// Unique (within the quoting domain) identifier for the quote, as attributed by the quoting system
	ID string `json:"id,omitempty"`

	// If this flag is set to Yes, Buyer requests to have instant quoting to be provided in operation POST response
	// Required: true
	InstantSyncQuoting bool `json:"instantSyncQuoting"`

	// note
	Note []*Note `json:"note" gorm:"foreignkey:ID"`

	// This value MAY be assigned by the Buyer/Seller to identify a project the quoting request is associated with.
	ProjectID string `json:"projectId,omitempty"`

	// Date when the quote was created
	// Format: date-time
	QuoteDate strfmt.DateTime `json:"quoteDate,omitempty"`

	// quote item
	// Required: true
	QuoteItem []*QuoteItem `json:"quoteItem" gorm:"foreignkey:ID"`

	// quote level
	// Required: true
	QuoteLevel models.QuoteLevel `json:"quoteLevel"`

	// related party
	// Required: true
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// This is the date wished by the requester to have the quote completed (meaning priced)
	// Required: true
	// Format: date-time
	RequestedQuoteCompletionDate *strfmt.DateTime `json:"requestedQuoteCompletionDate"`

	// state
	// Required: true
	State models.QuoteStateType `json:"state"`

	// valid for
	ValidFor *models.TimePeriod `json:"validFor,omitempty" gorm:"embedded"`
}

func (m *Quote) ToQuoteSummaryView() *models.QuoteSummaryView {
	var roles []*models.RelatedPartyRole
	//FIXME: How to convert this??? store role in db???
	for _, p := range m.RelatedParty {
		r := ""
		if len(p.Role) > 0 {
			r = p.Role[0]
		}
		roles = append(roles, &models.RelatedPartyRole{
			AtReferredType: p.AtReferredType,
			ID:             swag.StringValue(p.ID),
			RelatedParty:   p.To(),
			Role:           swag.String(r),
		})
	}
	var quoteItem []*models.QuoteItem
	_ = util.Convert(m.QuoteItem, &quoteItem)

	return &models.QuoteSummaryView{
		AtBaseType:                   m.AtBaseType,
		AtSchemaLocation:             m.AtSchemaLocation,
		AtType:                       m.AtType,
		Category:                     "",
		EffectiveQuoteCompletionDate: m.EffectiveQuoteCompletionDate,
		ExpectedFulfillmentStartDate: m.ExpectedFulfillmentStartDate,
		ExpectedQuoteCompletionDate:  m.ExpectedQuoteCompletionDate,
		ExternalID:                   m.ExternalID,
		Href:                         m.Href,
		ID:                           m.ID,
		QuoteDate:                    m.QuoteDate,
		QuoteItem:                    quoteItem,
		QuoteLevel:                   m.QuoteLevel,
		RequestedQuoteCompletionDate: m.RequestedQuoteCompletionDate,
		State:                        models.QuoteState(m.State),
		ValidFor:                     m.ValidFor,
		RelatedPartyRole:             roles,
	}
}

func (m *Quote) ToQuoteFind() *models.QuoteFind {
	return &models.QuoteFind{
		Description:                  m.Description,
		EffectiveQuoteCompletionDate: m.EffectiveQuoteCompletionDate,
		ExpectedQuoteCompletionDate:  m.ExpectedQuoteCompletionDate,
		ExternalID:                   m.ExternalID,
		Href:                         m.Href,
		ID:                           m.ID,
		ProjectID:                    m.ProjectID,
		QuoteDate:                    m.QuoteDate,
		QuoteLevel:                   m.QuoteLevel,
		RequestedQuoteCompletionDate: m.RequestedQuoteCompletionDate,
		State:                        m.State,
	}
}

func (m *Quote) ToQuote() *models.Quote {
	return &models.Quote{
		AtBaseType:                   m.AtBaseType,
		AtSchemaLocation:             m.AtSchemaLocation,
		AtType:                       m.AtType,
		Agreement:                    m.Agreement,
		Description:                  m.Description,
		EffectiveQuoteCompletionDate: m.EffectiveQuoteCompletionDate,
		ExpectedFulfillmentStartDate: m.ExpectedFulfillmentStartDate,
		ExpectedQuoteCompletionDate:  m.ExpectedQuoteCompletionDate,
		ExternalID:                   m.ExternalID,
		Href:                         m.Href,
		ID:                           m.ID,
		InstantSyncQuoting:           m.InstantSyncQuoting,
		Note:                         ToNotes(m.Note),
		ProjectID:                    m.ProjectID,
		QuoteDate:                    m.QuoteDate,
		QuoteItem:                    ToQuoteItem(m.QuoteItem),
		QuoteLevel:                   m.QuoteLevel,
		RelatedParty:                 ToRelatedParty(m.RelatedParty),
		RequestedQuoteCompletionDate: m.RequestedQuoteCompletionDate,
		State:                        m.State,
		ValidFor:                     m.ValidFor,
	}
}

type QuoteItem struct {

	// Link to the schema describing this REST resource
	AtSchemaLocation string `json:"@schemaLocation,omitempty"`

	// Indicates the base (class) type of the quote Item.
	AtType string `json:"@type,omitempty"`

	// action
	// Required: true
	Action models.ProductActionType `json:"action"`

	// Identifier of the quote item (generally it is a sequence number 01, 02, 03, ...).
	// Required: true
	ID *string `json:"id"`

	// note
	Note []*Note `json:"note"`

	// product
	Product *Product `json:"product,omitempty"`

	// product offering
	ProductOffering *models.ProductOfferingRef `json:"productOffering,omitempty"`

	// qualification
	Qualification []*models.ProductOfferingQualificationRef `json:"qualification"`

	// quote item price
	QuoteItemPrice []*models.QuotePrice `json:"quoteItemPrice"`

	// quote item relationship
	QuoteItemRelationship []*models.QuoteItemRelationship `json:"quoteItemRelationship"`

	// quote item term
	QuoteItemTerm *models.ItemTerm `json:"quoteItemTerm,omitempty"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty"`

	// requested quote item term
	RequestedQuoteItemTerm *models.ItemTerm `json:"requestedQuoteItemTerm,omitempty"`

	// state
	// Required: true
	State models.QuoteItemStateType `json:"state"`
}

func (m *QuoteItem) To() *models.QuoteItem {
	return &models.QuoteItem{
		AtSchemaLocation:       m.AtSchemaLocation,
		AtType:                 m.AtType,
		Action:                 m.Action,
		ID:                     m.ID,
		Note:                   ToNotes(m.Note),
		Product:                m.Product.To(),
		ProductOffering:        m.ProductOffering,
		Qualification:          m.Qualification,
		QuoteItemPrice:         m.QuoteItemPrice,
		QuoteItemRelationship:  m.QuoteItemRelationship,
		QuoteItemTerm:          m.QuoteItemTerm,
		RelatedParty:           ToRelatedParty(m.RelatedParty),
		RequestedQuoteItemTerm: m.RequestedQuoteItemTerm,
		State:                  m.State,
	}
}

func ToQuoteItem(item []*QuoteItem) []*models.QuoteItem {
	var to []*models.QuoteItem
	for _, i := range item {
		to = append(to, i.To())
	}
	return to
}
