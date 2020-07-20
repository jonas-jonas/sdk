// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra-client-go/models"
)

// DeleteJSONWebKeyReader is a Reader for the DeleteJSONWebKey structure.
type DeleteJSONWebKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJSONWebKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteJSONWebKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteJSONWebKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteJSONWebKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJSONWebKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteJSONWebKeyNoContent creates a DeleteJSONWebKeyNoContent with default headers values
func NewDeleteJSONWebKeyNoContent() *DeleteJSONWebKeyNoContent {
	return &DeleteJSONWebKeyNoContent{}
}

/*DeleteJSONWebKeyNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteJSONWebKeyNoContent struct {
}

func (o *DeleteJSONWebKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}/{kid}][%d] deleteJsonWebKeyNoContent ", 204)
}

func (o *DeleteJSONWebKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJSONWebKeyUnauthorized creates a DeleteJSONWebKeyUnauthorized with default headers values
func NewDeleteJSONWebKeyUnauthorized() *DeleteJSONWebKeyUnauthorized {
	return &DeleteJSONWebKeyUnauthorized{}
}

/*DeleteJSONWebKeyUnauthorized handles this case with default header values.

genericError
*/
type DeleteJSONWebKeyUnauthorized struct {
	Payload *models.GenericError
}

func (o *DeleteJSONWebKeyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}/{kid}][%d] deleteJsonWebKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteJSONWebKeyUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteJSONWebKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJSONWebKeyForbidden creates a DeleteJSONWebKeyForbidden with default headers values
func NewDeleteJSONWebKeyForbidden() *DeleteJSONWebKeyForbidden {
	return &DeleteJSONWebKeyForbidden{}
}

/*DeleteJSONWebKeyForbidden handles this case with default header values.

genericError
*/
type DeleteJSONWebKeyForbidden struct {
	Payload *models.GenericError
}

func (o *DeleteJSONWebKeyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}/{kid}][%d] deleteJsonWebKeyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteJSONWebKeyForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteJSONWebKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJSONWebKeyInternalServerError creates a DeleteJSONWebKeyInternalServerError with default headers values
func NewDeleteJSONWebKeyInternalServerError() *DeleteJSONWebKeyInternalServerError {
	return &DeleteJSONWebKeyInternalServerError{}
}

/*DeleteJSONWebKeyInternalServerError handles this case with default header values.

genericError
*/
type DeleteJSONWebKeyInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteJSONWebKeyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}/{kid}][%d] deleteJsonWebKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteJSONWebKeyInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteJSONWebKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
