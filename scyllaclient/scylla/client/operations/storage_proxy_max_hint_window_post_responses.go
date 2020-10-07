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

// StorageProxyMaxHintWindowPostReader is a Reader for the StorageProxyMaxHintWindowPost structure.
type StorageProxyMaxHintWindowPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMaxHintWindowPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMaxHintWindowPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMaxHintWindowPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMaxHintWindowPostOK creates a StorageProxyMaxHintWindowPostOK with default headers values
func NewStorageProxyMaxHintWindowPostOK() *StorageProxyMaxHintWindowPostOK {
	return &StorageProxyMaxHintWindowPostOK{}
}

/*StorageProxyMaxHintWindowPostOK handles this case with default header values.

StorageProxyMaxHintWindowPostOK storage proxy max hint window post o k
*/
type StorageProxyMaxHintWindowPostOK struct {
}

func (o *StorageProxyMaxHintWindowPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyMaxHintWindowPostDefault creates a StorageProxyMaxHintWindowPostDefault with default headers values
func NewStorageProxyMaxHintWindowPostDefault(code int) *StorageProxyMaxHintWindowPostDefault {
	return &StorageProxyMaxHintWindowPostDefault{
		_statusCode: code,
	}
}

/*StorageProxyMaxHintWindowPostDefault handles this case with default header values.

internal server error
*/
type StorageProxyMaxHintWindowPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy max hint window post default response
func (o *StorageProxyMaxHintWindowPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMaxHintWindowPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMaxHintWindowPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMaxHintWindowPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
