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

// StorageServiceDataFileLocationsGetReader is a Reader for the StorageServiceDataFileLocationsGet structure.
type StorageServiceDataFileLocationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceDataFileLocationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceDataFileLocationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceDataFileLocationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceDataFileLocationsGetOK creates a StorageServiceDataFileLocationsGetOK with default headers values
func NewStorageServiceDataFileLocationsGetOK() *StorageServiceDataFileLocationsGetOK {
	return &StorageServiceDataFileLocationsGetOK{}
}

/*StorageServiceDataFileLocationsGetOK handles this case with default header values.

StorageServiceDataFileLocationsGetOK storage service data file locations get o k
*/
type StorageServiceDataFileLocationsGetOK struct {
	Payload []string
}

func (o *StorageServiceDataFileLocationsGetOK) GetPayload() []string {
	return o.Payload
}

func (o *StorageServiceDataFileLocationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceDataFileLocationsGetDefault creates a StorageServiceDataFileLocationsGetDefault with default headers values
func NewStorageServiceDataFileLocationsGetDefault(code int) *StorageServiceDataFileLocationsGetDefault {
	return &StorageServiceDataFileLocationsGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceDataFileLocationsGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceDataFileLocationsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service data file locations get default response
func (o *StorageServiceDataFileLocationsGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceDataFileLocationsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceDataFileLocationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceDataFileLocationsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
