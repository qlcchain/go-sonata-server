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

// QuoteManagementHubDeleteReader is a Reader for the QuoteManagementHubDelete structure.
type QuoteManagementHubDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteManagementHubDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewQuoteManagementHubDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteManagementHubDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuoteManagementHubDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQuoteManagementHubDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQuoteManagementHubDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewQuoteManagementHubDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQuoteManagementHubDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuoteManagementHubDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQuoteManagementHubDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuoteManagementHubDeleteNoContent creates a QuoteManagementHubDeleteNoContent with default headers values
func NewQuoteManagementHubDeleteNoContent() *QuoteManagementHubDeleteNoContent {
	return &QuoteManagementHubDeleteNoContent{}
}

/* QuoteManagementHubDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type QuoteManagementHubDeleteNoContent struct {
}

func (o *QuoteManagementHubDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteNoContent ", 204)
}

func (o *QuoteManagementHubDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQuoteManagementHubDeleteBadRequest creates a QuoteManagementHubDeleteBadRequest with default headers values
func NewQuoteManagementHubDeleteBadRequest() *QuoteManagementHubDeleteBadRequest {
	return &QuoteManagementHubDeleteBadRequest{}
}

/* QuoteManagementHubDeleteBadRequest describes a response with status code 400, with default header values.

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
type QuoteManagementHubDeleteBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteManagementHubDeleteBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteUnauthorized creates a QuoteManagementHubDeleteUnauthorized with default headers values
func NewQuoteManagementHubDeleteUnauthorized() *QuoteManagementHubDeleteUnauthorized {
	return &QuoteManagementHubDeleteUnauthorized{}
}

/* QuoteManagementHubDeleteUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type QuoteManagementHubDeleteUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *QuoteManagementHubDeleteUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteForbidden creates a QuoteManagementHubDeleteForbidden with default headers values
func NewQuoteManagementHubDeleteForbidden() *QuoteManagementHubDeleteForbidden {
	return &QuoteManagementHubDeleteForbidden{}
}

/* QuoteManagementHubDeleteForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type QuoteManagementHubDeleteForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteForbidden  %+v", 403, o.Payload)
}
func (o *QuoteManagementHubDeleteForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteNotFound creates a QuoteManagementHubDeleteNotFound with default headers values
func NewQuoteManagementHubDeleteNotFound() *QuoteManagementHubDeleteNotFound {
	return &QuoteManagementHubDeleteNotFound{}
}

/* QuoteManagementHubDeleteNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type QuoteManagementHubDeleteNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteNotFound  %+v", 404, o.Payload)
}
func (o *QuoteManagementHubDeleteNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteMethodNotAllowed creates a QuoteManagementHubDeleteMethodNotAllowed with default headers values
func NewQuoteManagementHubDeleteMethodNotAllowed() *QuoteManagementHubDeleteMethodNotAllowed {
	return &QuoteManagementHubDeleteMethodNotAllowed{}
}

/* QuoteManagementHubDeleteMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type QuoteManagementHubDeleteMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *QuoteManagementHubDeleteMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteUnprocessableEntity creates a QuoteManagementHubDeleteUnprocessableEntity with default headers values
func NewQuoteManagementHubDeleteUnprocessableEntity() *QuoteManagementHubDeleteUnprocessableEntity {
	return &QuoteManagementHubDeleteUnprocessableEntity{}
}

/* QuoteManagementHubDeleteUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error
*/
type QuoteManagementHubDeleteUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *QuoteManagementHubDeleteUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteInternalServerError creates a QuoteManagementHubDeleteInternalServerError with default headers values
func NewQuoteManagementHubDeleteInternalServerError() *QuoteManagementHubDeleteInternalServerError {
	return &QuoteManagementHubDeleteInternalServerError{}
}

/* QuoteManagementHubDeleteInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type QuoteManagementHubDeleteInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *QuoteManagementHubDeleteInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteManagementHubDeleteServiceUnavailable creates a QuoteManagementHubDeleteServiceUnavailable with default headers values
func NewQuoteManagementHubDeleteServiceUnavailable() *QuoteManagementHubDeleteServiceUnavailable {
	return &QuoteManagementHubDeleteServiceUnavailable{}
}

/* QuoteManagementHubDeleteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type QuoteManagementHubDeleteServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *QuoteManagementHubDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /quoteManagement/v2/hub/{HubId}][%d] quoteManagementHubDeleteServiceUnavailable  %+v", 503, o.Payload)
}
func (o *QuoteManagementHubDeleteServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *QuoteManagementHubDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
