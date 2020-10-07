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

// StorageProxyTruncateRPCTimeoutPostReader is a Reader for the StorageProxyTruncateRPCTimeoutPost structure.
type StorageProxyTruncateRPCTimeoutPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyTruncateRPCTimeoutPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyTruncateRPCTimeoutPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyTruncateRPCTimeoutPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyTruncateRPCTimeoutPostOK creates a StorageProxyTruncateRPCTimeoutPostOK with default headers values
func NewStorageProxyTruncateRPCTimeoutPostOK() *StorageProxyTruncateRPCTimeoutPostOK {
	return &StorageProxyTruncateRPCTimeoutPostOK{}
}

/*StorageProxyTruncateRPCTimeoutPostOK handles this case with default header values.

StorageProxyTruncateRPCTimeoutPostOK storage proxy truncate Rpc timeout post o k
*/
type StorageProxyTruncateRPCTimeoutPostOK struct {
}

func (o *StorageProxyTruncateRPCTimeoutPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyTruncateRPCTimeoutPostDefault creates a StorageProxyTruncateRPCTimeoutPostDefault with default headers values
func NewStorageProxyTruncateRPCTimeoutPostDefault(code int) *StorageProxyTruncateRPCTimeoutPostDefault {
	return &StorageProxyTruncateRPCTimeoutPostDefault{
		_statusCode: code,
	}
}

/*StorageProxyTruncateRPCTimeoutPostDefault handles this case with default header values.

internal server error
*/
type StorageProxyTruncateRPCTimeoutPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy truncate Rpc timeout post default response
func (o *StorageProxyTruncateRPCTimeoutPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyTruncateRPCTimeoutPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyTruncateRPCTimeoutPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyTruncateRPCTimeoutPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
