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

// NotificationProductOrderStateChangeNotificationReader is a Reader for the NotificationProductOrderStateChangeNotification structure.
type NotificationProductOrderStateChangeNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationProductOrderStateChangeNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNotificationProductOrderStateChangeNotificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationProductOrderStateChangeNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationProductOrderStateChangeNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationProductOrderStateChangeNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationProductOrderStateChangeNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewNotificationProductOrderStateChangeNotificationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewNotificationProductOrderStateChangeNotificationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewNotificationProductOrderStateChangeNotificationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationProductOrderStateChangeNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNotificationProductOrderStateChangeNotificationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationProductOrderStateChangeNotificationNoContent creates a NotificationProductOrderStateChangeNotificationNoContent with default headers values
func NewNotificationProductOrderStateChangeNotificationNoContent() *NotificationProductOrderStateChangeNotificationNoContent {
	return &NotificationProductOrderStateChangeNotificationNoContent{}
}

/* NotificationProductOrderStateChangeNotificationNoContent describes a response with status code 204, with default header values.

No Content
*/
type NotificationProductOrderStateChangeNotificationNoContent struct {
}

func (o *NotificationProductOrderStateChangeNotificationNoContent) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationNoContent ", 204)
}

func (o *NotificationProductOrderStateChangeNotificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationProductOrderStateChangeNotificationBadRequest creates a NotificationProductOrderStateChangeNotificationBadRequest with default headers values
func NewNotificationProductOrderStateChangeNotificationBadRequest() *NotificationProductOrderStateChangeNotificationBadRequest {
	return &NotificationProductOrderStateChangeNotificationBadRequest{}
}

/* NotificationProductOrderStateChangeNotificationBadRequest describes a response with status code 400, with default header values.

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
type NotificationProductOrderStateChangeNotificationBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationUnauthorized creates a NotificationProductOrderStateChangeNotificationUnauthorized with default headers values
func NewNotificationProductOrderStateChangeNotificationUnauthorized() *NotificationProductOrderStateChangeNotificationUnauthorized {
	return &NotificationProductOrderStateChangeNotificationUnauthorized{}
}

/* NotificationProductOrderStateChangeNotificationUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type NotificationProductOrderStateChangeNotificationUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationForbidden creates a NotificationProductOrderStateChangeNotificationForbidden with default headers values
func NewNotificationProductOrderStateChangeNotificationForbidden() *NotificationProductOrderStateChangeNotificationForbidden {
	return &NotificationProductOrderStateChangeNotificationForbidden{}
}

/* NotificationProductOrderStateChangeNotificationForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type NotificationProductOrderStateChangeNotificationForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationForbidden) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationForbidden  %+v", 403, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationNotFound creates a NotificationProductOrderStateChangeNotificationNotFound with default headers values
func NewNotificationProductOrderStateChangeNotificationNotFound() *NotificationProductOrderStateChangeNotificationNotFound {
	return &NotificationProductOrderStateChangeNotificationNotFound{}
}

/* NotificationProductOrderStateChangeNotificationNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type NotificationProductOrderStateChangeNotificationNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationNotFound) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationNotFound  %+v", 404, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationMethodNotAllowed creates a NotificationProductOrderStateChangeNotificationMethodNotAllowed with default headers values
func NewNotificationProductOrderStateChangeNotificationMethodNotAllowed() *NotificationProductOrderStateChangeNotificationMethodNotAllowed {
	return &NotificationProductOrderStateChangeNotificationMethodNotAllowed{}
}

/* NotificationProductOrderStateChangeNotificationMethodNotAllowed describes a response with status code 405, with default header values.

 Method Not Allowed

List of supported error codes:
- 61: Method not allowed
*/
type NotificationProductOrderStateChangeNotificationMethodNotAllowed struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationMethodNotAllowed) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationRequestTimeout creates a NotificationProductOrderStateChangeNotificationRequestTimeout with default headers values
func NewNotificationProductOrderStateChangeNotificationRequestTimeout() *NotificationProductOrderStateChangeNotificationRequestTimeout {
	return &NotificationProductOrderStateChangeNotificationRequestTimeout{}
}

/* NotificationProductOrderStateChangeNotificationRequestTimeout describes a response with status code 408, with default header values.

 Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type NotificationProductOrderStateChangeNotificationRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationRequestTimeout  %+v", 408, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationUnprocessableEntity creates a NotificationProductOrderStateChangeNotificationUnprocessableEntity with default headers values
func NewNotificationProductOrderStateChangeNotificationUnprocessableEntity() *NotificationProductOrderStateChangeNotificationUnprocessableEntity {
	return &NotificationProductOrderStateChangeNotificationUnprocessableEntity{}
}

/* NotificationProductOrderStateChangeNotificationUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error
*/
type NotificationProductOrderStateChangeNotificationUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationInternalServerError creates a NotificationProductOrderStateChangeNotificationInternalServerError with default headers values
func NewNotificationProductOrderStateChangeNotificationInternalServerError() *NotificationProductOrderStateChangeNotificationInternalServerError {
	return &NotificationProductOrderStateChangeNotificationInternalServerError{}
}

/* NotificationProductOrderStateChangeNotificationInternalServerError describes a response with status code 500, with default header values.

 Internal Server Error

List of supported error codes:
- 1: Internal error
*/
type NotificationProductOrderStateChangeNotificationInternalServerError struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationInternalServerError  %+v", 500, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationInternalServerError) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOrderStateChangeNotificationServiceUnavailable creates a NotificationProductOrderStateChangeNotificationServiceUnavailable with default headers values
func NewNotificationProductOrderStateChangeNotificationServiceUnavailable() *NotificationProductOrderStateChangeNotificationServiceUnavailable {
	return &NotificationProductOrderStateChangeNotificationServiceUnavailable{}
}

/* NotificationProductOrderStateChangeNotificationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type NotificationProductOrderStateChangeNotificationServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOrderStateChangeNotificationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /productOrderNotification/v3/notification/productOrderStateChangeNotification][%d] notificationProductOrderStateChangeNotificationServiceUnavailable  %+v", 503, o.Payload)
}
func (o *NotificationProductOrderStateChangeNotificationServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOrderStateChangeNotificationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
