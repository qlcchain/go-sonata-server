// Code generated by go-swagger; DO NOT EDIT.

package product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
)

// ProductOrderFindOKCode is the HTTP code returned for type ProductOrderFindOK
const ProductOrderFindOKCode int = 200

/*ProductOrderFindOK Ok

swagger:response productOrderFindOK
*/
type ProductOrderFindOK struct {
	/*The number of resources retrieved in the response

	 */
	XResultCount int32 `json:"X-Result-Count"`
	/*The total number of matching resources

	 */
	XTotalCount int32 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.ProductOrderSummary `json:"body,omitempty"`
}

// NewProductOrderFindOK creates ProductOrderFindOK with default headers values
func NewProductOrderFindOK() *ProductOrderFindOK {

	return &ProductOrderFindOK{}
}

// WithXResultCount adds the xResultCount to the product order find o k response
func (o *ProductOrderFindOK) WithXResultCount(xResultCount int32) *ProductOrderFindOK {
	o.XResultCount = xResultCount
	return o
}

// SetXResultCount sets the xResultCount to the product order find o k response
func (o *ProductOrderFindOK) SetXResultCount(xResultCount int32) {
	o.XResultCount = xResultCount
}

// WithXTotalCount adds the xTotalCount to the product order find o k response
func (o *ProductOrderFindOK) WithXTotalCount(xTotalCount int32) *ProductOrderFindOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the product order find o k response
func (o *ProductOrderFindOK) SetXTotalCount(xTotalCount int32) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the product order find o k response
func (o *ProductOrderFindOK) WithPayload(payload []*models.ProductOrderSummary) *ProductOrderFindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find o k response
func (o *ProductOrderFindOK) SetPayload(payload []*models.ProductOrderSummary) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Result-Count

	xResultCount := swag.FormatInt32(o.XResultCount)
	if xResultCount != "" {
		rw.Header().Set("X-Result-Count", xResultCount)
	}

	// response header X-Total-Count

	xTotalCount := swag.FormatInt32(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ProductOrderSummary, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ProductOrderFindBadRequestCode is the HTTP code returned for type ProductOrderFindBadRequest
const ProductOrderFindBadRequestCode int = 400

/*ProductOrderFindBadRequest Bad Request

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

swagger:response productOrderFindBadRequest
*/
type ProductOrderFindBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindBadRequest creates ProductOrderFindBadRequest with default headers values
func NewProductOrderFindBadRequest() *ProductOrderFindBadRequest {

	return &ProductOrderFindBadRequest{}
}

// WithPayload adds the payload to the product order find bad request response
func (o *ProductOrderFindBadRequest) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find bad request response
func (o *ProductOrderFindBadRequest) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindUnauthorizedCode is the HTTP code returned for type ProductOrderFindUnauthorized
const ProductOrderFindUnauthorizedCode int = 401

/*ProductOrderFindUnauthorized Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials

swagger:response productOrderFindUnauthorized
*/
type ProductOrderFindUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindUnauthorized creates ProductOrderFindUnauthorized with default headers values
func NewProductOrderFindUnauthorized() *ProductOrderFindUnauthorized {

	return &ProductOrderFindUnauthorized{}
}

// WithPayload adds the payload to the product order find unauthorized response
func (o *ProductOrderFindUnauthorized) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find unauthorized response
func (o *ProductOrderFindUnauthorized) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindForbiddenCode is the HTTP code returned for type ProductOrderFindForbidden
const ProductOrderFindForbiddenCode int = 403

/*ProductOrderFindForbidden Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests

swagger:response productOrderFindForbidden
*/
type ProductOrderFindForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindForbidden creates ProductOrderFindForbidden with default headers values
func NewProductOrderFindForbidden() *ProductOrderFindForbidden {

	return &ProductOrderFindForbidden{}
}

// WithPayload adds the payload to the product order find forbidden response
func (o *ProductOrderFindForbidden) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find forbidden response
func (o *ProductOrderFindForbidden) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindNotFoundCode is the HTTP code returned for type ProductOrderFindNotFound
const ProductOrderFindNotFoundCode int = 404

/*ProductOrderFindNotFound Not Found

List of supported error codes:
- 60: Resource not found

swagger:response productOrderFindNotFound
*/
type ProductOrderFindNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindNotFound creates ProductOrderFindNotFound with default headers values
func NewProductOrderFindNotFound() *ProductOrderFindNotFound {

	return &ProductOrderFindNotFound{}
}

// WithPayload adds the payload to the product order find not found response
func (o *ProductOrderFindNotFound) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find not found response
func (o *ProductOrderFindNotFound) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindMethodNotAllowedCode is the HTTP code returned for type ProductOrderFindMethodNotAllowed
const ProductOrderFindMethodNotAllowedCode int = 405

/*ProductOrderFindMethodNotAllowed Method Not Allowed

List of supported error codes:
- 61: Method not allowed

swagger:response productOrderFindMethodNotAllowed
*/
type ProductOrderFindMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindMethodNotAllowed creates ProductOrderFindMethodNotAllowed with default headers values
func NewProductOrderFindMethodNotAllowed() *ProductOrderFindMethodNotAllowed {

	return &ProductOrderFindMethodNotAllowed{}
}

// WithPayload adds the payload to the product order find method not allowed response
func (o *ProductOrderFindMethodNotAllowed) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find method not allowed response
func (o *ProductOrderFindMethodNotAllowed) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindRequestTimeoutCode is the HTTP code returned for type ProductOrderFindRequestTimeout
const ProductOrderFindRequestTimeoutCode int = 408

/*ProductOrderFindRequestTimeout Request Time-out

List of supported error codes:
- 63: Request time-out

swagger:response productOrderFindRequestTimeout
*/
type ProductOrderFindRequestTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindRequestTimeout creates ProductOrderFindRequestTimeout with default headers values
func NewProductOrderFindRequestTimeout() *ProductOrderFindRequestTimeout {

	return &ProductOrderFindRequestTimeout{}
}

// WithPayload adds the payload to the product order find request timeout response
func (o *ProductOrderFindRequestTimeout) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindRequestTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find request timeout response
func (o *ProductOrderFindRequestTimeout) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(408)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindUnprocessableEntityCode is the HTTP code returned for type ProductOrderFindUnprocessableEntity
const ProductOrderFindUnprocessableEntityCode int = 422

/*ProductOrderFindUnprocessableEntity Unprocessable entity

Functional error





 - code: 100
message: Too many records retrieved - please restrict requested parameter value(s)
description:

swagger:response productOrderFindUnprocessableEntity
*/
type ProductOrderFindUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindUnprocessableEntity creates ProductOrderFindUnprocessableEntity with default headers values
func NewProductOrderFindUnprocessableEntity() *ProductOrderFindUnprocessableEntity {

	return &ProductOrderFindUnprocessableEntity{}
}

// WithPayload adds the payload to the product order find unprocessable entity response
func (o *ProductOrderFindUnprocessableEntity) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find unprocessable entity response
func (o *ProductOrderFindUnprocessableEntity) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindInternalServerErrorCode is the HTTP code returned for type ProductOrderFindInternalServerError
const ProductOrderFindInternalServerErrorCode int = 500

/*ProductOrderFindInternalServerError Internal Server Error

List of supported error codes:
- 1: Internal error

swagger:response productOrderFindInternalServerError
*/
type ProductOrderFindInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindInternalServerError creates ProductOrderFindInternalServerError with default headers values
func NewProductOrderFindInternalServerError() *ProductOrderFindInternalServerError {

	return &ProductOrderFindInternalServerError{}
}

// WithPayload adds the payload to the product order find internal server error response
func (o *ProductOrderFindInternalServerError) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find internal server error response
func (o *ProductOrderFindInternalServerError) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ProductOrderFindServiceUnavailableCode is the HTTP code returned for type ProductOrderFindServiceUnavailable
const ProductOrderFindServiceUnavailableCode int = 503

/*ProductOrderFindServiceUnavailable Service Unavailable



swagger:response productOrderFindServiceUnavailable
*/
type ProductOrderFindServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorRepresentation `json:"body,omitempty"`
}

// NewProductOrderFindServiceUnavailable creates ProductOrderFindServiceUnavailable with default headers values
func NewProductOrderFindServiceUnavailable() *ProductOrderFindServiceUnavailable {

	return &ProductOrderFindServiceUnavailable{}
}

// WithPayload adds the payload to the product order find service unavailable response
func (o *ProductOrderFindServiceUnavailable) WithPayload(payload *models.ErrorRepresentation) *ProductOrderFindServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product order find service unavailable response
func (o *ProductOrderFindServiceUnavailable) SetPayload(payload *models.ErrorRepresentation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductOrderFindServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
