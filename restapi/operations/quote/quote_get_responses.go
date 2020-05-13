// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-sonata-server/models"
)

// QuoteGetOKCode is the HTTP code returned for type QuoteGetOK
const QuoteGetOKCode int = 200

/*QuoteGetOK Ok

swagger:response quoteGetOK
*/
type QuoteGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Quote `json:"body,omitempty"`
}

// NewQuoteGetOK creates QuoteGetOK with default headers values
func NewQuoteGetOK() *QuoteGetOK {

	return &QuoteGetOK{}
}

// WithPayload adds the payload to the quote get o k response
func (o *QuoteGetOK) WithPayload(payload *models.Quote) *QuoteGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get o k response
func (o *QuoteGetOK) SetPayload(payload *models.Quote) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetBadRequestCode is the HTTP code returned for type QuoteGetBadRequest
const QuoteGetBadRequestCode int = 400

/*QuoteGetBadRequest Bad Request

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

swagger:response quoteGetBadRequest
*/
type QuoteGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetBadRequest creates QuoteGetBadRequest with default headers values
func NewQuoteGetBadRequest() *QuoteGetBadRequest {

	return &QuoteGetBadRequest{}
}

// WithPayload adds the payload to the quote get bad request response
func (o *QuoteGetBadRequest) WithPayload(payload *models.ErrorRepresentation) *QuoteGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get bad request response
func (o *QuoteGetBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetUnauthorizedCode is the HTTP code returned for type QuoteGetUnauthorized
const QuoteGetUnauthorizedCode int = 401

/*QuoteGetUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response quoteGetUnauthorized
*/
type QuoteGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetUnauthorized creates QuoteGetUnauthorized with default headers values
func NewQuoteGetUnauthorized() *QuoteGetUnauthorized {

	return &QuoteGetUnauthorized{}
}

// WithPayload adds the payload to the quote get unauthorized response
func (o *QuoteGetUnauthorized) WithPayload(payload *models.ErrorRepresentation) *QuoteGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get unauthorized response
func (o *QuoteGetUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetForbiddenCode is the HTTP code returned for type QuoteGetForbidden
const QuoteGetForbiddenCode int = 403

/*QuoteGetForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response quoteGetForbidden
*/
type QuoteGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetForbidden creates QuoteGetForbidden with default headers values
func NewQuoteGetForbidden() *QuoteGetForbidden {

	return &QuoteGetForbidden{}
}

// WithPayload adds the payload to the quote get forbidden response
func (o *QuoteGetForbidden) WithPayload(payload *models.ErrorRepresentation) *QuoteGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get forbidden response
func (o *QuoteGetForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetNotFoundCode is the HTTP code returned for type QuoteGetNotFound
const QuoteGetNotFoundCode int = 404

/*QuoteGetNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response quoteGetNotFound
*/
type QuoteGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetNotFound creates QuoteGetNotFound with default headers values
func NewQuoteGetNotFound() *QuoteGetNotFound {

	return &QuoteGetNotFound{}
}

// WithPayload adds the payload to the quote get not found response
func (o *QuoteGetNotFound) WithPayload(payload *models.ErrorRepresentation) *QuoteGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get not found response
func (o *QuoteGetNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetMethodNotAllowedCode is the HTTP code returned for type QuoteGetMethodNotAllowed
const QuoteGetMethodNotAllowedCode int = 405

/*QuoteGetMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response quoteGetMethodNotAllowed
*/
type QuoteGetMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetMethodNotAllowed creates QuoteGetMethodNotAllowed with default headers values
func NewQuoteGetMethodNotAllowed() *QuoteGetMethodNotAllowed {

	return &QuoteGetMethodNotAllowed{}
}

// WithPayload adds the payload to the quote get method not allowed response
func (o *QuoteGetMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *QuoteGetMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get method not allowed response
func (o *QuoteGetMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetUnprocessableEntityCode is the HTTP code returned for type QuoteGetUnprocessableEntity
const QuoteGetUnprocessableEntityCode int = 422

/*QuoteGetUnprocessableEntity Unprocessable entity

Functional error

swagger:response quoteGetUnprocessableEntity
*/
type QuoteGetUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetUnprocessableEntity creates QuoteGetUnprocessableEntity with default headers values
func NewQuoteGetUnprocessableEntity() *QuoteGetUnprocessableEntity {

	return &QuoteGetUnprocessableEntity{}
}

// WithPayload adds the payload to the quote get unprocessable entity response
func (o *QuoteGetUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *QuoteGetUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get unprocessable entity response
func (o *QuoteGetUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetInternalServerErrorCode is the HTTP code returned for type QuoteGetInternalServerError
const QuoteGetInternalServerErrorCode int = 500

/*QuoteGetInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response quoteGetInternalServerError
*/
type QuoteGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetInternalServerError creates QuoteGetInternalServerError with default headers values
func NewQuoteGetInternalServerError() *QuoteGetInternalServerError {

	return &QuoteGetInternalServerError{}
}

// WithPayload adds the payload to the quote get internal server error response
func (o *QuoteGetInternalServerError) WithPayload(payload *models.ErrorRepresentation) *QuoteGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get internal server error response
func (o *QuoteGetInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteGetServiceUnavailableCode is the HTTP code returned for type QuoteGetServiceUnavailable
const QuoteGetServiceUnavailableCode int = 503

/*QuoteGetServiceUnavailable Service Unavailable



swagger:response quoteGetServiceUnavailable
*/
type QuoteGetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteGetServiceUnavailable creates QuoteGetServiceUnavailable with default headers values
func NewQuoteGetServiceUnavailable() *QuoteGetServiceUnavailable {

	return &QuoteGetServiceUnavailable{}
}

// WithPayload adds the payload to the quote get service unavailable response
func (o *QuoteGetServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *QuoteGetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote get service unavailable response
func (o *QuoteGetServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteGetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
