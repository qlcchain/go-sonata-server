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

// ProductOfferingQualificationCreateReader is a Reader for the ProductOfferingQualificationCreate structure.
type ProductOfferingQualificationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductOfferingQualificationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProductOfferingQualificationCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProductOfferingQualificationCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProductOfferingQualificationCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProductOfferingQualificationCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProductOfferingQualificationCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewProductOfferingQualificationCreateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewProductOfferingQualificationCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewProductOfferingQualificationCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProductOfferingQualificationCreateCreated creates a ProductOfferingQualificationCreateCreated with default headers values
func NewProductOfferingQualificationCreateCreated() *ProductOfferingQualificationCreateCreated {
	return &ProductOfferingQualificationCreateCreated{}
}

/* ProductOfferingQualificationCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ProductOfferingQualificationCreateCreated struct {
	Payload *models.ProductOfferingQualification
}

func (o *ProductOfferingQualificationCreateCreated) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateCreated  %+v", 201, o.Payload)
}
func (o *ProductOfferingQualificationCreateCreated) GetPayload() *models.ProductOfferingQualification {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOfferingQualification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateBadRequest creates a ProductOfferingQualificationCreateBadRequest with default headers values
func NewProductOfferingQualificationCreateBadRequest() *ProductOfferingQualificationCreateBadRequest {
	return &ProductOfferingQualificationCreateBadRequest{}
}

/* ProductOfferingQualificationCreateBadRequest describes a response with status code 400, with default header values.

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
type ProductOfferingQualificationCreateBadRequest struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateBadRequest  %+v", 400, o.Payload)
}
func (o *ProductOfferingQualificationCreateBadRequest) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateUnauthorized creates a ProductOfferingQualificationCreateUnauthorized with default headers values
func NewProductOfferingQualificationCreateUnauthorized() *ProductOfferingQualificationCreateUnauthorized {
	return &ProductOfferingQualificationCreateUnauthorized{}
}

/* ProductOfferingQualificationCreateUnauthorized describes a response with status code 401, with default header values.

 Unauthorized

List of supported error codes:
- 40: Missing credentials
- 41: Invalid credentials
- 42: Expired credentials
*/
type ProductOfferingQualificationCreateUnauthorized struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *ProductOfferingQualificationCreateUnauthorized) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateForbidden creates a ProductOfferingQualificationCreateForbidden with default headers values
func NewProductOfferingQualificationCreateForbidden() *ProductOfferingQualificationCreateForbidden {
	return &ProductOfferingQualificationCreateForbidden{}
}

/* ProductOfferingQualificationCreateForbidden describes a response with status code 403, with default header values.

 Forbidden

List of supported error codes:
- 50: Access denied
- 51: Forbidden requester
- 52: Forbidden user
- 53: Too many requests
*/
type ProductOfferingQualificationCreateForbidden struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateForbidden  %+v", 403, o.Payload)
}
func (o *ProductOfferingQualificationCreateForbidden) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateNotFound creates a ProductOfferingQualificationCreateNotFound with default headers values
func NewProductOfferingQualificationCreateNotFound() *ProductOfferingQualificationCreateNotFound {
	return &ProductOfferingQualificationCreateNotFound{}
}

/* ProductOfferingQualificationCreateNotFound describes a response with status code 404, with default header values.

 Not Found

List of supported error codes:
- 60: Resource not found
*/
type ProductOfferingQualificationCreateNotFound struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateNotFound  %+v", 404, o.Payload)
}
func (o *ProductOfferingQualificationCreateNotFound) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateRequestTimeout creates a ProductOfferingQualificationCreateRequestTimeout with default headers values
func NewProductOfferingQualificationCreateRequestTimeout() *ProductOfferingQualificationCreateRequestTimeout {
	return &ProductOfferingQualificationCreateRequestTimeout{}
}

/* ProductOfferingQualificationCreateRequestTimeout describes a response with status code 408, with default header values.

 Request Time-out

List of supported error codes:
- 63: Request time-out
*/
type ProductOfferingQualificationCreateRequestTimeout struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateRequestTimeout  %+v", 408, o.Payload)
}
func (o *ProductOfferingQualificationCreateRequestTimeout) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateUnprocessableEntity creates a ProductOfferingQualificationCreateUnprocessableEntity with default headers values
func NewProductOfferingQualificationCreateUnprocessableEntity() *ProductOfferingQualificationCreateUnprocessableEntity {
	return &ProductOfferingQualificationCreateUnprocessableEntity{}
}

/* ProductOfferingQualificationCreateUnprocessableEntity describes a response with status code 422, with default header values.

 Unprocessable entity

Functional error





 - code: 100
message: A relatedParty - at productOfferingQualification level - with a role 'Buyer' must be provided (including contact information)
description:


 - code: 101
message: A least a productOffering OR a productSpecification OR a Product must be provided for a POQItem
description:


 - code: 102
message: Provided Product Offering Identifier is unknown
description:


 - code: 103
message: Provided Product Specification Identifier is unknown
description:


 - code: 104
message: Provided Product Identifier is unknown
description:


 - code: 105
message: The place information provided are invalid
description:


 - code: 106
message: A least one date provided is invalid
description:


 - code: 107
message: Incorrect related party role provided
description:
*/
type ProductOfferingQualificationCreateUnprocessableEntity struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ProductOfferingQualificationCreateUnprocessableEntity) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductOfferingQualificationCreateServiceUnavailable creates a ProductOfferingQualificationCreateServiceUnavailable with default headers values
func NewProductOfferingQualificationCreateServiceUnavailable() *ProductOfferingQualificationCreateServiceUnavailable {
	return &ProductOfferingQualificationCreateServiceUnavailable{}
}

/* ProductOfferingQualificationCreateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable


*/
type ProductOfferingQualificationCreateServiceUnavailable struct {
	Payload *models.ErrorRepresentation
}

func (o *ProductOfferingQualificationCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /productOfferingQualificationManagement/v3/productOfferingQualification][%d] productOfferingQualificationCreateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ProductOfferingQualificationCreateServiceUnavailable) GetPayload() *models.ErrorRepresentation {
	return o.Payload
}

func (o *ProductOfferingQualificationCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
