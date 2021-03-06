// Code generated by go-swagger; DO NOT EDIT.

package product_offering_qualification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new product offering qualification API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for product offering qualification API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ProductOfferingQualificationCreate(params *ProductOfferingQualificationCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProductOfferingQualificationCreateCreated, error)

	ProductOfferingQualificationFind(params *ProductOfferingQualificationFindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProductOfferingQualificationFindOK, error)

	ProductOfferingQualificationGet(params *ProductOfferingQualificationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProductOfferingQualificationGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProductOfferingQualificationCreate creates a product offering qualification

  A request initiated by the Buyer to determine whether the Seller can feasibly deliver a particular Product (or Products) to a specific set of geographic locations specified by a set of Site/Address filter criteria. The Seller also provides estimated time intervals to complete these deliveries.
*/
func (a *Client) ProductOfferingQualificationCreate(params *ProductOfferingQualificationCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProductOfferingQualificationCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductOfferingQualificationCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "productOfferingQualificationCreate",
		Method:             "POST",
		PathPattern:        "/productOfferingQualificationManagement/v3/productOfferingQualification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductOfferingQualificationCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductOfferingQualificationCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for productOfferingQualificationCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductOfferingQualificationFind retrieves a list of product offering qualifications based on a set of criteria

  The Buyer requests a list of POQs (in any state) from the Seller based on a set of POQ filter criteria.  For each POQ returned, the Seller also provides a POQ Identifier that uniquely identifies this POQ within the Seller.
*/
func (a *Client) ProductOfferingQualificationFind(params *ProductOfferingQualificationFindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProductOfferingQualificationFindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductOfferingQualificationFindParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "productOfferingQualificationFind",
		Method:             "GET",
		PathPattern:        "/productOfferingQualificationManagement/v3/productOfferingQualification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductOfferingQualificationFindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductOfferingQualificationFindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for productOfferingQualificationFind: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProductOfferingQualificationGet gets a product offering qualification based on its id

  The Buyer requests the full details of a single Product Offering Qualification based on a POQ identifier.
*/
func (a *Client) ProductOfferingQualificationGet(params *ProductOfferingQualificationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProductOfferingQualificationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProductOfferingQualificationGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "productOfferingQualificationGet",
		Method:             "GET",
		PathPattern:        "/productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProductOfferingQualificationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProductOfferingQualificationGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for productOfferingQualificationGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
