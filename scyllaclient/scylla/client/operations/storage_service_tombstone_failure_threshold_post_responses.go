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

// StorageServiceTombstoneFailureThresholdPostReader is a Reader for the StorageServiceTombstoneFailureThresholdPost structure.
type StorageServiceTombstoneFailureThresholdPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTombstoneFailureThresholdPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTombstoneFailureThresholdPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTombstoneFailureThresholdPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTombstoneFailureThresholdPostOK creates a StorageServiceTombstoneFailureThresholdPostOK with default headers values
func NewStorageServiceTombstoneFailureThresholdPostOK() *StorageServiceTombstoneFailureThresholdPostOK {
	return &StorageServiceTombstoneFailureThresholdPostOK{}
}

/*StorageServiceTombstoneFailureThresholdPostOK handles this case with default header values.

StorageServiceTombstoneFailureThresholdPostOK storage service tombstone failure threshold post o k
*/
type StorageServiceTombstoneFailureThresholdPostOK struct {
}

func (o *StorageServiceTombstoneFailureThresholdPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceTombstoneFailureThresholdPostDefault creates a StorageServiceTombstoneFailureThresholdPostDefault with default headers values
func NewStorageServiceTombstoneFailureThresholdPostDefault(code int) *StorageServiceTombstoneFailureThresholdPostDefault {
	return &StorageServiceTombstoneFailureThresholdPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceTombstoneFailureThresholdPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceTombstoneFailureThresholdPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service tombstone failure threshold post default response
func (o *StorageServiceTombstoneFailureThresholdPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTombstoneFailureThresholdPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTombstoneFailureThresholdPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTombstoneFailureThresholdPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
