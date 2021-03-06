// Code generated by go-swagger; DO NOT EDIT.

package product_offering_qualification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// ProductOfferingQualificationGetReader is a Reader for the ProductOfferingQualificationGet structure.
type ProductOfferingQualificationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductOfferingQualificationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductOfferingQualificationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProductOfferingQualificationGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProductOfferingQualificationGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProductOfferingQualificationGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProductOfferingQualificationGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewProductOfferingQualificationGetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProductOfferingQualificationGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewProductOfferingQualificationGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProductOfferingQualificationGetOK creates a ProductOfferingQualificationGetOK with default headers values
func NewProductOfferingQualificationGetOK() *ProductOfferingQualificationGetOK {
	return &ProductOfferingQualificationGetOK{}
}

/* ProductOfferingQualificationGetOK describes a response with status code 200, with default header values.

Ok
*/
type ProductOfferingQualificationGetOK struct {
	Payload *models.ProductOfferingQualification
}

func (o *ProductOfferingQualificationGetOK) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetOK  %+v", 200, o.Payload)
}
func (o *ProductOfferingQualificationGetOK) GetPayload() *models.ProductOfferingQualification {
	return o.Payload
}

func (o *ProductOfferingQualificationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOfferingQualification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetBadRequest creates a ProductOfferingQualificationGetBadRequest with default headers values
func NewProductOfferingQualificationGetBadRequest() *ProductOfferingQualificationGetBadRequest {
	return &ProductOfferingQualificationGetBadRequest{}
}

/* ProductOfferingQualificationGetBadRequest describes a response with status code 400, with default header values.

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
type ProductOfferingQualificationGetBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetBadRequest  %+v", 400, o.Payload)
}
func (o *ProductOfferingQualificationGetBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetUnauthorized creates a ProductOfferingQualificationGetUnauthorized with default headers values
func NewProductOfferingQualificationGetUnauthorized() *ProductOfferingQualificationGetUnauthorized {
	return &ProductOfferingQualificationGetUnauthorized{}
}

/* ProductOfferingQualificationGetUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type ProductOfferingQualificationGetUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetUnauthorized  %+v", 401, o.Payload)
}
func (o *ProductOfferingQualificationGetUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetForbidden creates a ProductOfferingQualificationGetForbidden with default headers values
func NewProductOfferingQualificationGetForbidden() *ProductOfferingQualificationGetForbidden {
	return &ProductOfferingQualificationGetForbidden{}
}

/* ProductOfferingQualificationGetForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type ProductOfferingQualificationGetForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetForbidden) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetForbidden  %+v", 403, o.Payload)
}
func (o *ProductOfferingQualificationGetForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetNotFound creates a ProductOfferingQualificationGetNotFound with default headers values
func NewProductOfferingQualificationGetNotFound() *ProductOfferingQualificationGetNotFound {
	return &ProductOfferingQualificationGetNotFound{}
}

/* ProductOfferingQualificationGetNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type ProductOfferingQualificationGetNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetNotFound) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetNotFound  %+v", 404, o.Payload)
}
func (o *ProductOfferingQualificationGetNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetRequestTimeout creates a ProductOfferingQualificationGetRequestTimeout with default headers values
func NewProductOfferingQualificationGetRequestTimeout() *ProductOfferingQualificationGetRequestTimeout {
	return &ProductOfferingQualificationGetRequestTimeout{}
}

/* ProductOfferingQualificationGetRequestTimeout describes a response with status code 408, with default header values.

 Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type ProductOfferingQualificationGetRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetRequestTimeout  %+v", 408, o.Payload)
}
func (o *ProductOfferingQualificationGetRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetUnprocessableEntity creates a ProductOfferingQualificationGetUnprocessableEntity with default headers values
func NewProductOfferingQualificationGetUnprocessableEntity() *ProductOfferingQualificationGetUnprocessableEntity {
	return &ProductOfferingQualificationGetUnprocessableEntity{}
}

/* ProductOfferingQualificationGetUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error
*/
type ProductOfferingQualificationGetUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ProductOfferingQualificationGetUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationGetServiceUnavailable creates a ProductOfferingQualificationGetServiceUnavailable with default headers values
func NewProductOfferingQualificationGetServiceUnavailable() *ProductOfferingQualificationGetServiceUnavailable {
	return &ProductOfferingQualificationGetServiceUnavailable{}
}

/* ProductOfferingQualificationGetServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type ProductOfferingQualificationGetServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationGetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/productOfferingQualification/{ProductOfferingQualificationId}][%d] productOfferingQualificationGetServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ProductOfferingQualificationGetServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
