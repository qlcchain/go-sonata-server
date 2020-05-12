// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quote API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	QuoteCreate(params *QuoteCreateParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteCreateCreated, error)

	QuoteFind(params *QuoteFindParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteFindOK, error)

	QuoteGet(params *QuoteGetParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteGetOK, error)

	QuoteRequestStateChange(params *QuoteRequestStateChangeParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteRequestStateChangeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  QuoteCreate creates a quote

  This operation is used to create a new quote.
*/
func (a *Client) QuoteCreate(params *QuoteCreateParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteCreate",
		Method:             "POST",
		PathPattern:        "/quoteManagement/v2/quote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuoteCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for quoteCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QuoteFind lists quotes

  This operation is used to retrieve quote information using filter criteria.
*/
func (a *Client) QuoteFind(params *QuoteFindParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteFindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteFindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteFind",
		Method:             "GET",
		PathPattern:        "/quoteManagement/v2/quote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteFindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuoteFindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for quoteFind: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QuoteGet retrieves a quote

  This operation is used to retrieve quote information using the ID.
*/
func (a *Client) QuoteGet(params *QuoteGetParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteGet",
		Method:             "GET",
		PathPattern:        "/quoteManagement/v2/quote/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuoteGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for quoteGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QuoteRequestStateChange requests a quote state change

  Request from buyer to cancel or reject a quote.
When seller receive cancel request, seller will shift quote state to CANCELLED (no change on order item state)
When seller receive reject request, seller will shift quote state to REJECTED (no change on order item state)
*/
func (a *Client) QuoteRequestStateChange(params *QuoteRequestStateChangeParams, authInfo runtime.ClientAuthInfoWriter) (*QuoteRequestStateChangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteRequestStateChangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quoteRequestStateChange",
		Method:             "POST",
		PathPattern:        "/quoteManagement/v2/quote/requestStateChange",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteRequestStateChangeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuoteRequestStateChangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for quoteRequestStateChange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
