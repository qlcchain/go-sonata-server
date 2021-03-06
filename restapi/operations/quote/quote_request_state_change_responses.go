// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-sonata-server/models"
)

// QuoteRequestStateChangeOKCode is the HTTP code returned for type QuoteRequestStateChangeOK
const QuoteRequestStateChangeOKCode int = 200

/*QuoteRequestStateChangeOK Ok

swagger:response quoteRequestStateChangeOK
*/
type QuoteRequestStateChangeOK struct {

	/*
	  In: Body
	*/
	Payload *models.ChangeQuoteStateResponse `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeOK creates QuoteRequestStateChangeOK with default headers values
func NewQuoteRequestStateChangeOK() *QuoteRequestStateChangeOK {

	return &QuoteRequestStateChangeOK{}
}

// WithPayload adds the payload to the quote request state change o k response
func (o *QuoteRequestStateChangeOK) WithPayload(payload *models.ChangeQuoteStateResponse) *QuoteRequestStateChangeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change o k response
func (o *QuoteRequestStateChangeOK) SetPayload(payload *models.ChangeQuoteStateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeBadRequestCode is the HTTP code returned for type QuoteRequestStateChangeBadRequest
const QuoteRequestStateChangeBadRequestCode int = 400

/*QuoteRequestStateChangeBadRequest Bad Request

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

swagger:response quoteRequestStateChangeBadRequest
*/
type QuoteRequestStateChangeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeBadRequest creates QuoteRequestStateChangeBadRequest with default headers values
func NewQuoteRequestStateChangeBadRequest() *QuoteRequestStateChangeBadRequest {

	return &QuoteRequestStateChangeBadRequest{}
}

// WithPayload adds the payload to the quote request state change bad request response
func (o *QuoteRequestStateChangeBadRequest) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change bad request response
func (o *QuoteRequestStateChangeBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeUnauthorizedCode is the HTTP code returned for type QuoteRequestStateChangeUnauthorized
const QuoteRequestStateChangeUnauthorizedCode int = 401

/*QuoteRequestStateChangeUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response quoteRequestStateChangeUnauthorized
*/
type QuoteRequestStateChangeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeUnauthorized creates QuoteRequestStateChangeUnauthorized with default headers values
func NewQuoteRequestStateChangeUnauthorized() *QuoteRequestStateChangeUnauthorized {

	return &QuoteRequestStateChangeUnauthorized{}
}

// WithPayload adds the payload to the quote request state change unauthorized response
func (o *QuoteRequestStateChangeUnauthorized) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change unauthorized response
func (o *QuoteRequestStateChangeUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeForbiddenCode is the HTTP code returned for type QuoteRequestStateChangeForbidden
const QuoteRequestStateChangeForbiddenCode int = 403

/*QuoteRequestStateChangeForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response quoteRequestStateChangeForbidden
*/
type QuoteRequestStateChangeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeForbidden creates QuoteRequestStateChangeForbidden with default headers values
func NewQuoteRequestStateChangeForbidden() *QuoteRequestStateChangeForbidden {

	return &QuoteRequestStateChangeForbidden{}
}

// WithPayload adds the payload to the quote request state change forbidden response
func (o *QuoteRequestStateChangeForbidden) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change forbidden response
func (o *QuoteRequestStateChangeForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeNotFoundCode is the HTTP code returned for type QuoteRequestStateChangeNotFound
const QuoteRequestStateChangeNotFoundCode int = 404

/*QuoteRequestStateChangeNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response quoteRequestStateChangeNotFound
*/
type QuoteRequestStateChangeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeNotFound creates QuoteRequestStateChangeNotFound with default headers values
func NewQuoteRequestStateChangeNotFound() *QuoteRequestStateChangeNotFound {

	return &QuoteRequestStateChangeNotFound{}
}

// WithPayload adds the payload to the quote request state change not found response
func (o *QuoteRequestStateChangeNotFound) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change not found response
func (o *QuoteRequestStateChangeNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeMethodNotAllowedCode is the HTTP code returned for type QuoteRequestStateChangeMethodNotAllowed
const QuoteRequestStateChangeMethodNotAllowedCode int = 405

/*QuoteRequestStateChangeMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response quoteRequestStateChangeMethodNotAllowed
*/
type QuoteRequestStateChangeMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeMethodNotAllowed creates QuoteRequestStateChangeMethodNotAllowed with default headers values
func NewQuoteRequestStateChangeMethodNotAllowed() *QuoteRequestStateChangeMethodNotAllowed {

	return &QuoteRequestStateChangeMethodNotAllowed{}
}

// WithPayload adds the payload to the quote request state change method not allowed response
func (o *QuoteRequestStateChangeMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change method not allowed response
func (o *QuoteRequestStateChangeMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeUnprocessableEntityCode is the HTTP code returned for type QuoteRequestStateChangeUnprocessableEntity
const QuoteRequestStateChangeUnprocessableEntityCode int = 422

/*QuoteRequestStateChangeUnprocessableEntity Unprocessable entity

Functional error





 - code: 100
message: Quote current status is incompatible with requested quote state change
description:


 - code: 101
message: Quote external Id provided did not match
description:

swagger:response quoteRequestStateChangeUnprocessableEntity
*/
type QuoteRequestStateChangeUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeUnprocessableEntity creates QuoteRequestStateChangeUnprocessableEntity with default headers values
func NewQuoteRequestStateChangeUnprocessableEntity() *QuoteRequestStateChangeUnprocessableEntity {

	return &QuoteRequestStateChangeUnprocessableEntity{}
}

// WithPayload adds the payload to the quote request state change unprocessable entity response
func (o *QuoteRequestStateChangeUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change unprocessable entity response
func (o *QuoteRequestStateChangeUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeInternalServerErrorCode is the HTTP code returned for type QuoteRequestStateChangeInternalServerError
const QuoteRequestStateChangeInternalServerErrorCode int = 500

/*QuoteRequestStateChangeInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response quoteRequestStateChangeInternalServerError
*/
type QuoteRequestStateChangeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeInternalServerError creates QuoteRequestStateChangeInternalServerError with default headers values
func NewQuoteRequestStateChangeInternalServerError() *QuoteRequestStateChangeInternalServerError {

	return &QuoteRequestStateChangeInternalServerError{}
}

// WithPayload adds the payload to the quote request state change internal server error response
func (o *QuoteRequestStateChangeInternalServerError) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change internal server error response
func (o *QuoteRequestStateChangeInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// QuoteRequestStateChangeServiceUnavailableCode is the HTTP code returned for type QuoteRequestStateChangeServiceUnavailable
const QuoteRequestStateChangeServiceUnavailableCode int = 503

/*QuoteRequestStateChangeServiceUnavailable Service Unavailable



swagger:response quoteRequestStateChangeServiceUnavailable
*/
type QuoteRequestStateChangeServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewQuoteRequestStateChangeServiceUnavailable creates QuoteRequestStateChangeServiceUnavailable with default headers values
func NewQuoteRequestStateChangeServiceUnavailable() *QuoteRequestStateChangeServiceUnavailable {

	return &QuoteRequestStateChangeServiceUnavailable{}
}

// WithPayload adds the payload to the quote request state change service unavailable response
func (o *QuoteRequestStateChangeServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *QuoteRequestStateChangeServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the quote request state change service unavailable response
func (o *QuoteRequestStateChangeServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QuoteRequestStateChangeServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
