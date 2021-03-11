// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new notification API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for notification API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NotificationProductOfferingQualificationCreationNotification(params *NotificationProductOfferingQualificationCreationNotificationParams, opts ...ClientOption) (*NotificationProductOfferingQualificationCreationNotificationNoContent, error)

	NotificationProductOfferingQualificationStateChangeNotification(params *NotificationProductOfferingQualificationStateChangeNotificationParams, opts ...ClientOption) (*NotificationProductOfferingQualificationStateChangeNotificationNoContent, error)

	NotificationProductOrderAttributeValueChangeNotification(params *NotificationProductOrderAttributeValueChangeNotificationParams, opts ...ClientOption) (*NotificationProductOrderAttributeValueChangeNotificationNoContent, error)

	NotificationProductOrderCreationNotification(params *NotificationProductOrderCreationNotificationParams, opts ...ClientOption) (*NotificationProductOrderCreationNotificationNoContent, error)

	NotificationProductOrderInformationRequiredNotification(params *NotificationProductOrderInformationRequiredNotificationParams, opts ...ClientOption) (*NotificationProductOrderInformationRequiredNotificationNoContent, error)

	NotificationProductOrderStateChangeNotification(params *NotificationProductOrderStateChangeNotificationParams, opts ...ClientOption) (*NotificationProductOrderStateChangeNotificationNoContent, error)

	NotificationQuoteAttributeValueChangeNotification(params *NotificationQuoteAttributeValueChangeNotificationParams, opts ...ClientOption) (*NotificationQuoteAttributeValueChangeNotificationNoContent, error)

	NotificationQuoteCreationNotification(params *NotificationQuoteCreationNotificationParams, opts ...ClientOption) (*NotificationQuoteCreationNotificationNoContent, error)

	NotificationQuoteInformationRequiredNotification(params *NotificationQuoteInformationRequiredNotificationParams, opts ...ClientOption) (*NotificationQuoteInformationRequiredNotificationNoContent, error)

	NotificationQuoteStateChangeNotification(params *NotificationQuoteStateChangeNotificationParams, opts ...ClientOption) (*NotificationQuoteStateChangeNotificationNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  NotificationProductOfferingQualificationCreationNotification products offering qualification creation notification structure

  Product Offering Qualification Creation Notification structure definition
*/
func (a *Client) NotificationProductOfferingQualificationCreationNotification(params *NotificationProductOfferingQualificationCreationNotificationParams, opts ...ClientOption) (*NotificationProductOfferingQualificationCreationNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOfferingQualificationCreationNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationProductOfferingQualificationCreationNotification",
		Method:             "POST",
		PathPattern:        "/productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOfferingQualificationCreationNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationProductOfferingQualificationCreationNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationProductOfferingQualificationCreationNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationProductOfferingQualificationStateChangeNotification products offering qualification state change notification structure

  Product Offering Qualification State Change Notification structure definition
*/
func (a *Client) NotificationProductOfferingQualificationStateChangeNotification(params *NotificationProductOfferingQualificationStateChangeNotificationParams, opts ...ClientOption) (*NotificationProductOfferingQualificationStateChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOfferingQualificationStateChangeNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationProductOfferingQualificationStateChangeNotification",
		Method:             "POST",
		PathPattern:        "/productOfferingQualificationNotification/v3/notification/productOfferingQualificationStateChangeNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOfferingQualificationStateChangeNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationProductOfferingQualificationStateChangeNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationProductOfferingQualificationStateChangeNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationProductOrderAttributeValueChangeNotification products order attribute value change structure

  Product Order attribute value change structure description
*/
func (a *Client) NotificationProductOrderAttributeValueChangeNotification(params *NotificationProductOrderAttributeValueChangeNotificationParams, opts ...ClientOption) (*NotificationProductOrderAttributeValueChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderAttributeValueChangeNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationProductOrderAttributeValueChangeNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderAttributeValueChangeNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderAttributeValueChangeNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationProductOrderAttributeValueChangeNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationProductOrderAttributeValueChangeNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationProductOrderCreationNotification products order creation notification structure

  Product order creation notification structure description
*/
func (a *Client) NotificationProductOrderCreationNotification(params *NotificationProductOrderCreationNotificationParams, opts ...ClientOption) (*NotificationProductOrderCreationNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderCreationNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationProductOrderCreationNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderCreationNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderCreationNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationProductOrderCreationNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationProductOrderCreationNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationProductOrderInformationRequiredNotification products order information required structure

  Product Order information required structure description
*/
func (a *Client) NotificationProductOrderInformationRequiredNotification(params *NotificationProductOrderInformationRequiredNotificationParams, opts ...ClientOption) (*NotificationProductOrderInformationRequiredNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderInformationRequiredNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationProductOrderInformationRequiredNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderInformationRequiredNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderInformationRequiredNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationProductOrderInformationRequiredNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationProductOrderInformationRequiredNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationProductOrderStateChangeNotification products order state change structure

  Product order state change structure description
*/
func (a *Client) NotificationProductOrderStateChangeNotification(params *NotificationProductOrderStateChangeNotificationParams, opts ...ClientOption) (*NotificationProductOrderStateChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderStateChangeNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationProductOrderStateChangeNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderStateChangeNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderStateChangeNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationProductOrderStateChangeNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationProductOrderStateChangeNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationQuoteAttributeValueChangeNotification quotes attribute value change notification structure

  Quote attribute value change notification structure description.
Attribute resourcePatch allows to target quote but also quoteItem - example: resourcePath":"/quote/42/quoteItem/3" is the item #3 of quote #42
Attribute fieldPath allows to target attribute with value changed.

Specific business errors for current operation will be encapsulated in

HTTP Response 422 Unprocessable entity

*/
func (a *Client) NotificationQuoteAttributeValueChangeNotification(params *NotificationQuoteAttributeValueChangeNotificationParams, opts ...ClientOption) (*NotificationQuoteAttributeValueChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteAttributeValueChangeNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationQuoteAttributeValueChangeNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/notification/quoteAttributeValueChangeNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteAttributeValueChangeNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationQuoteAttributeValueChangeNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationQuoteAttributeValueChangeNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationQuoteCreationNotification quotes creation notification structure

  Quote creation notification structure definition

Specific business errors for current operation will be encapsulated in

HTTP Response 422 Unprocessable entity

*/
func (a *Client) NotificationQuoteCreationNotification(params *NotificationQuoteCreationNotificationParams, opts ...ClientOption) (*NotificationQuoteCreationNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteCreationNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationQuoteCreationNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/notification/quoteCreationNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteCreationNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationQuoteCreationNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationQuoteCreationNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationQuoteInformationRequiredNotification quotes information required notification structure

  Quote information required notification structure description.
Attribute resourcePatch allows to target quote but also quoteItem - example: resourcePath":"/quote/42/quoteItem/3" is the item #3 of quote #42
Attribute fieldPath allows to target missing information: fieldPath":"missing=quote.relatedParty.Role.value&party.id=46" means role information is missing for party 46.

Specific business errors for current operation will be encapsulated in

HTTP Response 422 Unprocessable entity

*/
func (a *Client) NotificationQuoteInformationRequiredNotification(params *NotificationQuoteInformationRequiredNotificationParams, opts ...ClientOption) (*NotificationQuoteInformationRequiredNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteInformationRequiredNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationQuoteInformationRequiredNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/notification/quoteInformationRequiredNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteInformationRequiredNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationQuoteInformationRequiredNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationQuoteInformationRequiredNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NotificationQuoteStateChangeNotification quotes state change notification structure

  Quote state change notification structure description

Specific business errors for current operation will be encapsulated in

HTTP Response 422 Unprocessable entity

*/
func (a *Client) NotificationQuoteStateChangeNotification(params *NotificationQuoteStateChangeNotificationParams, opts ...ClientOption) (*NotificationQuoteStateChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteStateChangeNotificationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "notificationQuoteStateChangeNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/notification/quoteStateChangeNotification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteStateChangeNotificationReader{formats: a.formats},
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
	success, ok := result.(*NotificationQuoteStateChangeNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for notificationQuoteStateChangeNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
