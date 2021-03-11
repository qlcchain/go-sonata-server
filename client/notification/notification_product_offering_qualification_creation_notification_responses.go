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

// NotificationProductOfferingQualificationCreationNotificationReader is a Reader for the NotificationProductOfferingQualificationCreationNotification structure.
type NotificationProductOfferingQualificationCreationNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationProductOfferingQualificationCreationNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNotificationProductOfferingQualificationCreationNotificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationProductOfferingQualificationCreationNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationProductOfferingQualificationCreationNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationProductOfferingQualificationCreationNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationProductOfferingQualificationCreationNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewNotificationProductOfferingQualificationCreationNotificationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewNotificationProductOfferingQualificationCreationNotificationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNotificationProductOfferingQualificationCreationNotificationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationProductOfferingQualificationCreationNotificationNoContent creates a NotificationProductOfferingQualificationCreationNotificationNoContent with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationNoContent() *NotificationProductOfferingQualificationCreationNotificationNoContent {
	return &NotificationProductOfferingQualificationCreationNotificationNoContent{}
}

/* NotificationProductOfferingQualificationCreationNotificationNoContent describes a response with status code 204, with default header values.

No Content
*/
type NotificationProductOfferingQualificationCreationNotificationNoContent struct {
}

func (o *NotificationProductOfferingQualificationCreationNotificationNoContent) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationNoContent ", 204)
}

func (o *NotificationProductOfferingQualificationCreationNotificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationBadRequest creates a NotificationProductOfferingQualificationCreationNotificationBadRequest with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationBadRequest() *NotificationProductOfferingQualificationCreationNotificationBadRequest {
	return &NotificationProductOfferingQualificationCreationNotificationBadRequest{}
}

/* NotificationProductOfferingQualificationCreationNotificationBadRequest describes a response with status code 400, with default header values.

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
type NotificationProductOfferingQualificationCreationNotificationBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationUnauthorized creates a NotificationProductOfferingQualificationCreationNotificationUnauthorized with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationUnauthorized() *NotificationProductOfferingQualificationCreationNotificationUnauthorized {
	return &NotificationProductOfferingQualificationCreationNotificationUnauthorized{}
}

/* NotificationProductOfferingQualificationCreationNotificationUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type NotificationProductOfferingQualificationCreationNotificationUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationForbidden creates a NotificationProductOfferingQualificationCreationNotificationForbidden with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationForbidden() *NotificationProductOfferingQualificationCreationNotificationForbidden {
	return &NotificationProductOfferingQualificationCreationNotificationForbidden{}
}

/* NotificationProductOfferingQualificationCreationNotificationForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type NotificationProductOfferingQualificationCreationNotificationForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationForbidden) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationForbidden  %+v", 403, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationNotFound creates a NotificationProductOfferingQualificationCreationNotificationNotFound with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationNotFound() *NotificationProductOfferingQualificationCreationNotificationNotFound {
	return &NotificationProductOfferingQualificationCreationNotificationNotFound{}
}

/* NotificationProductOfferingQualificationCreationNotificationNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type NotificationProductOfferingQualificationCreationNotificationNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationNotFound) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationNotFound  %+v", 404, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationRequestTimeout creates a NotificationProductOfferingQualificationCreationNotificationRequestTimeout with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationRequestTimeout() *NotificationProductOfferingQualificationCreationNotificationRequestTimeout {
	return &NotificationProductOfferingQualificationCreationNotificationRequestTimeout{}
}

/* NotificationProductOfferingQualificationCreationNotificationRequestTimeout describes a response with status code 408, with default header values.

 Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type NotificationProductOfferingQualificationCreationNotificationRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationRequestTimeout  %+v", 408, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationUnprocessableEntity creates a NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationUnprocessableEntity() *NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity {
	return &NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity{}
}

/* NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error
*/
type NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationProductOfferingQualificationCreationNotificationServiceUnavailable creates a NotificationProductOfferingQualificationCreationNotificationServiceUnavailable with default headers values
func NewNotificationProductOfferingQualificationCreationNotificationServiceUnavailable() *NotificationProductOfferingQualificationCreationNotificationServiceUnavailable {
	return &NotificationProductOfferingQualificationCreationNotificationServiceUnavailable{}
}

/* NotificationProductOfferingQualificationCreationNotificationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type NotificationProductOfferingQualificationCreationNotificationServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *NotificationProductOfferingQualificationCreationNotificationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationNotification/v3/notification/productOfferingQualificationCreationNotification][%d] notificationProductOfferingQualificationCreationNotificationServiceUnavailable  %+v", 503, o.Payload)
}
func (o *NotificationProductOfferingQualificationCreationNotificationServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *NotificationProductOfferingQualificationCreationNotificationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
