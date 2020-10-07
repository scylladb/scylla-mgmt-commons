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

// StorageProxyTruncateRPCTimeoutGetReader is a Reader for the StorageProxyTruncateRPCTimeoutGet structure.
type StorageProxyTruncateRPCTimeoutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyTruncateRPCTimeoutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyTruncateRPCTimeoutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyTruncateRPCTimeoutGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyTruncateRPCTimeoutGetOK creates a StorageProxyTruncateRPCTimeoutGetOK with default headers values
func NewStorageProxyTruncateRPCTimeoutGetOK() *StorageProxyTruncateRPCTimeoutGetOK {
	return &StorageProxyTruncateRPCTimeoutGetOK{}
}

/*StorageProxyTruncateRPCTimeoutGetOK handles this case with default header values.

StorageProxyTruncateRPCTimeoutGetOK storage proxy truncate Rpc timeout get o k
*/
type StorageProxyTruncateRPCTimeoutGetOK struct {
	Payload interface{}
}

func (o *StorageProxyTruncateRPCTimeoutGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyTruncateRPCTimeoutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyTruncateRPCTimeoutGetDefault creates a StorageProxyTruncateRPCTimeoutGetDefault with default headers values
func NewStorageProxyTruncateRPCTimeoutGetDefault(code int) *StorageProxyTruncateRPCTimeoutGetDefault {
	return &StorageProxyTruncateRPCTimeoutGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyTruncateRPCTimeoutGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyTruncateRPCTimeoutGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy truncate Rpc timeout get default response
func (o *StorageProxyTruncateRPCTimeoutGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyTruncateRPCTimeoutGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyTruncateRPCTimeoutGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyTruncateRPCTimeoutGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
