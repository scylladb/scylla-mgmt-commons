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

// StorageServiceTruncateByKeyspacePostReader is a Reader for the StorageServiceTruncateByKeyspacePost structure.
type StorageServiceTruncateByKeyspacePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTruncateByKeyspacePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceTruncateByKeyspacePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceTruncateByKeyspacePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceTruncateByKeyspacePostOK creates a StorageServiceTruncateByKeyspacePostOK with default headers values
func NewStorageServiceTruncateByKeyspacePostOK() *StorageServiceTruncateByKeyspacePostOK {
	return &StorageServiceTruncateByKeyspacePostOK{}
}

/*StorageServiceTruncateByKeyspacePostOK handles this case with default header values.

StorageServiceTruncateByKeyspacePostOK storage service truncate by keyspace post o k
*/
type StorageServiceTruncateByKeyspacePostOK struct {
}

func (o *StorageServiceTruncateByKeyspacePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceTruncateByKeyspacePostDefault creates a StorageServiceTruncateByKeyspacePostDefault with default headers values
func NewStorageServiceTruncateByKeyspacePostDefault(code int) *StorageServiceTruncateByKeyspacePostDefault {
	return &StorageServiceTruncateByKeyspacePostDefault{
		_statusCode: code,
	}
}

/*StorageServiceTruncateByKeyspacePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceTruncateByKeyspacePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service truncate by keyspace post default response
func (o *StorageServiceTruncateByKeyspacePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceTruncateByKeyspacePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceTruncateByKeyspacePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceTruncateByKeyspacePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
