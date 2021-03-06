// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/qlcchain/go-sonata-server/models"
)

// ProductOrderManagementHubDeleteNoContentCode is the HTTP code returned for type ProductOrderManagementHubDeleteNoContent
const ProductOrderManagementHubDeleteNoContentCode int = 204

/*ProductOrderManagementHubDeleteNoContent No Content

swagger:response productOrderManagementHubDeleteNoContent
*/
type ProductOrderManagementHubDeleteNoContent struct {
}

// NewProductOrderManagementHubDeleteNoContent creates ProductOrderManagementHubDeleteNoContent with default headers values
func NewProductOrderManagementHubDeleteNoContent() *ProductOrderManagementHubDeleteNoContent {

	return &ProductOrderManagementHubDeleteNoContent{}
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ProductOrderManagementHubDeleteBadRequestCode is the HTTP code returned for type ProductOrderManagementHubDeleteBadRequest
const ProductOrderManagementHubDeleteBadRequestCode int = 400

/*ProductOrderManagementHubDeleteBadRequest Bad Request

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

swagger:response productOrderManagementHubDeleteBadRequest
*/
type ProductOrderManagementHubDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteBadRequest creates ProductOrderManagementHubDeleteBadRequest with default headers values
func NewProductOrderManagementHubDeleteBadRequest() *ProductOrderManagementHubDeleteBadRequest {

	return &ProductOrderManagementHubDeleteBadRequest{}
}

// WithPayload adds the payload to the product order management hub delete bad request response
func (o *ProductOrderManagementHubDeleteBadRequest) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete bad request response
func (o *ProductOrderManagementHubDeleteBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteUnauthorizedCode is the HTTP code returned for type ProductOrderManagementHubDeleteUnauthorized
const ProductOrderManagementHubDeleteUnauthorizedCode int = 401

/*ProductOrderManagementHubDeleteUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response productOrderManagementHubDeleteUnauthorized
*/
type ProductOrderManagementHubDeleteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteUnauthorized creates ProductOrderManagementHubDeleteUnauthorized with default headers values
func NewProductOrderManagementHubDeleteUnauthorized() *ProductOrderManagementHubDeleteUnauthorized {

	return &ProductOrderManagementHubDeleteUnauthorized{}
}

// WithPayload adds the payload to the product order management hub delete unauthorized response
func (o *ProductOrderManagementHubDeleteUnauthorized) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete unauthorized response
func (o *ProductOrderManagementHubDeleteUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteForbiddenCode is the HTTP code returned for type ProductOrderManagementHubDeleteForbidden
const ProductOrderManagementHubDeleteForbiddenCode int = 403

/*ProductOrderManagementHubDeleteForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response productOrderManagementHubDeleteForbidden
*/
type ProductOrderManagementHubDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteForbidden creates ProductOrderManagementHubDeleteForbidden with default headers values
func NewProductOrderManagementHubDeleteForbidden() *ProductOrderManagementHubDeleteForbidden {

	return &ProductOrderManagementHubDeleteForbidden{}
}

// WithPayload adds the payload to the product order management hub delete forbidden response
func (o *ProductOrderManagementHubDeleteForbidden) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete forbidden response
func (o *ProductOrderManagementHubDeleteForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteNotFoundCode is the HTTP code returned for type ProductOrderManagementHubDeleteNotFound
const ProductOrderManagementHubDeleteNotFoundCode int = 404

/*ProductOrderManagementHubDeleteNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response productOrderManagementHubDeleteNotFound
*/
type ProductOrderManagementHubDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteNotFound creates ProductOrderManagementHubDeleteNotFound with default headers values
func NewProductOrderManagementHubDeleteNotFound() *ProductOrderManagementHubDeleteNotFound {

	return &ProductOrderManagementHubDeleteNotFound{}
}

// WithPayload adds the payload to the product order management hub delete not found response
func (o *ProductOrderManagementHubDeleteNotFound) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete not found response
func (o *ProductOrderManagementHubDeleteNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteMethodNotAllowedCode is the HTTP code returned for type ProductOrderManagementHubDeleteMethodNotAllowed
const ProductOrderManagementHubDeleteMethodNotAllowedCode int = 405

/*ProductOrderManagementHubDeleteMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response productOrderManagementHubDeleteMethodNotAllowed
*/
type ProductOrderManagementHubDeleteMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteMethodNotAllowed creates ProductOrderManagementHubDeleteMethodNotAllowed with default headers values
func NewProductOrderManagementHubDeleteMethodNotAllowed() *ProductOrderManagementHubDeleteMethodNotAllowed {

	return &ProductOrderManagementHubDeleteMethodNotAllowed{}
}

// WithPayload adds the payload to the product order management hub delete method not allowed response
func (o *ProductOrderManagementHubDeleteMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete method not allowed response
func (o *ProductOrderManagementHubDeleteMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteRequestTimeoutCode is the HTTP code returned for type ProductOrderManagementHubDeleteRequestTimeout
const ProductOrderManagementHubDeleteRequestTimeoutCode int = 408

/*ProductOrderManagementHubDeleteRequestTimeout Request Time-out

List of supported error codes:
- 63: Request time-out

swagger:response productOrderManagementHubDeleteRequestTimeout
*/
type ProductOrderManagementHubDeleteRequestTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteRequestTimeout creates ProductOrderManagementHubDeleteRequestTimeout with default headers values
func NewProductOrderManagementHubDeleteRequestTimeout() *ProductOrderManagementHubDeleteRequestTimeout {

	return &ProductOrderManagementHubDeleteRequestTimeout{}
}

// WithPayload adds the payload to the product order management hub delete request timeout response
func (o *ProductOrderManagementHubDeleteRequestTimeout) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteRequestTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete request timeout response
func (o *ProductOrderManagementHubDeleteRequestTimeout) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(408)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteUnprocessableEntityCode is the HTTP code returned for type ProductOrderManagementHubDeleteUnprocessableEntity
const ProductOrderManagementHubDeleteUnprocessableEntityCode int = 422

/*ProductOrderManagementHubDeleteUnprocessableEntity Unprocessable entity

Functional error

swagger:response productOrderManagementHubDeleteUnprocessableEntity
*/
type ProductOrderManagementHubDeleteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteUnprocessableEntity creates ProductOrderManagementHubDeleteUnprocessableEntity with default headers values
func NewProductOrderManagementHubDeleteUnprocessableEntity() *ProductOrderManagementHubDeleteUnprocessableEntity {

	return &ProductOrderManagementHubDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the product order management hub delete unprocessable entity response
func (o *ProductOrderManagementHubDeleteUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete unprocessable entity response
func (o *ProductOrderManagementHubDeleteUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteInternalServerErrorCode is the HTTP code returned for type ProductOrderManagementHubDeleteInternalServerError
const ProductOrderManagementHubDeleteInternalServerErrorCode int = 500

/*ProductOrderManagementHubDeleteInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response productOrderManagementHubDeleteInternalServerError
*/
type ProductOrderManagementHubDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteInternalServerError creates ProductOrderManagementHubDeleteInternalServerError with default headers values
func NewProductOrderManagementHubDeleteInternalServerError() *ProductOrderManagementHubDeleteInternalServerError {

	return &ProductOrderManagementHubDeleteInternalServerError{}
}

// WithPayload adds the payload to the product order management hub delete internal server error response
func (o *ProductOrderManagementHubDeleteInternalServerError) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete internal server error response
func (o *ProductOrderManagementHubDeleteInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderManagementHubDeleteServiceUnavailableCode is the HTTP code returned for type ProductOrderManagementHubDeleteServiceUnavailable
const ProductOrderManagementHubDeleteServiceUnavailableCode int = 503

/*ProductOrderManagementHubDeleteServiceUnavailable Service Unavailable



swagger:response productOrderManagementHubDeleteServiceUnavailable
*/
type ProductOrderManagementHubDeleteServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderManagementHubDeleteServiceUnavailable creates ProductOrderManagementHubDeleteServiceUnavailable with default headers values
func NewProductOrderManagementHubDeleteServiceUnavailable() *ProductOrderManagementHubDeleteServiceUnavailable {

	return &ProductOrderManagementHubDeleteServiceUnavailable{}
}

// WithPayload adds the payload to the product order management hub delete service unavailable response
func (o *ProductOrderManagementHubDeleteServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *ProductOrderManagementHubDeleteServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order management hub delete service unavailable response
func (o *ProductOrderManagementHubDeleteServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderManagementHubDeleteServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
