// Code generated by go-swagger; DO NOT EDIT.

package product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// ProductOrderCreateReader is a Reader for the ProductOrderCreate structure.
type ProductOrderCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductOrderCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProductOrderCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProductOrderCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProductOrderCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProductOrderCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProductOrderCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewProductOrderCreateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewProductOrderCreateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProductOrderCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProductOrderCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewProductOrderCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProductOrderCreateCreated creates a ProductOrderCreateCreated with default headers values
func NewProductOrderCreateCreated() *ProductOrderCreateCreated {
	return &ProductOrderCreateCreated{}
}

/* ProductOrderCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ProductOrderCreateCreated struct {
	Payload *models.ProductOrder
}

func (o *ProductOrderCreateCreated) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateCreated  %+v", 201, o.Payload)
}
func (o *ProductOrderCreateCreated) GetPayload() *models.ProductOrder {
	return o.Payload
}

func (o *ProductOrderCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateBadRequest creates a ProductOrderCreateBadRequest with default headers values
func NewProductOrderCreateBadRequest() *ProductOrderCreateBadRequest {
	return &ProductOrderCreateBadRequest{}
}

/* ProductOrderCreateBadRequest describes a response with status code 400, with default header values.

 Bad Request

List of supported error codes:
- 20: Invalid URL parameter value
- 21: Missing body
- 22: Invalid body
- 23: Missing body field
- 24: Invalid body field
- 25: Missing header
- 26: Invalid header value
- 27: Missing query-string parameter
- 28: Invalid query-string parameter value
*/
type ProductOrderCreateBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateBadRequest  %+v", 400, o.Payload)
}
func (o *ProductOrderCreateBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateUnauthorized creates a ProductOrderCreateUnauthorized with default headers values
func NewProductOrderCreateUnauthorized() *ProductOrderCreateUnauthorized {
	return &ProductOrderCreateUnauthorized{}
}

/* ProductOrderCreateUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type ProductOrderCreateUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *ProductOrderCreateUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateForbidden creates a ProductOrderCreateForbidden with default headers values
func NewProductOrderCreateForbidden() *ProductOrderCreateForbidden {
	return &ProductOrderCreateForbidden{}
}

/* ProductOrderCreateForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type ProductOrderCreateForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateForbidden  %+v", 403, o.Payload)
}
func (o *ProductOrderCreateForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateNotFound creates a ProductOrderCreateNotFound with default headers values
func NewProductOrderCreateNotFound() *ProductOrderCreateNotFound {
	return &ProductOrderCreateNotFound{}
}

/* ProductOrderCreateNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type ProductOrderCreateNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateNotFound  %+v", 404, o.Payload)
}
func (o *ProductOrderCreateNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateMethodNotAllowed creates a ProductOrderCreateMethodNotAllowed with default headers values
func NewProductOrderCreateMethodNotAllowed() *ProductOrderCreateMethodNotAllowed {
	return &ProductOrderCreateMethodNotAllowed{}
}

/* ProductOrderCreateMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type ProductOrderCreateMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *ProductOrderCreateMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateRequestTimeout creates a ProductOrderCreateRequestTimeout with default headers values
func NewProductOrderCreateRequestTimeout() *ProductOrderCreateRequestTimeout {
	return &ProductOrderCreateRequestTimeout{}
}

/* ProductOrderCreateRequestTimeout describes a response with status code 408, with default header values.

 Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type ProductOrderCreateRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateRequestTimeout  %+v", 408, o.Payload)
}
func (o *ProductOrderCreateRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateUnprocessableEntity creates a ProductOrderCreateUnprocessableEntity with default headers values
func NewProductOrderCreateUnprocessableEntity() *ProductOrderCreateUnprocessableEntity {
	return &ProductOrderCreateUnprocessableEntity{}
}

/* ProductOrderCreateUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error





 - code: 100
message: Missing order item (minimum 1)
description: At least one order item must be provided


 - code: 101
message: Missing Buyer at order level
description: One and only one related partyRole with a "Buyer" role should be provided at the product order level.


 - code: 114
message: Missing Seller at order level
description: One and only one related partyRole with a "Seller" role should be provided at the product order level.


 - code: 102
message: A relatedParty is at the wrong level
description: The partyRole provided is not at the correct level - MEF allows to have
"Buyer", "Seller", "Billing Contact", "Order Contact", "Implementation Contact", "Technical Contact" roles at product order level and "UNI Site Contact", "UNI Alt Site Contact", "ENNI Site Contact", "ENNI Alt Site Contact" at product order item level.


 - code: 103
message: Missing Buyer Order Contact at order level
description: One and only one related partyRole with a "Order Contact" role should be provided at the product order level.
Buyer Order Contact name & Telephone number must be provided.


 - code: 104
message: Missing Buyer Implementation Contact at order level
description: One and only one related partyRole with a "Implementation Contact" role should be provided at the product order level.
Implementation Contact name & Telephone number must be provided.


 - code: 105
message: Missing Buyer Technical contact at order level
description: One and only one related partyRole with a "Technical Contact" role should be provided at the product order level.
Technical Contact name, Telephone number and email address must be provided.


 - code: 106
message: Address information must match place type
description: If place type is 'Formatted Address' : addrLine1, city, stateOrProvince, postCode and country must be there.
If place type is 'Fielded Address' : streetName, streetType, city, stateOrProvince, postCode and country must be there.


 - code: 107
message: postCode extension requires postcode value to be filled
description: A postCode extension must not be present without a postcode being present


 - code: 108
message: Product id is required for all OrderItem Actions other than INSTALL
description: If orderItemAction is not INSTALL, orderItem.product.id is mandatory


 - code: 109
message: Order Activity must match all OrderItem Actions for INSTALL
description: If orderActivity is set to INSTALL, all orderItemAction must be INSTALL


 - code: 110
message: Referred quote cannot be used for ordering due to its status
description: Quote cannot be used in the order if its status is in CANCELLED, UNABLE TO PROVIDE, REJECTED or EXPIRED state.


 - code: 111
message: Billing Account information must not be both at order level and order item level
description: Billing Account must not be present both at order header level and order item level.


 - code: 112
message: PricingMethod, pricing Reference & pricing term attributes  must not be both at order level and order item level
description: Pricing data  must not be present both at order header level and order item level.


 - code: 113
message: Referred Serviceability request is expired
description: Serviceability information are expired.


 - code: 114
message: A reccuring price is mentionned without a charge period.
description:


 - code: 115
message: Referred Quote is not existing
description:


 - code: 116
message: Referred ProductOfferingQualification is not existing
description:


 - code: 117
message: Product /item relationship is missing
description:


 - code: 118
message: Product Id refered in a relationship is not existing
description:
*/
type ProductOrderCreateUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ProductOrderCreateUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateInternalServerError creates a ProductOrderCreateInternalServerError with default headers values
func NewProductOrderCreateInternalServerError() *ProductOrderCreateInternalServerError {
	return &ProductOrderCreateInternalServerError{}
}

/* ProductOrderCreateInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type ProductOrderCreateInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *ProductOrderCreateInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOrderCreateServiceUnavailable creates a ProductOrderCreateServiceUnavailable with default headers values
func NewProductOrderCreateServiceUnavailable() *ProductOrderCreateServiceUnavailable {
	return &ProductOrderCreateServiceUnavailable{}
}

/* ProductOrderCreateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type ProductOrderCreateServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOrderCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /productOrderManagement/v3/productOrder][%d] productOrderCreateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ProductOrderCreateServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOrderCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
