// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// QuoteManagementHubFindReader is a Reader for the QuoteManagementHubFind structure.
type QuoteManagementHubFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteManagementHubFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteManagementHubFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteManagementHubFindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuoteManagementHubFindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuoteManagementHubFindForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuoteManagementHubFindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewQuoteManagementHubFindMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQuoteManagementHubFindUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuoteManagementHubFindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQuoteManagementHubFindServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuoteManagementHubFindOK creates a QuoteManagementHubFindOK with default headers values
func NewQuoteManagementHubFindOK() *QuoteManagementHubFindOK {
	return &QuoteManagementHubFindOK{}
}

/* QuoteManagementHubFindOK describes a response with status code 200, with default header values.

Ok
*/
type QuoteManagementHubFindOK struct {
	Payload []*models.Hub
}

func (o *QuoteManagementHubFindOK) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindOK  %+v", 200, o.Payload)
}
func (o *QuoteManagementHubFindOK) GetPayload() []*models.Hub {
	return o.Payload
}

func (o *QuoteManagementHubFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindBadRequest creates a QuoteManagementHubFindBadRequest with default headers values
func NewQuoteManagementHubFindBadRequest() *QuoteManagementHubFindBadRequest {
	return &QuoteManagementHubFindBadRequest{}
}

/* QuoteManagementHubFindBadRequest describes a response with status code 400, with default header values.

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
type QuoteManagementHubFindBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindBadRequest) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteManagementHubFindBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindUnauthorized creates a QuoteManagementHubFindUnauthorized with default headers values
func NewQuoteManagementHubFindUnauthorized() *QuoteManagementHubFindUnauthorized {
	return &QuoteManagementHubFindUnauthorized{}
}

/* QuoteManagementHubFindUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type QuoteManagementHubFindUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindUnauthorized  %+v", 401, o.Payload)
}
func (o *QuoteManagementHubFindUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindForbidden creates a QuoteManagementHubFindForbidden with default headers values
func NewQuoteManagementHubFindForbidden() *QuoteManagementHubFindForbidden {
	return &QuoteManagementHubFindForbidden{}
}

/* QuoteManagementHubFindForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type QuoteManagementHubFindForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindForbidden) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindForbidden  %+v", 403, o.Payload)
}
func (o *QuoteManagementHubFindForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindNotFound creates a QuoteManagementHubFindNotFound with default headers values
func NewQuoteManagementHubFindNotFound() *QuoteManagementHubFindNotFound {
	return &QuoteManagementHubFindNotFound{}
}

/* QuoteManagementHubFindNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type QuoteManagementHubFindNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindNotFound) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindNotFound  %+v", 404, o.Payload)
}
func (o *QuoteManagementHubFindNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindMethodNotAllowed creates a QuoteManagementHubFindMethodNotAllowed with default headers values
func NewQuoteManagementHubFindMethodNotAllowed() *QuoteManagementHubFindMethodNotAllowed {
	return &QuoteManagementHubFindMethodNotAllowed{}
}

/* QuoteManagementHubFindMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type QuoteManagementHubFindMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *QuoteManagementHubFindMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindUnprocessableEntity creates a QuoteManagementHubFindUnprocessableEntity with default headers values
func NewQuoteManagementHubFindUnprocessableEntity() *QuoteManagementHubFindUnprocessableEntity {
	return &QuoteManagementHubFindUnprocessableEntity{}
}

/* QuoteManagementHubFindUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error
*/
type QuoteManagementHubFindUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *QuoteManagementHubFindUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindInternalServerError creates a QuoteManagementHubFindInternalServerError with default headers values
func NewQuoteManagementHubFindInternalServerError() *QuoteManagementHubFindInternalServerError {
	return &QuoteManagementHubFindInternalServerError{}
}

/* QuoteManagementHubFindInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type QuoteManagementHubFindInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindInternalServerError  %+v", 500, o.Payload)
}
func (o *QuoteManagementHubFindInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubFindServiceUnavailable creates a QuoteManagementHubFindServiceUnavailable with default headers values
func NewQuoteManagementHubFindServiceUnavailable() *QuoteManagementHubFindServiceUnavailable {
	return &QuoteManagementHubFindServiceUnavailable{}
}

/* QuoteManagementHubFindServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type QuoteManagementHubFindServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubFindServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /quoteManagement/v2/hub][%d] quoteManagementHubFindServiceUnavailable  %+v", 503, o.Payload)
}
func (o *QuoteManagementHubFindServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubFindServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
