package schema

import (
	"strconv"

	"github.com/rs/xid"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type Quote struct {

	// Indicates the base (class) type of the quote.
	AtBaseTypeField string `json:"@baseType,omitempty"`

	// Link to the schema describing the REST resource.
	// Technical attribute to extend this class.
	AtSchemaLocationField string `json:"@schemaLocation,omitempty"`

	// Indicates the (class) type of the quote.
	// Technical attribute to extend this class.
	Type string `json:"@type,omitempty"`

	// agreement
	AgreementField []*models.AgreementRef `json:"agreement" gorm:"foreignkey:ID"`

	// Description of the quote
	DescriptionField string `json:"description,omitempty"`

	// This is the date wished by the requester to have the requested quote item(s) delivered
	// Format: date
	ExpectedFulfillmentStartDateField strfmt.Date `json:"expectedFulfillmentStartDate,omitempty"`
	//This is the date filled by the seller to indicate expected quote completion date.
	ExpectedQuoteCompletionDateField strfmt.Date `json:"expectedQuoteCompletionDate,omitempty"`
	//Date when the quoted was Cancelled or Rejected or Accepted
	EffectiveQuoteCompletionDateField strfmt.DateTime `json:"effectiveQuoteCompletionDate,omitempty"`

	// ID given by the consumer and only understandable by him (to facilitate his searches afterwards)
	ExternalIDField string `json:"externalId,omitempty"`

	// If this flag is set to Yes, Buyer requests to have instant quoting to be provided in operation POST response
	InstantSyncQuotingField *bool `json:"instantSyncQuoting,omitempty"`

	// note
	NoteField []*Note `json:"note" gorm:"foreignkey:ID"`

	// This value MAY be assigned by the Buyer/Seller to identify a project the quoting request is associated with.
	ProjectIDField string `json:"projectID,omitempty"`

	// quote item
	// Required: true
	QuoteItemField []*QuoteItem `json:"quoteItem" gorm:"foreignkey:ID"`

	// quote level
	QuoteLevelField models.QuoteLevel `json:"quoteLevel,omitempty"`

	// related party
	// Required: true
	RelatedPartyField []*RelatedParty `json:"relatedParty" gorm:"foreignkey:ID"`

	// This is the date wished by the requester to have the quote completed (meaning priced).
	// This attribute is not considered when instantSyncQuoting is set to Yes.
	// Required: true
	// Format: date-time
	RequestedQuoteCompletionDateField *strfmt.DateTime `json:"requestedQuoteCompletionDate"`

	HrefField string `json:"href,omitempty"`
	IDField   string `json:"id,omitempty"`
	//Date when the quote was created
	QuoteDateField strfmt.DateTime       `json:"quoteDate,omitempty"`
	StateField     models.QuoteStateType `json:"state,omitempty"`
	ValidForField  *models.TimePeriod    `json:"validFor,omitempty"`
}

func (m *Quote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgreement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedFulfillmentStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstantSyncQuoting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuoteLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedQuoteCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quote) validateAgreement(formats strfmt.Registry) error {

	if swag.IsZero(m.Agreement()) { // not required
		return nil
	}

	for i := 0; i < len(m.Agreement()); i++ {
		if swag.IsZero(m.AgreementField[i]) { // not required
			continue
		}

		if m.AgreementField[i] != nil {
			if err := m.AgreementField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agreement" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Quote) validateEffectiveQuoteCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveQuoteCompletionDate()) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveQuoteCompletionDate", "body", "date-time", m.EffectiveQuoteCompletionDate().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateExpectedFulfillmentStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedFulfillmentStartDate()) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedFulfillmentStartDate", "body", "date", m.ExpectedFulfillmentStartDate().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateExpectedQuoteCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpectedQuoteCompletionDate()) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedQuoteCompletionDate", "body", "date", m.ExpectedQuoteCompletionDate().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateInstantSyncQuoting(formats strfmt.Registry) error {

	if err := validate.Required("instantSyncQuoting", "body", bool(m.InstantSyncQuoting())); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note()) { // not required
		return nil
	}

	for i := 0; i < len(m.Note()); i++ {
		if swag.IsZero(m.NoteField[i]) { // not required
			continue
		}

		if m.NoteField[i] != nil {
			to := &models.Note{}
			if err := util.Convert(m.NoteField[i], to); err == nil {
				if err := to.Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("note" + "." + strconv.Itoa(i))
					}
					return err
				}
			}
		}

	}

	return nil
}

func (m *Quote) validateQuoteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.QuoteDate()) { // not required
		return nil
	}

	if err := validate.FormatOf("quoteDate", "body", "date-time", m.QuoteDate().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateQuoteItem(formats strfmt.Registry) error {

	if err := validate.Required("quoteItem", "body", m.QuoteItem()); err != nil {
		return err
	}

	for i := 0; i < len(m.QuoteItem()); i++ {
		if swag.IsZero(m.QuoteItemField[i]) { // not required
			continue
		}

		if m.QuoteItemField[i] != nil {
			qi := &models.QuoteItem{}
			if err := util.Convert(m.QuoteItemField[i], qi); err == nil {
				if err := qi.Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("quoteItem" + "." + strconv.Itoa(i))
					}
					return err
				}
			} else {
				return err
			}
		}

	}

	return nil
}

func (m *Quote) validateQuoteLevel(formats strfmt.Registry) error {

	if err := m.QuoteLevel().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quoteLevel")
		}
		return err
	}

	return nil
}

func (m *Quote) validateRelatedParty(formats strfmt.Registry) error {

	if err := validate.Required("relatedParty", "body", m.RelatedParty()); err != nil {
		return err
	}

	for i := 0; i < len(m.RelatedParty()); i++ {
		if swag.IsZero(m.RelatedPartyField[i]) { // not required
			continue
		}

		if m.RelatedPartyField[i] != nil {
			to := &models.RelatedParty{}
			if err := util.Convert(m.RelatedPartyField[i], to); err == nil {
				if err := to.Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
					}
					return err
				}
			} else {
				return err
			}
		}

	}

	return nil
}

func (m *Quote) validateRequestedQuoteCompletionDate(formats strfmt.Registry) error {

	if err := validate.Required("requestedQuoteCompletionDate", "body", m.RequestedQuoteCompletionDate()); err != nil {
		return err
	}

	if err := validate.FormatOf("requestedQuoteCompletionDate", "body", "date-time", m.RequestedQuoteCompletionDate().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateState(formats strfmt.Registry) error {

	if err := m.State().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Quote) validateValidFor(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFor()) { // not required
		return nil
	}

	if m.ValidFor() != nil {
		if err := m.ValidFor().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validFor")
			}
			return err
		}
	}

	return nil
}

func (m *Quote) AtBaseType() string {
	return m.AtBaseTypeField
}

func (m *Quote) SetAtBaseType(s string) {
	m.AtBaseTypeField = s
}

func (m *Quote) AtSchemaLocation() string {
	return m.AtSchemaLocationField
}

func (m *Quote) SetAtSchemaLocation(s string) {
	m.AtSchemaLocationField = s
}

func (m *Quote) AtType() string {
	return m.Type
}

func (m *Quote) SetAtType(s string) {
	m.Type = s
}

func (m *Quote) Agreement() []*models.AgreementRef {
	return m.AgreementField
}

func (m *Quote) SetAgreement(refs []*models.AgreementRef) {
	m.AgreementField = refs
}

func (m *Quote) Description() string {
	return m.DescriptionField
}

func (m *Quote) SetDescription(s string) {
	m.DescriptionField = s
}

func (m *Quote) EffectiveQuoteCompletionDate() strfmt.DateTime {
	return m.EffectiveQuoteCompletionDateField
}

func (m *Quote) SetEffectiveQuoteCompletionDate(time strfmt.DateTime) {
	m.EffectiveQuoteCompletionDateField = time
}

func (m *Quote) ExpectedFulfillmentStartDate() strfmt.Date {
	return m.ExpectedFulfillmentStartDateField
}

func (m *Quote) SetExpectedFulfillmentStartDate(date strfmt.Date) {
	m.ExpectedFulfillmentStartDateField = date
}

func (m *Quote) ExpectedQuoteCompletionDate() strfmt.Date {
	return m.ExpectedQuoteCompletionDateField
}

func (m *Quote) SetExpectedQuoteCompletionDate(date strfmt.Date) {
	m.ExpectedQuoteCompletionDateField = date
}

func (m *Quote) ExternalID() string {
	return m.ExternalIDField
}

func (m *Quote) SetExternalID(s string) {
	m.ExternalIDField = s
}

func (m *Quote) Href() string {
	return m.HrefField
}

func (m *Quote) SetHref(s string) {
	m.HrefField = s
}

func (m *Quote) ID() string {
	return m.IDField
}

func (m *Quote) SetID(s string) {
	m.IDField = s
}

func (m *Quote) InstantSyncQuoting() bool {
	return *m.InstantSyncQuotingField
}

func (m *Quote) SetInstantSyncQuoting(b bool) {
	*m.InstantSyncQuotingField = b
}

func (m *Quote) Note() []*models.Note {
	var notes []*models.Note
	for _, note := range m.NoteField {
		to := &models.Note{}
		if err := util.Convert(note, to); err == nil {
			notes = append(notes)
		}
	}
	return notes
}

func (m *Quote) SetNote(notes []*models.Note) {
	var r []*Note
	for _, n := range notes {
		to := &Note{}
		if err := util.Convert(n, to); err == nil {
			to.ID = swag.String(xid.New().String())
			r = append(r, to)
		}
	}
	m.NoteField = r
}

func (m *Quote) ProjectID() string {
	return m.ProjectID()
}

func (m *Quote) SetProjectID(s string) {
	m.ProjectIDField = s
}

func (m *Quote) QuoteDate() strfmt.DateTime {
	return m.QuoteDateField
}

func (m *Quote) SetQuoteDate(time strfmt.DateTime) {
	m.QuoteDateField = time
}

func (m *Quote) QuoteItem() []*models.QuoteItem {
	if m.QuoteItemField == nil {
		return nil
	} else {
		var to []*models.QuoteItem
		_ = util.Convert(m.QuoteItemField, &to)
		return to
	}
}

func (m *Quote) SetQuoteItem(items []*models.QuoteItem) {
	if items != nil {
		var to []*QuoteItem
		_ = util.Convert(items, &to)
		m.QuoteItemField = to
	}
}

func (m *Quote) QuoteLevel() models.QuoteLevel {
	return m.QuoteLevelField
}

func (m *Quote) SetQuoteLevel(level models.QuoteLevel) {
	m.QuoteLevelField = level
}

func (m *Quote) RelatedParty() []*models.RelatedParty {
	var to = []*models.RelatedParty{}
	_ = util.Convert(m.RelatedPartyField, &to)
	return to
}

func (m *Quote) SetRelatedParty(parties []*models.RelatedParty) {
	var to []*RelatedParty
	_ = util.Convert(parties, &to)
	m.RelatedPartyField = to
}

func (m *Quote) RequestedQuoteCompletionDate() *strfmt.DateTime {
	return m.RequestedQuoteCompletionDateField
}

func (m *Quote) SetRequestedQuoteCompletionDate(time *strfmt.DateTime) {
	m.RequestedQuoteCompletionDateField = time
}

func (m *Quote) State() models.QuoteStateType {
	return m.StateField
}

func (m *Quote) SetState(stateType models.QuoteStateType) {
	m.StateField = stateType
}

func (m *Quote) ValidFor() *models.TimePeriod {
	return m.ValidForField
}

func (m *Quote) SetValidFor(period *models.TimePeriod) {
	m.ValidForField = period
}

func (m *Quote) ToQuoteSummaryView() models.QuoteSummaryView {
	to := &QuoteSummaryView{}
	to.AtBaseTypeField = m.AtBaseTypeField
	to.AtSchemaLocationField = m.AtSchemaLocationField
	to.AtTypeField = m.Type
	to.CategoryField = ""
	to.EffectiveQuoteCompletionDateField = m.EffectiveQuoteCompletionDateField
	to.ExpectedFulfillmentStartDateField = m.ExpectedFulfillmentStartDateField
	to.ExpectedQuoteCompletionDateField = m.ExpectedQuoteCompletionDateField
	to.ExternalIdField = m.ExternalIDField
	to.HrefField = m.HrefField
	to.IDField = m.IDField
	to.QuoteDateField = m.QuoteDateField
	to.QuoteItemField = m.QuoteItem()
	to.QuoteLevelField = m.QuoteLevelField
	var roles []*models.RelatedPartyRole
	//FIXME: How to convert this??? store role in db???
	for _, p := range m.RelatedParty() {
		r := ""
		if len(p.Role) > 0 {
			r = p.Role[0]
		}
		roles = append(roles, &models.RelatedPartyRole{
			AtReferredType: p.AtReferredType,
			ID:             swag.StringValue(p.ID),
			RelatedParty:   p,
			Role:           swag.String(r),
		})
	}
	to.RelatedPartyRoleField = roles
	to.RequestedQuoteCompletionDateField = m.RequestedQuoteCompletionDateField
	to.StateField = models.QuoteState(m.StateField)
	to.ValidForField = m.ValidForField
	return to
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
	Note []*Note `json:"note" gorm:"foreignkey:ID"`

	// product
	Product *Product `json:"product,omitempty" gorm:"foreignkey:ID"`

	// product offering
	ProductOffering *models.ProductOfferingRef `json:"productOffering,omitempty" gorm:"foreignkey:ID"`

	// qualification
	Qualification *models.ProductOfferingQualificationRef `json:"qualification,omitempty" gorm:"foreignkey:ID"`

	// quote item relationship
	QuoteItemRelationship []*models.QuoteItemRelationship `json:"quoteItemRelationship" gorm:"foreignkey:ID"`

	// related party
	RelatedParty []*RelatedParty `json:"relatedParty" gorm:"foreignkey:ID"`

	// requested quote item term
	RequestedQuoteItemTerm *models.ItemTerm `json:"requestedQuoteItemTerm,omitempty" gorm:"embedded"`
}
