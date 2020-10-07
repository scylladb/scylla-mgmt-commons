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

// StorageServiceTombstoneWarnThresholdPostReader is a Reader for the StorageServiceTombstoneWarnThresholdPost structure.
type StorageServiceTombstoneWarnThresholdPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTombstoneWarnThresholdPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTombstoneWarnThresholdPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTombstoneWarnThresholdPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTombstoneWarnThresholdPostOK creates a StorageServiceTombstoneWarnThresholdPostOK with default headers values
func NewStorageServiceTombstoneWarnThresholdPostOK() *StorageServiceTombstoneWarnThresholdPostOK {
	return &StorageServiceTombstoneWarnThresholdPostOK{}
}

/*StorageServiceTombstoneWarnThresholdPostOK handles this case with default header values.

StorageServiceTombstoneWarnThresholdPostOK storage service tombstone warn threshold post o k
*/
type StorageServiceTombstoneWarnThresholdPostOK struct {
}

func (o *StorageServiceTombstoneWarnThresholdPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceTombstoneWarnThresholdPostDefault creates a StorageServiceTombstoneWarnThresholdPostDefault with default headers values
func NewStorageServiceTombstoneWarnThresholdPostDefault(code int) *StorageServiceTombstoneWarnThresholdPostDefault {
	return &StorageServiceTombstoneWarnThresholdPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceTombstoneWarnThresholdPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceTombstoneWarnThresholdPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service tombstone warn threshold post default response
func (o *StorageServiceTombstoneWarnThresholdPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTombstoneWarnThresholdPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTombstoneWarnThresholdPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTombstoneWarnThresholdPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
