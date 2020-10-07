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

// StorageServiceUpdateSnitchPostReader is a Reader for the StorageServiceUpdateSnitchPost structure.
type StorageServiceUpdateSnitchPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceUpdateSnitchPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceUpdateSnitchPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceUpdateSnitchPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceUpdateSnitchPostOK creates a StorageServiceUpdateSnitchPostOK with default headers values
func NewStorageServiceUpdateSnitchPostOK() *StorageServiceUpdateSnitchPostOK {
	return &StorageServiceUpdateSnitchPostOK{}
}

/*StorageServiceUpdateSnitchPostOK handles this case with default header values.

StorageServiceUpdateSnitchPostOK storage service update snitch post o k
*/
type StorageServiceUpdateSnitchPostOK struct {
}

func (o *StorageServiceUpdateSnitchPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceUpdateSnitchPostDefault creates a StorageServiceUpdateSnitchPostDefault with default headers values
func NewStorageServiceUpdateSnitchPostDefault(code int) *StorageServiceUpdateSnitchPostDefault {
	return &StorageServiceUpdateSnitchPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceUpdateSnitchPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceUpdateSnitchPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service update snitch post default response
func (o *StorageServiceUpdateSnitchPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceUpdateSnitchPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceUpdateSnitchPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceUpdateSnitchPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
