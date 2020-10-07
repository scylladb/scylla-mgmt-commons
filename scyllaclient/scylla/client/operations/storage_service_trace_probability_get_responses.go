// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla/models"
)

// StorageServiceTraceProbabilityGetReader is a Reader for the StorageServiceTraceProbabilityGet structure.
type StorageServiceTraceProbabilityGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTraceProbabilityGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTraceProbabilityGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTraceProbabilityGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTraceProbabilityGetOK creates a StorageServiceTraceProbabilityGetOK with default headers values
func NewStorageServiceTraceProbabilityGetOK() *StorageServiceTraceProbabilityGetOK {
	return &StorageServiceTraceProbabilityGetOK{}
}

/*StorageServiceTraceProbabilityGetOK handles this case with default header values.

StorageServiceTraceProbabilityGetOK storage service trace probability get o k
*/
type StorageServiceTraceProbabilityGetOK struct {
	Payload interface{}
}

func (o *StorageServiceTraceProbabilityGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageServiceTraceProbabilityGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceTraceProbabilityGetDefault creates a StorageServiceTraceProbabilityGetDefault with default headers values
func NewStorageServiceTraceProbabilityGetDefault(code int) *StorageServiceTraceProbabilityGetDefault {
	return &StorageServiceTraceProbabilityGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceTraceProbabilityGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceTraceProbabilityGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service trace probability get default response
func (o *StorageServiceTraceProbabilityGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTraceProbabilityGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTraceProbabilityGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTraceProbabilityGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
