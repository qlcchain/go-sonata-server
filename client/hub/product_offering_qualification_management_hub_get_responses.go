// Code generated by go-swagger; DO NOT EDIT.

package hub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

// ProductOfferingQualificationManagementHubGetReader is a Reader for the ProductOfferingQualificationManagementHubGet structure.
type ProductOfferingQualificationManagementHubGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductOfferingQualificationManagementHubGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductOfferingQualificationManagementHubGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProductOfferingQualificationManagementHubGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProductOfferingQualificationManagementHubGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProductOfferingQualificationManagementHubGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProductOfferingQualificationManagementHubGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewProductOfferingQualificationManagementHubGetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProductOfferingQualificationManagementHubGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewProductOfferingQualificationManagementHubGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProductOfferingQualificationManagementHubGetOK creates a ProductOfferingQualificationManagementHubGetOK with default headers values
func NewProductOfferingQualificationManagementHubGetOK() *ProductOfferingQualificationManagementHubGetOK {
	return &ProductOfferingQualificationManagementHubGetOK{}
}

/*ProductOfferingQualificationManagementHubGetOK handles this case with default header values.

Ok
*/
type ProductOfferingQualificationManagementHubGetOK struct {
	Payload []*models.Hub
}

func (o *ProductOfferingQualificationManagementHubGetOK) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetOK  %+v", 200, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetOK) GetPayload() []*models.Hub {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetBadRequest creates a ProductOfferingQualificationManagementHubGetBadRequest with default headers values
func NewProductOfferingQualificationManagementHubGetBadRequest() *ProductOfferingQualificationManagementHubGetBadRequest {
	return &ProductOfferingQualificationManagementHubGetBadRequest{}
}

/*ProductOfferingQualificationManagementHubGetBadRequest handles this case with default header values.

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
type ProductOfferingQualificationManagementHubGetBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetBadRequest  %+v", 400, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetUnauthorized creates a ProductOfferingQualificationManagementHubGetUnauthorized with default headers values
func NewProductOfferingQualificationManagementHubGetUnauthorized() *ProductOfferingQualificationManagementHubGetUnauthorized {
	return &ProductOfferingQualificationManagementHubGetUnauthorized{}
}

/*ProductOfferingQualificationManagementHubGetUnauthorized handles this case with default header values.

Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type ProductOfferingQualificationManagementHubGetUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetForbidden creates a ProductOfferingQualificationManagementHubGetForbidden with default headers values
func NewProductOfferingQualificationManagementHubGetForbidden() *ProductOfferingQualificationManagementHubGetForbidden {
	return &ProductOfferingQualificationManagementHubGetForbidden{}
}

/*ProductOfferingQualificationManagementHubGetForbidden handles this case with default header values.

Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type ProductOfferingQualificationManagementHubGetForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetForbidden) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetForbidden  %+v", 403, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetNotFound creates a ProductOfferingQualificationManagementHubGetNotFound with default headers values
func NewProductOfferingQualificationManagementHubGetNotFound() *ProductOfferingQualificationManagementHubGetNotFound {
	return &ProductOfferingQualificationManagementHubGetNotFound{}
}

/*ProductOfferingQualificationManagementHubGetNotFound handles this case with default header values.

Not Found

List of supported error codes:
- 60: Resource not found
*/
type ProductOfferingQualificationManagementHubGetNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetNotFound) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetNotFound  %+v", 404, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetRequestTimeout creates a ProductOfferingQualificationManagementHubGetRequestTimeout with default headers values
func NewProductOfferingQualificationManagementHubGetRequestTimeout() *ProductOfferingQualificationManagementHubGetRequestTimeout {
	return &ProductOfferingQualificationManagementHubGetRequestTimeout{}
}

/*ProductOfferingQualificationManagementHubGetRequestTimeout handles this case with default header values.

Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type ProductOfferingQualificationManagementHubGetRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetRequestTimeout  %+v", 408, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetUnprocessableEntity creates a ProductOfferingQualificationManagementHubGetUnprocessableEntity with default headers values
func NewProductOfferingQualificationManagementHubGetUnprocessableEntity() *ProductOfferingQualificationManagementHubGetUnprocessableEntity {
	return &ProductOfferingQualificationManagementHubGetUnprocessableEntity{}
}

/*ProductOfferingQualificationManagementHubGetUnprocessableEntity handles this case with default header values.

Unprocessable entity

Functional error
*/
type ProductOfferingQualificationManagementHubGetUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationManagementHubGetServiceUnavailable creates a ProductOfferingQualificationManagementHubGetServiceUnavailable with default headers values
func NewProductOfferingQualificationManagementHubGetServiceUnavailable() *ProductOfferingQualificationManagementHubGetServiceUnavailable {
	return &ProductOfferingQualificationManagementHubGetServiceUnavailable{}
}

/*ProductOfferingQualificationManagementHubGetServiceUnavailable handles this case with default header values.

Service Unavailable


*/
type ProductOfferingQualificationManagementHubGetServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationManagementHubGetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /productOfferingQualificationManagement/v3/hub][%d] productOfferingQualificationManagementHubGetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ProductOfferingQualificationManagementHubGetServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationManagementHubGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
