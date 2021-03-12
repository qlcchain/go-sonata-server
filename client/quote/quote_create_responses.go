// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// QuoteCreateReader is a Reader for the QuoteCreate structure.
type QuoteCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewQuoteCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuoteCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuoteCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuoteCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewQuoteCreateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQuoteCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuoteCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQuoteCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuoteCreateCreated creates a QuoteCreateCreated with default headers values
func NewQuoteCreateCreated() *QuoteCreateCreated {
	return &QuoteCreateCreated{}
}

/* QuoteCreateCreated describes a response with status code 201, with default header values.

Created
*/
type QuoteCreateCreated struct {
	Payload *models.Quote
}

func (o *QuoteCreateCreated) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateCreated  %+v", 201, o.Payload)
}
func (o *QuoteCreateCreated) GetPayload() *models.Quote {
	return o.Payload
}

func (o *QuoteCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Quote)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateBadRequest creates a QuoteCreateBadRequest with default headers values
func NewQuoteCreateBadRequest() *QuoteCreateBadRequest {
	return &QuoteCreateBadRequest{}
}

/* QuoteCreateBadRequest describes a response with status code 400, with default header values.

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
type QuoteCreateBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteCreateBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateUnauthorized creates a QuoteCreateUnauthorized with default headers values
func NewQuoteCreateUnauthorized() *QuoteCreateUnauthorized {
	return &QuoteCreateUnauthorized{}
}

/* QuoteCreateUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type QuoteCreateUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *QuoteCreateUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateForbidden creates a QuoteCreateForbidden with default headers values
func NewQuoteCreateForbidden() *QuoteCreateForbidden {
	return &QuoteCreateForbidden{}
}

/* QuoteCreateForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type QuoteCreateForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateForbidden  %+v", 403, o.Payload)
}
func (o *QuoteCreateForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateNotFound creates a QuoteCreateNotFound with default headers values
func NewQuoteCreateNotFound() *QuoteCreateNotFound {
	return &QuoteCreateNotFound{}
}

/* QuoteCreateNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type QuoteCreateNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateNotFound  %+v", 404, o.Payload)
}
func (o *QuoteCreateNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateMethodNotAllowed creates a QuoteCreateMethodNotAllowed with default headers values
func NewQuoteCreateMethodNotAllowed() *QuoteCreateMethodNotAllowed {
	return &QuoteCreateMethodNotAllowed{}
}

/* QuoteCreateMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type QuoteCreateMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *QuoteCreateMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateUnprocessableEntity creates a QuoteCreateUnprocessableEntity with default headers values
func NewQuoteCreateUnprocessableEntity() *QuoteCreateUnprocessableEntity {
	return &QuoteCreateUnprocessableEntity{}
}

/* QuoteCreateUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error





 - code: 101
message: A relatedPartyRole for Buyer role must be provided
description:


 - code: 103
message: A relatedPartyRole is incomplete (must have partyRole.id or role.id+relatedParty)
description:


 - code: 104
message: A relatedParty is incomplete (must have a party.id or a party.name)
description:


 - code: 105
message: Qualification Id provided did not exist or expired
description:


 - code: 106
message: At least a productOffering Id or a product.id or a productSpec.id must be provided in quote item
description:


 - code: 107
message: inconsistency between  requestedQuoteCompletionDate and expectedFulfillmentStartDate
description:


 - code: 108
message: Seller requires agreement under which the buyer is requesting but is missing from the request.
description:


 - code: 109
message: The agreement provided cannot be validated by the seller
description:


 - code: 110
message: Product requested is not part of the provided agreement
description:


 - code: 111
message: Mandatory product/item relationship is missing
description:
*/
type QuoteCreateUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *QuoteCreateUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateInternalServerError creates a QuoteCreateInternalServerError with default headers values
func NewQuoteCreateInternalServerError() *QuoteCreateInternalServerError {
	return &QuoteCreateInternalServerError{}
}

/* QuoteCreateInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type QuoteCreateInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *QuoteCreateInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteCreateServiceUnavailable creates a QuoteCreateServiceUnavailable with default headers values
func NewQuoteCreateServiceUnavailable() *QuoteCreateServiceUnavailable {
	return &QuoteCreateServiceUnavailable{}
}

/* QuoteCreateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type QuoteCreateServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote][%d] quoteCreateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *QuoteCreateServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
