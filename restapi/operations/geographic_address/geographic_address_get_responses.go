// Code generated by go-swagger; DO NOT EDIT.

package geographic_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-sonata-server/models"
)

// GeographicAddressGetOKCode is the HTTP code returned for type GeographicAddressGetOK
const GeographicAddressGetOKCode int = 200

/*GeographicAddressGetOK Ok

swagger:response geographicAddressGetOK
*/
type GeographicAddressGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.GeographicAddress `json:"body,omitempty"`
}

// NewGeographicAddressGetOK creates GeographicAddressGetOK with default headers values
func NewGeographicAddressGetOK() *GeographicAddressGetOK {

	return &GeographicAddressGetOK{}
}

// WithPayload adds the payload to the geographic address get o k response
func (o *GeographicAddressGetOK) WithPayload(payload *models.GeographicAddress) *GeographicAddressGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get o k response
func (o *GeographicAddressGetOK) SetPayload(payload *models.GeographicAddress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetBadRequestCode is the HTTP code returned for type GeographicAddressGetBadRequest
const GeographicAddressGetBadRequestCode int = 400

/*GeographicAddressGetBadRequest Bad Request

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

swagger:response geographicAddressGetBadRequest
*/
type GeographicAddressGetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetBadRequest creates GeographicAddressGetBadRequest with default headers values
func NewGeographicAddressGetBadRequest() *GeographicAddressGetBadRequest {

	return &GeographicAddressGetBadRequest{}
}

// WithPayload adds the payload to the geographic address get bad request response
func (o *GeographicAddressGetBadRequest) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get bad request response
func (o *GeographicAddressGetBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetUnauthorizedCode is the HTTP code returned for type GeographicAddressGetUnauthorized
const GeographicAddressGetUnauthorizedCode int = 401

/*GeographicAddressGetUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response geographicAddressGetUnauthorized
*/
type GeographicAddressGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetUnauthorized creates GeographicAddressGetUnauthorized with default headers values
func NewGeographicAddressGetUnauthorized() *GeographicAddressGetUnauthorized {

	return &GeographicAddressGetUnauthorized{}
}

// WithPayload adds the payload to the geographic address get unauthorized response
func (o *GeographicAddressGetUnauthorized) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get unauthorized response
func (o *GeographicAddressGetUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetForbiddenCode is the HTTP code returned for type GeographicAddressGetForbidden
const GeographicAddressGetForbiddenCode int = 403

/*GeographicAddressGetForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response geographicAddressGetForbidden
*/
type GeographicAddressGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetForbidden creates GeographicAddressGetForbidden with default headers values
func NewGeographicAddressGetForbidden() *GeographicAddressGetForbidden {

	return &GeographicAddressGetForbidden{}
}

// WithPayload adds the payload to the geographic address get forbidden response
func (o *GeographicAddressGetForbidden) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get forbidden response
func (o *GeographicAddressGetForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetNotFoundCode is the HTTP code returned for type GeographicAddressGetNotFound
const GeographicAddressGetNotFoundCode int = 404

/*GeographicAddressGetNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response geographicAddressGetNotFound
*/
type GeographicAddressGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetNotFound creates GeographicAddressGetNotFound with default headers values
func NewGeographicAddressGetNotFound() *GeographicAddressGetNotFound {

	return &GeographicAddressGetNotFound{}
}

// WithPayload adds the payload to the geographic address get not found response
func (o *GeographicAddressGetNotFound) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get not found response
func (o *GeographicAddressGetNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetMethodNotAllowedCode is the HTTP code returned for type GeographicAddressGetMethodNotAllowed
const GeographicAddressGetMethodNotAllowedCode int = 405

/*GeographicAddressGetMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response geographicAddressGetMethodNotAllowed
*/
type GeographicAddressGetMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetMethodNotAllowed creates GeographicAddressGetMethodNotAllowed with default headers values
func NewGeographicAddressGetMethodNotAllowed() *GeographicAddressGetMethodNotAllowed {

	return &GeographicAddressGetMethodNotAllowed{}
}

// WithPayload adds the payload to the geographic address get method not allowed response
func (o *GeographicAddressGetMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get method not allowed response
func (o *GeographicAddressGetMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetUnprocessableEntityCode is the HTTP code returned for type GeographicAddressGetUnprocessableEntity
const GeographicAddressGetUnprocessableEntityCode int = 422

/*GeographicAddressGetUnprocessableEntity Unprocessable entity

Functional error

swagger:response geographicAddressGetUnprocessableEntity
*/
type GeographicAddressGetUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetUnprocessableEntity creates GeographicAddressGetUnprocessableEntity with default headers values
func NewGeographicAddressGetUnprocessableEntity() *GeographicAddressGetUnprocessableEntity {

	return &GeographicAddressGetUnprocessableEntity{}
}

// WithPayload adds the payload to the geographic address get unprocessable entity response
func (o *GeographicAddressGetUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get unprocessable entity response
func (o *GeographicAddressGetUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetInternalServerErrorCode is the HTTP code returned for type GeographicAddressGetInternalServerError
const GeographicAddressGetInternalServerErrorCode int = 500

/*GeographicAddressGetInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response geographicAddressGetInternalServerError
*/
type GeographicAddressGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetInternalServerError creates GeographicAddressGetInternalServerError with default headers values
func NewGeographicAddressGetInternalServerError() *GeographicAddressGetInternalServerError {

	return &GeographicAddressGetInternalServerError{}
}

// WithPayload adds the payload to the geographic address get internal server error response
func (o *GeographicAddressGetInternalServerError) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get internal server error response
func (o *GeographicAddressGetInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GeographicAddressGetServiceUnavailableCode is the HTTP code returned for type GeographicAddressGetServiceUnavailable
const GeographicAddressGetServiceUnavailableCode int = 503

/*GeographicAddressGetServiceUnavailable Service Unavailable



swagger:response geographicAddressGetServiceUnavailable
*/
type GeographicAddressGetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewGeographicAddressGetServiceUnavailable creates GeographicAddressGetServiceUnavailable with default headers values
func NewGeographicAddressGetServiceUnavailable() *GeographicAddressGetServiceUnavailable {

	return &GeographicAddressGetServiceUnavailable{}
}

// WithPayload adds the payload to the geographic address get service unavailable response
func (o *GeographicAddressGetServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *GeographicAddressGetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the geographic address get service unavailable response
func (o *GeographicAddressGetServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GeographicAddressGetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
