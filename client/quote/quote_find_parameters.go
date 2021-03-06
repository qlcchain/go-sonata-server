// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewQuoteFindParams creates a new QuoteFindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuoteFindParams() *QuoteFindParams {
	return &QuoteFindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteFindParamsWithTimeout creates a new QuoteFindParams object
// with the ability to set a timeout on a request.
func NewQuoteFindParamsWithTimeout(timeout time.Duration) *QuoteFindParams {
	return &QuoteFindParams{
		timeout: timeout,
	}
}

// NewQuoteFindParamsWithContext creates a new QuoteFindParams object
// with the ability to set a context for a request.
func NewQuoteFindParamsWithContext(ctx context.Context) *QuoteFindParams {
	return &QuoteFindParams{
		Context: ctx,
	}
}

// NewQuoteFindParamsWithHTTPClient creates a new QuoteFindParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuoteFindParamsWithHTTPClient(client *http.Client) *QuoteFindParams {
	return &QuoteFindParams{
		HTTPClient: client,
	}
}

/* QuoteFindParams contains all the parameters to send to the API endpoint
   for the quote find operation.

   Typically these are written to a http.Request.
*/
type QuoteFindParams struct {

	/* ExternalID.

	   ID given by the consumer and only understandable by him (to facilitate his searches afterwards)
	*/
	ExternalID *string

	/* Limit.

	   Requested number of resources to be provided in response requested by client
	*/
	Limit *string

	/* Offset.

	   Requested index for start of resources to be provided in response requested by client
	*/
	Offset *string

	/* ProjectID.

	   This value MAY be assigned by the Buyer/Seller to identify a project the quoting request is associated with.
	*/
	ProjectID *string

	/* QuoteCompletionDateGt.

	   Date when the quote was completed -  greater than

	   Format: date-time
	*/
	QuoteCompletionDateGt *strfmt.DateTime

	/* QuoteCompletionDateLt.

	   Date when the quote was completed -  lower than

	   Format: date-time
	*/
	QuoteCompletionDateLt *strfmt.DateTime

	/* QuoteDateGt.

	   Date when the quote was created - greater than

	   Format: date-time
	*/
	QuoteDateGt *strfmt.DateTime

	/* QuoteDateLt.

	   Date when the quote was created - lower than

	   Format: date-time
	*/
	QuoteDateLt *strfmt.DateTime

	/* QuoteLevel.

	   Level of the quote - could be indicative, budgetary or firm
	*/
	QuoteLevel *string

	/* State.

	   State of the Quote
	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quote find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteFindParams) WithDefaults() *QuoteFindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quote find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteFindParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the quote find params
func (o *QuoteFindParams) WithTimeout(timeout time.Duration) *QuoteFindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote find params
func (o *QuoteFindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote find params
func (o *QuoteFindParams) WithContext(ctx context.Context) *QuoteFindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote find params
func (o *QuoteFindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote find params
func (o *QuoteFindParams) WithHTTPClient(client *http.Client) *QuoteFindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote find params
func (o *QuoteFindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalID adds the externalID to the quote find params
func (o *QuoteFindParams) WithExternalID(externalID *string) *QuoteFindParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the quote find params
func (o *QuoteFindParams) SetExternalID(externalID *string) {
	o.ExternalID = externalID
}

// WithLimit adds the limit to the quote find params
func (o *QuoteFindParams) WithLimit(limit *string) *QuoteFindParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the quote find params
func (o *QuoteFindParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithOffset adds the offset to the quote find params
func (o *QuoteFindParams) WithOffset(offset *string) *QuoteFindParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the quote find params
func (o *QuoteFindParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithProjectID adds the projectID to the quote find params
func (o *QuoteFindParams) WithProjectID(projectID *string) *QuoteFindParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the quote find params
func (o *QuoteFindParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithQuoteCompletionDateGt adds the quoteCompletionDateGt to the quote find params
func (o *QuoteFindParams) WithQuoteCompletionDateGt(quoteCompletionDateGt *strfmt.DateTime) *QuoteFindParams {
	o.SetQuoteCompletionDateGt(quoteCompletionDateGt)
	return o
}

// SetQuoteCompletionDateGt adds the quoteCompletionDateGt to the quote find params
func (o *QuoteFindParams) SetQuoteCompletionDateGt(quoteCompletionDateGt *strfmt.DateTime) {
	o.QuoteCompletionDateGt = quoteCompletionDateGt
}

// WithQuoteCompletionDateLt adds the quoteCompletionDateLt to the quote find params
func (o *QuoteFindParams) WithQuoteCompletionDateLt(quoteCompletionDateLt *strfmt.DateTime) *QuoteFindParams {
	o.SetQuoteCompletionDateLt(quoteCompletionDateLt)
	return o
}

// SetQuoteCompletionDateLt adds the quoteCompletionDateLt to the quote find params
func (o *QuoteFindParams) SetQuoteCompletionDateLt(quoteCompletionDateLt *strfmt.DateTime) {
	o.QuoteCompletionDateLt = quoteCompletionDateLt
}

// WithQuoteDateGt adds the quoteDateGt to the quote find params
func (o *QuoteFindParams) WithQuoteDateGt(quoteDateGt *strfmt.DateTime) *QuoteFindParams {
	o.SetQuoteDateGt(quoteDateGt)
	return o
}

// SetQuoteDateGt adds the quoteDateGt to the quote find params
func (o *QuoteFindParams) SetQuoteDateGt(quoteDateGt *strfmt.DateTime) {
	o.QuoteDateGt = quoteDateGt
}

// WithQuoteDateLt adds the quoteDateLt to the quote find params
func (o *QuoteFindParams) WithQuoteDateLt(quoteDateLt *strfmt.DateTime) *QuoteFindParams {
	o.SetQuoteDateLt(quoteDateLt)
	return o
}

// SetQuoteDateLt adds the quoteDateLt to the quote find params
func (o *QuoteFindParams) SetQuoteDateLt(quoteDateLt *strfmt.DateTime) {
	o.QuoteDateLt = quoteDateLt
}

// WithQuoteLevel adds the quoteLevel to the quote find params
func (o *QuoteFindParams) WithQuoteLevel(quoteLevel *string) *QuoteFindParams {
	o.SetQuoteLevel(quoteLevel)
	return o
}

// SetQuoteLevel adds the quoteLevel to the quote find params
func (o *QuoteFindParams) SetQuoteLevel(quoteLevel *string) {
	o.QuoteLevel = quoteLevel
}

// WithState adds the state to the quote find params
func (o *QuoteFindParams) WithState(state *string) *QuoteFindParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the quote find params
func (o *QuoteFindParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteFindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExternalID != nil {

		// query param externalId
		var qrExternalID string

		if o.ExternalID != nil {
			qrExternalID = *o.ExternalID
		}
		qExternalID := qrExternalID
		if qExternalID != "" {

			if err := r.SetQueryParam("externalId", qExternalID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.QuoteCompletionDateGt != nil {

		// query param quoteCompletionDate.gt
		var qrQuoteCompletionDateGt strfmt.DateTime

		if o.QuoteCompletionDateGt != nil {
			qrQuoteCompletionDateGt = *o.QuoteCompletionDateGt
		}
		qQuoteCompletionDateGt := qrQuoteCompletionDateGt.String()
		if qQuoteCompletionDateGt != "" {

			if err := r.SetQueryParam("quoteCompletionDate.gt", qQuoteCompletionDateGt); err != nil {
				return err
			}
		}
	}

	if o.QuoteCompletionDateLt != nil {

		// query param quoteCompletionDate.lt
		var qrQuoteCompletionDateLt strfmt.DateTime

		if o.QuoteCompletionDateLt != nil {
			qrQuoteCompletionDateLt = *o.QuoteCompletionDateLt
		}
		qQuoteCompletionDateLt := qrQuoteCompletionDateLt.String()
		if qQuoteCompletionDateLt != "" {

			if err := r.SetQueryParam("quoteCompletionDate.lt", qQuoteCompletionDateLt); err != nil {
				return err
			}
		}
	}

	if o.QuoteDateGt != nil {

		// query param quoteDate.gt
		var qrQuoteDateGt strfmt.DateTime

		if o.QuoteDateGt != nil {
			qrQuoteDateGt = *o.QuoteDateGt
		}
		qQuoteDateGt := qrQuoteDateGt.String()
		if qQuoteDateGt != "" {

			if err := r.SetQueryParam("quoteDate.gt", qQuoteDateGt); err != nil {
				return err
			}
		}
	}

	if o.QuoteDateLt != nil {

		// query param quoteDate.lt
		var qrQuoteDateLt strfmt.DateTime

		if o.QuoteDateLt != nil {
			qrQuoteDateLt = *o.QuoteDateLt
		}
		qQuoteDateLt := qrQuoteDateLt.String()
		if qQuoteDateLt != "" {

			if err := r.SetQueryParam("quoteDate.lt", qQuoteDateLt); err != nil {
				return err
			}
		}
	}

	if o.QuoteLevel != nil {

		// query param quoteLevel
		var qrQuoteLevel string

		if o.QuoteLevel != nil {
			qrQuoteLevel = *o.QuoteLevel
		}
		qQuoteLevel := qrQuoteLevel
		if qQuoteLevel != "" {

			if err := r.SetQueryParam("quoteLevel", qQuoteLevel); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
