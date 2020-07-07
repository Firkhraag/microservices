// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firkhraag/product-api/sdk/models"
)

// GetProductReader is a Reader for the GetProduct structure.
type GetProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductOK creates a GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {
	return &GetProductOK{}
}

/*GetProductOK handles this case with default header values.

This API endpoint returns a product
*/
type GetProductOK struct {
	Payload *models.Product
}

func (o *GetProductOK) Error() string {
	return fmt.Sprintf("[GET /{id}][%d] getProductOK  %+v", 200, o.Payload)
}

func (o *GetProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *GetProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductNotFound creates a GetProductNotFound with default headers values
func NewGetProductNotFound() *GetProductNotFound {
	return &GetProductNotFound{}
}

/*GetProductNotFound handles this case with default header values.

Generic error message returned as a string
*/
type GetProductNotFound struct {
	Payload *models.GenericError
}

func (o *GetProductNotFound) Error() string {
	return fmt.Sprintf("[GET /{id}][%d] getProductNotFound  %+v", 404, o.Payload)
}

func (o *GetProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductInternalServerError creates a GetProductInternalServerError with default headers values
func NewGetProductInternalServerError() *GetProductInternalServerError {
	return &GetProductInternalServerError{}
}

/*GetProductInternalServerError handles this case with default header values.

Generic error message returned as a string
*/
type GetProductInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetProductInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{id}][%d] getProductInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProductInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
