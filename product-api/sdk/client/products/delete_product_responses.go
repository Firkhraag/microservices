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

// DeleteProductReader is a Reader for the DeleteProduct structure.
type DeleteProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductOK creates a DeleteProductOK with default headers values
func NewDeleteProductOK() *DeleteProductOK {
	return &DeleteProductOK{}
}

/*DeleteProductOK handles this case with default header values.

This API endpoint returns no content
*/
type DeleteProductOK struct {
}

func (o *DeleteProductOK) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductOK ", 200)
}

func (o *DeleteProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductNotFound creates a DeleteProductNotFound with default headers values
func NewDeleteProductNotFound() *DeleteProductNotFound {
	return &DeleteProductNotFound{}
}

/*DeleteProductNotFound handles this case with default header values.

Generic error message returned as a string
*/
type DeleteProductNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteProductNotFound) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
