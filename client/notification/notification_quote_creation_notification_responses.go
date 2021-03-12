// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// NotificationQuoteCreationNotificationReader is a Reader for the NotificationQuoteCreationNotification structure.
type NotificationQuoteCreationNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationQuoteCreationNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNotificationQuoteCreationNotificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationQuoteCreationNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationQuoteCreationNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationQuoteCreationNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationQuoteCreationNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewNotificationQuoteCreationNotificationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewNotificationQuoteCreationNotificationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationQuoteCreationNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNotificationQuoteCreationNotificationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationQuoteCreationNotificationNoContent creates a NotificationQuoteCreationNotificationNoContent with default headers values
func NewNotificationQuoteCreationNotificationNoContent() *NotificationQuoteCreationNotificationNoContent {
	return &NotificationQuoteCreationNotificationNoContent{}
}

/* NotificationQuoteCreationNotificationNoContent describes a response with status code 204, with default header values.

Success
*/
type NotificationQuoteCreationNotificationNoContent struct {
}

func (o *NotificationQuoteCreationNotificationNoContent) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationNoContent ", 204)
}

func (o *NotificationQuoteCreationNotificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationQuoteCreationNotificationBadRequest creates a NotificationQuoteCreationNotificationBadRequest with default headers values
func NewNotificationQuoteCreationNotificationBadRequest() *NotificationQuoteCreationNotificationBadRequest {
	return &NotificationQuoteCreationNotificationBadRequest{}
}

/* NotificationQuoteCreationNotificationBadRequest describes a response with status code 400, with default header values.

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
type NotificationQuoteCreationNotificationBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationQuoteCreationNotificationBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationUnauthorized creates a NotificationQuoteCreationNotificationUnauthorized with default headers values
func NewNotificationQuoteCreationNotificationUnauthorized() *NotificationQuoteCreationNotificationUnauthorized {
	return &NotificationQuoteCreationNotificationUnauthorized{}
}

/* NotificationQuoteCreationNotificationUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type NotificationQuoteCreationNotificationUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationQuoteCreationNotificationUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationForbidden creates a NotificationQuoteCreationNotificationForbidden with default headers values
func NewNotificationQuoteCreationNotificationForbidden() *NotificationQuoteCreationNotificationForbidden {
	return &NotificationQuoteCreationNotificationForbidden{}
}

/* NotificationQuoteCreationNotificationForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type NotificationQuoteCreationNotificationForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationForbidden) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationForbidden  %+v", 403, o.Payload)
}
func (o *NotificationQuoteCreationNotificationForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationNotFound creates a NotificationQuoteCreationNotificationNotFound with default headers values
func NewNotificationQuoteCreationNotificationNotFound() *NotificationQuoteCreationNotificationNotFound {
	return &NotificationQuoteCreationNotificationNotFound{}
}

/* NotificationQuoteCreationNotificationNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type NotificationQuoteCreationNotificationNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationNotFound) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationNotFound  %+v", 404, o.Payload)
}
func (o *NotificationQuoteCreationNotificationNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationMethodNotAllowed creates a NotificationQuoteCreationNotificationMethodNotAllowed with default headers values
func NewNotificationQuoteCreationNotificationMethodNotAllowed() *NotificationQuoteCreationNotificationMethodNotAllowed {
	return &NotificationQuoteCreationNotificationMethodNotAllowed{}
}

/* NotificationQuoteCreationNotificationMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type NotificationQuoteCreationNotificationMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *NotificationQuoteCreationNotificationMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationUnprocessableEntity creates a NotificationQuoteCreationNotificationUnprocessableEntity with default headers values
func NewNotificationQuoteCreationNotificationUnprocessableEntity() *NotificationQuoteCreationNotificationUnprocessableEntity {
	return &NotificationQuoteCreationNotificationUnprocessableEntity{}
}

/* NotificationQuoteCreationNotificationUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error
*/
type NotificationQuoteCreationNotificationUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *NotificationQuoteCreationNotificationUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationInternalServerError creates a NotificationQuoteCreationNotificationInternalServerError with default headers values
func NewNotificationQuoteCreationNotificationInternalServerError() *NotificationQuoteCreationNotificationInternalServerError {
	return &NotificationQuoteCreationNotificationInternalServerError{}
}

/* NotificationQuoteCreationNotificationInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type NotificationQuoteCreationNotificationInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationInternalServerError  %+v", 500, o.Payload)
}
func (o *NotificationQuoteCreationNotificationInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationQuoteCreationNotificationServiceUnavailable creates a NotificationQuoteCreationNotificationServiceUnavailable with default headers values
func NewNotificationQuoteCreationNotificationServiceUnavailable() *NotificationQuoteCreationNotificationServiceUnavailable {
	return &NotificationQuoteCreationNotificationServiceUnavailable{}
}

/* NotificationQuoteCreationNotificationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type NotificationQuoteCreationNotificationServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationQuoteCreationNotificationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /quoteNotification/v1/notification/quoteCreationNotification][%d] notificationQuoteCreationNotificationServiceUnavailable  %+v", 503, o.Payload)
}
func (o *NotificationQuoteCreationNotificationServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationQuoteCreationNotificationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
