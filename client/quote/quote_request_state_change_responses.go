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

// QuoteRequestStateChangeReader is a Reader for the QuoteRequestStateChange structure.
type QuoteRequestStateChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteRequestStateChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteRequestStateChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteRequestStateChangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuoteRequestStateChangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuoteRequestStateChangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuoteRequestStateChangeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewQuoteRequestStateChangeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQuoteRequestStateChangeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuoteRequestStateChangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQuoteRequestStateChangeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuoteRequestStateChangeOK creates a QuoteRequestStateChangeOK with default headers values
func NewQuoteRequestStateChangeOK() *QuoteRequestStateChangeOK {
	return &QuoteRequestStateChangeOK{}
}

/* QuoteRequestStateChangeOK describes a response with status code 200, with default header values.

Ok
*/
type QuoteRequestStateChangeOK struct {
	Payload *models.ChangeQuoteStateResponse
}

func (o *QuoteRequestStateChangeOK) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeOK  %+v", 200, o.Payload)
}
func (o *QuoteRequestStateChangeOK) GetPayload() *models.ChangeQuoteStateResponse {
	return o.Payload
}

func (o *QuoteRequestStateChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeQuoteStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeBadRequest creates a QuoteRequestStateChangeBadRequest with default headers values
func NewQuoteRequestStateChangeBadRequest() *QuoteRequestStateChangeBadRequest {
	return &QuoteRequestStateChangeBadRequest{}
}

/* QuoteRequestStateChangeBadRequest describes a response with status code 400, with default header values.

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
type QuoteRequestStateChangeBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteRequestStateChangeBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeUnauthorized creates a QuoteRequestStateChangeUnauthorized with default headers values
func NewQuoteRequestStateChangeUnauthorized() *QuoteRequestStateChangeUnauthorized {
	return &QuoteRequestStateChangeUnauthorized{}
}

/* QuoteRequestStateChangeUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type QuoteRequestStateChangeUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeUnauthorized  %+v", 401, o.Payload)
}
func (o *QuoteRequestStateChangeUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeForbidden creates a QuoteRequestStateChangeForbidden with default headers values
func NewQuoteRequestStateChangeForbidden() *QuoteRequestStateChangeForbidden {
	return &QuoteRequestStateChangeForbidden{}
}

/* QuoteRequestStateChangeForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type QuoteRequestStateChangeForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeForbidden) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeForbidden  %+v", 403, o.Payload)
}
func (o *QuoteRequestStateChangeForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeNotFound creates a QuoteRequestStateChangeNotFound with default headers values
func NewQuoteRequestStateChangeNotFound() *QuoteRequestStateChangeNotFound {
	return &QuoteRequestStateChangeNotFound{}
}

/* QuoteRequestStateChangeNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type QuoteRequestStateChangeNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeNotFound) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeNotFound  %+v", 404, o.Payload)
}
func (o *QuoteRequestStateChangeNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeMethodNotAllowed creates a QuoteRequestStateChangeMethodNotAllowed with default headers values
func NewQuoteRequestStateChangeMethodNotAllowed() *QuoteRequestStateChangeMethodNotAllowed {
	return &QuoteRequestStateChangeMethodNotAllowed{}
}

/* QuoteRequestStateChangeMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type QuoteRequestStateChangeMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *QuoteRequestStateChangeMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeUnprocessableEntity creates a QuoteRequestStateChangeUnprocessableEntity with default headers values
func NewQuoteRequestStateChangeUnprocessableEntity() *QuoteRequestStateChangeUnprocessableEntity {
	return &QuoteRequestStateChangeUnprocessableEntity{}
}

/* QuoteRequestStateChangeUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error





 - code: 100
message: Quote current status is incompatible with requested quote state change
description:


 - code: 101
message: Quote external Id provided did not match
description:
*/
type QuoteRequestStateChangeUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *QuoteRequestStateChangeUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeInternalServerError creates a QuoteRequestStateChangeInternalServerError with default headers values
func NewQuoteRequestStateChangeInternalServerError() *QuoteRequestStateChangeInternalServerError {
	return &QuoteRequestStateChangeInternalServerError{}
}

/* QuoteRequestStateChangeInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type QuoteRequestStateChangeInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeInternalServerError  %+v", 500, o.Payload)
}
func (o *QuoteRequestStateChangeInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteRequestStateChangeServiceUnavailable creates a QuoteRequestStateChangeServiceUnavailable with default headers values
func NewQuoteRequestStateChangeServiceUnavailable() *QuoteRequestStateChangeServiceUnavailable {
	return &QuoteRequestStateChangeServiceUnavailable{}
}

/* QuoteRequestStateChangeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type QuoteRequestStateChangeServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteRequestStateChangeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /quoteManagement/v2/quote/requestStateChange][%d] quoteRequestStateChangeServiceUnavailable  %+v", 503, o.Payload)
}
func (o *QuoteRequestStateChangeServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteRequestStateChangeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
