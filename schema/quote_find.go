package schema

import (
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

type QuoteFind struct {
	DescriptionField string `json:"description,omitempty"`

	EffectiveQuoteCompletionDateField strfmt.DateTime `json:"effectiveQuoteCompletionDate`

	ExpectedQuoteCompletionDateField strfmt.Date `json:"expectedQuoteCompletionDate"`

	ExternalIDField string `json:"externalId,omitempty"`

	HrefField string `json:"href"`

	IDField string `json:"id"`

	ProjectIDField string `json:"projectId"`

	QuoteDateField strfmt.DateTime `json:"quoteDate"`

	QuoteLevelField models.QuoteLevel `json:"quoteLevel"`

	RequestedQuoteCompletionDateField *strfmt.DateTime `json:"requestedQuoteCompletionDate"`

	StateField models.QuoteStateType `json:"state"`
}

// Description gets the description of this polymorphic type
func (m *QuoteFind) Description() string {
	return m.DescriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *QuoteFind) SetDescription(val string) {
	m.DescriptionField = val
}

// EffectiveQuoteCompletionDate gets the effective quote completion date of this polymorphic type
func (m *QuoteFind) EffectiveQuoteCompletionDate() strfmt.DateTime {
	return m.EffectiveQuoteCompletionDateField
}

// SetEffectiveQuoteCompletionDate sets the effective quote completion date of this polymorphic type
func (m *QuoteFind) SetEffectiveQuoteCompletionDate(val strfmt.DateTime) {
	m.EffectiveQuoteCompletionDateField = val
}

// ExpectedQuoteCompletionDate gets the expected quote completion date of this polymorphic type
func (m *QuoteFind) ExpectedQuoteCompletionDate() strfmt.Date {
	return m.ExpectedQuoteCompletionDateField
}

// SetExpectedQuoteCompletionDate sets the expected quote completion date of this polymorphic type
func (m *QuoteFind) SetExpectedQuoteCompletionDate(val strfmt.Date) {
	m.ExpectedQuoteCompletionDateField = val
}

// ExternalID gets the external Id of this polymorphic type
func (m *QuoteFind) ExternalID() string {
	return m.ExternalIDField
}

// SetExternalID sets the external Id of this polymorphic type
func (m *QuoteFind) SetExternalID(val string) {
	m.ExternalIDField = val
}

// Href gets the href of this polymorphic type
func (m *QuoteFind) Href() string {
	return m.HrefField
}

// SetHref sets the href of this polymorphic type
func (m *QuoteFind) SetHref(val string) {
	m.HrefField = val
}

// ID gets the id of this polymorphic type
func (m *QuoteFind) ID() string {
	return "Quote_Find"
}

// SetID sets the id of this polymorphic type
func (m *QuoteFind) SetID(val string) {
}

// ProjectID gets the project Id of this polymorphic type
func (m *QuoteFind) ProjectID() string {
	return m.ProjectIDField
}

// SetProjectID sets the project Id of this polymorphic type
func (m *QuoteFind) SetProjectID(val string) {
	m.ProjectIDField = val
}

// QuoteDate gets the quote date of this polymorphic type
func (m *QuoteFind) QuoteDate() strfmt.DateTime {
	return m.QuoteDateField
}

// SetQuoteDate sets the quote date of this polymorphic type
func (m *QuoteFind) SetQuoteDate(val strfmt.DateTime) {
	m.QuoteDateField = val
}

// QuoteLevel gets the quote level of this polymorphic type
func (m *QuoteFind) QuoteLevel() models.QuoteLevel {
	return m.QuoteLevelField
}

// SetQuoteLevel sets the quote level of this polymorphic type
func (m *QuoteFind) SetQuoteLevel(val models.QuoteLevel) {
	m.QuoteLevelField = val
}

// RequestedQuoteCompletionDate gets the requested quote completion date of this polymorphic type
func (m *QuoteFind) RequestedQuoteCompletionDate() *strfmt.DateTime {
	return m.RequestedQuoteCompletionDateField
}

// SetRequestedQuoteCompletionDate sets the requested quote completion date of this polymorphic type
func (m *QuoteFind) SetRequestedQuoteCompletionDate(val *strfmt.DateTime) {
	m.RequestedQuoteCompletionDateField = val
}

// State gets the state of this polymorphic type
func (m *QuoteFind) State() models.QuoteStateType {
	return m.StateField
}

// SetState sets the state of this polymorphic type
func (m *QuoteFind) SetState(val models.QuoteStateType) {
	m.StateField = val
}
