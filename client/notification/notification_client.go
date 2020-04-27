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

// ClientService is the interface for Client methods
type ClientService interface {
	NotificationProductOrderAttributeValueChangeNotification(params *NotificationProductOrderAttributeValueChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderAttributeValueChangeNotificationNoContent, error)

	NotificationProductOrderCreationNotification(params *NotificationProductOrderCreationNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderCreationNotificationNoContent, error)

	NotificationProductOrderInformationRequiredNotification(params *NotificationProductOrderInformationRequiredNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderInformationRequiredNotificationNoContent, error)

	NotificationProductOrderStateChangeNotification(params *NotificationProductOrderStateChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderStateChangeNotificationNoContent, error)

	NotificationQuoteAttributeValueChangeNotification(params *NotificationQuoteAttributeValueChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteAttributeValueChangeNotificationNoContent, error)

	NotificationQuoteCreationNotification(params *NotificationQuoteCreationNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteCreationNotificationNoContent, error)

	NotificationQuoteInformationRequiredNotification(params *NotificationQuoteInformationRequiredNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteInformationRequiredNotificationNoContent, error)

	NotificationQuoteStateChangeNotification(params *NotificationQuoteStateChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteStateChangeNotificationNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  NotificationProductOrderAttributeValueChangeNotification products order attribute value change structure

  Product Order attribute value change structure description
*/
func (a *Client) NotificationProductOrderAttributeValueChangeNotification(params *NotificationProductOrderAttributeValueChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderAttributeValueChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderAttributeValueChangeNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationProductOrderAttributeValueChangeNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderAttributeValueChangeNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderAttributeValueChangeNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationProductOrderCreationNotification(params *NotificationProductOrderCreationNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderCreationNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderCreationNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationProductOrderCreationNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderCreationNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderCreationNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationProductOrderInformationRequiredNotification(params *NotificationProductOrderInformationRequiredNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderInformationRequiredNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderInformationRequiredNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationProductOrderInformationRequiredNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderInformationRequiredNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderInformationRequiredNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationProductOrderStateChangeNotification(params *NotificationProductOrderStateChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationProductOrderStateChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationProductOrderStateChangeNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationProductOrderStateChangeNotification",
		Method:             "POST",
		PathPattern:        "/productOrderNotification/v3/notification/productOrderStateChangeNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationProductOrderStateChangeNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationQuoteAttributeValueChangeNotification(params *NotificationQuoteAttributeValueChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteAttributeValueChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteAttributeValueChangeNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationQuoteAttributeValueChangeNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/quoteNotification/v1/notification/quoteAttributeValueChangeNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteAttributeValueChangeNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationQuoteCreationNotification(params *NotificationQuoteCreationNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteCreationNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteCreationNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationQuoteCreationNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/notification/quoteCreationNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteCreationNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationQuoteInformationRequiredNotification(params *NotificationQuoteInformationRequiredNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteInformationRequiredNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteInformationRequiredNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationQuoteInformationRequiredNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/quoteNotification/v1/notification/quoteInformationRequiredNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteInformationRequiredNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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
func (a *Client) NotificationQuoteStateChangeNotification(params *NotificationQuoteStateChangeNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*NotificationQuoteStateChangeNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationQuoteStateChangeNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notificationQuoteStateChangeNotification",
		Method:             "POST",
		PathPattern:        "/quoteNotification/v1/notification/quoteStateChangeNotification",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationQuoteStateChangeNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
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