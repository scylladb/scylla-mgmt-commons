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

// StorageProxyMaxHintsInProgressPostReader is a Reader for the StorageProxyMaxHintsInProgressPost structure.
type StorageProxyMaxHintsInProgressPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMaxHintsInProgressPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMaxHintsInProgressPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMaxHintsInProgressPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMaxHintsInProgressPostOK creates a StorageProxyMaxHintsInProgressPostOK with default headers values
func NewStorageProxyMaxHintsInProgressPostOK() *StorageProxyMaxHintsInProgressPostOK {
	return &StorageProxyMaxHintsInProgressPostOK{}
}

/*StorageProxyMaxHintsInProgressPostOK handles this case with default header values.

StorageProxyMaxHintsInProgressPostOK storage proxy max hints in progress post o k
*/
type StorageProxyMaxHintsInProgressPostOK struct {
}

func (o *StorageProxyMaxHintsInProgressPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyMaxHintsInProgressPostDefault creates a StorageProxyMaxHintsInProgressPostDefault with default headers values
func NewStorageProxyMaxHintsInProgressPostDefault(code int) *StorageProxyMaxHintsInProgressPostDefault {
	return &StorageProxyMaxHintsInProgressPostDefault{
		_statusCode: code,
	}
}

/*StorageProxyMaxHintsInProgressPostDefault handles this case with default header values.

internal server error
*/
type StorageProxyMaxHintsInProgressPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy max hints in progress post default response
func (o *StorageProxyMaxHintsInProgressPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMaxHintsInProgressPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMaxHintsInProgressPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMaxHintsInProgressPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
