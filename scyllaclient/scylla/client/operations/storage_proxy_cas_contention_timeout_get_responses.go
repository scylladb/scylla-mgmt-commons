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

// StorageProxyCasContentionTimeoutGetReader is a Reader for the StorageProxyCasContentionTimeoutGet structure.
type StorageProxyCasContentionTimeoutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyCasContentionTimeoutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyCasContentionTimeoutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyCasContentionTimeoutGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyCasContentionTimeoutGetOK creates a StorageProxyCasContentionTimeoutGetOK with default headers values
func NewStorageProxyCasContentionTimeoutGetOK() *StorageProxyCasContentionTimeoutGetOK {
	return &StorageProxyCasContentionTimeoutGetOK{}
}

/*StorageProxyCasContentionTimeoutGetOK handles this case with default header values.

StorageProxyCasContentionTimeoutGetOK storage proxy cas contention timeout get o k
*/
type StorageProxyCasContentionTimeoutGetOK struct {
	Payload interface{}
}

func (o *StorageProxyCasContentionTimeoutGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyCasContentionTimeoutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyCasContentionTimeoutGetDefault creates a StorageProxyCasContentionTimeoutGetDefault with default headers values
func NewStorageProxyCasContentionTimeoutGetDefault(code int) *StorageProxyCasContentionTimeoutGetDefault {
	return &StorageProxyCasContentionTimeoutGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyCasContentionTimeoutGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyCasContentionTimeoutGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy cas contention timeout get default response
func (o *StorageProxyCasContentionTimeoutGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyCasContentionTimeoutGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyCasContentionTimeoutGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyCasContentionTimeoutGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
