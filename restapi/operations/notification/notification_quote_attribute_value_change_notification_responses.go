// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-sonata-server/models"
)

// NotificationQuoteAttributeValueChangeNotificationNoContentCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationNoContent
const NotificationQuoteAttributeValueChangeNotificationNoContentCode int = 204

/*NotificationQuoteAttributeValueChangeNotificationNoContent Success

swagger:response notificationQuoteAttributeValueChangeNotificationNoContent
*/
type NotificationQuoteAttributeValueChangeNotificationNoContent struct {
}

// NewNotificationQuoteAttributeValueChangeNotificationNoContent creates NotificationQuoteAttributeValueChangeNotificationNoContent with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationNoContent() *NotificationQuoteAttributeValueChangeNotificationNoContent {

	return &NotificationQuoteAttributeValueChangeNotificationNoContent{}
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// NotificationQuoteAttributeValueChangeNotificationBadRequestCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationBadRequest
const NotificationQuoteAttributeValueChangeNotificationBadRequestCode int = 400

/*NotificationQuoteAttributeValueChangeNotificationBadRequest Bad Request

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

swagger:response notificationQuoteAttributeValueChangeNotificationBadRequest
*/
type NotificationQuoteAttributeValueChangeNotificationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationBadRequest creates NotificationQuoteAttributeValueChangeNotificationBadRequest with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationBadRequest() *NotificationQuoteAttributeValueChangeNotificationBadRequest {

	return &NotificationQuoteAttributeValueChangeNotificationBadRequest{}
}

// WithPayload adds the payload to the notification quote attribute value change notification bad request response
func (o *NotificationQuoteAttributeValueChangeNotificationBadRequest) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification bad request response
func (o *NotificationQuoteAttributeValueChangeNotificationBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationUnauthorizedCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationUnauthorized
const NotificationQuoteAttributeValueChangeNotificationUnauthorizedCode int = 401

/*NotificationQuoteAttributeValueChangeNotificationUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response notificationQuoteAttributeValueChangeNotificationUnauthorized
*/
type NotificationQuoteAttributeValueChangeNotificationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationUnauthorized creates NotificationQuoteAttributeValueChangeNotificationUnauthorized with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationUnauthorized() *NotificationQuoteAttributeValueChangeNotificationUnauthorized {

	return &NotificationQuoteAttributeValueChangeNotificationUnauthorized{}
}

// WithPayload adds the payload to the notification quote attribute value change notification unauthorized response
func (o *NotificationQuoteAttributeValueChangeNotificationUnauthorized) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification unauthorized response
func (o *NotificationQuoteAttributeValueChangeNotificationUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationForbiddenCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationForbidden
const NotificationQuoteAttributeValueChangeNotificationForbiddenCode int = 403

/*NotificationQuoteAttributeValueChangeNotificationForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response notificationQuoteAttributeValueChangeNotificationForbidden
*/
type NotificationQuoteAttributeValueChangeNotificationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationForbidden creates NotificationQuoteAttributeValueChangeNotificationForbidden with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationForbidden() *NotificationQuoteAttributeValueChangeNotificationForbidden {

	return &NotificationQuoteAttributeValueChangeNotificationForbidden{}
}

// WithPayload adds the payload to the notification quote attribute value change notification forbidden response
func (o *NotificationQuoteAttributeValueChangeNotificationForbidden) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification forbidden response
func (o *NotificationQuoteAttributeValueChangeNotificationForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationNotFoundCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationNotFound
const NotificationQuoteAttributeValueChangeNotificationNotFoundCode int = 404

/*NotificationQuoteAttributeValueChangeNotificationNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response notificationQuoteAttributeValueChangeNotificationNotFound
*/
type NotificationQuoteAttributeValueChangeNotificationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationNotFound creates NotificationQuoteAttributeValueChangeNotificationNotFound with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationNotFound() *NotificationQuoteAttributeValueChangeNotificationNotFound {

	return &NotificationQuoteAttributeValueChangeNotificationNotFound{}
}

// WithPayload adds the payload to the notification quote attribute value change notification not found response
func (o *NotificationQuoteAttributeValueChangeNotificationNotFound) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification not found response
func (o *NotificationQuoteAttributeValueChangeNotificationNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationMethodNotAllowedCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed
const NotificationQuoteAttributeValueChangeNotificationMethodNotAllowedCode int = 405

/*NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response notificationQuoteAttributeValueChangeNotificationMethodNotAllowed
*/
type NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationMethodNotAllowed creates NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationMethodNotAllowed() *NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed {

	return &NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed{}
}

// WithPayload adds the payload to the notification quote attribute value change notification method not allowed response
func (o *NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification method not allowed response
func (o *NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationUnprocessableEntityCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity
const NotificationQuoteAttributeValueChangeNotificationUnprocessableEntityCode int = 422

/*NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity Unprocessable entity

Functional error

swagger:response notificationQuoteAttributeValueChangeNotificationUnprocessableEntity
*/
type NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationUnprocessableEntity creates NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationUnprocessableEntity() *NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity {

	return &NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity{}
}

// WithPayload adds the payload to the notification quote attribute value change notification unprocessable entity response
func (o *NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification unprocessable entity response
func (o *NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationInternalServerErrorCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationInternalServerError
const NotificationQuoteAttributeValueChangeNotificationInternalServerErrorCode int = 500

/*NotificationQuoteAttributeValueChangeNotificationInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response notificationQuoteAttributeValueChangeNotificationInternalServerError
*/
type NotificationQuoteAttributeValueChangeNotificationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationInternalServerError creates NotificationQuoteAttributeValueChangeNotificationInternalServerError with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationInternalServerError() *NotificationQuoteAttributeValueChangeNotificationInternalServerError {

	return &NotificationQuoteAttributeValueChangeNotificationInternalServerError{}
}

// WithPayload adds the payload to the notification quote attribute value change notification internal server error response
func (o *NotificationQuoteAttributeValueChangeNotificationInternalServerError) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification internal server error response
func (o *NotificationQuoteAttributeValueChangeNotificationInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NotificationQuoteAttributeValueChangeNotificationServiceUnavailableCode is the HTTP code returned for type NotificationQuoteAttributeValueChangeNotificationServiceUnavailable
const NotificationQuoteAttributeValueChangeNotificationServiceUnavailableCode int = 503

/*NotificationQuoteAttributeValueChangeNotificationServiceUnavailable Service Unavailable



swagger:response notificationQuoteAttributeValueChangeNotificationServiceUnavailable
*/
type NotificationQuoteAttributeValueChangeNotificationServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewNotificationQuoteAttributeValueChangeNotificationServiceUnavailable creates NotificationQuoteAttributeValueChangeNotificationServiceUnavailable with default headers values
func NewNotificationQuoteAttributeValueChangeNotificationServiceUnavailable() *NotificationQuoteAttributeValueChangeNotificationServiceUnavailable {

	return &NotificationQuoteAttributeValueChangeNotificationServiceUnavailable{}
}

// WithPayload adds the payload to the notification quote attribute value change notification service unavailable response
func (o *NotificationQuoteAttributeValueChangeNotificationServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *NotificationQuoteAttributeValueChangeNotificationServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification quote attribute value change notification service unavailable response
func (o *NotificationQuoteAttributeValueChangeNotificationServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationQuoteAttributeValueChangeNotificationServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
