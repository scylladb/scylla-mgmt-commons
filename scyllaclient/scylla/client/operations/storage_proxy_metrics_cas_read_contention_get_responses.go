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

// StorageProxyMetricsCasReadContentionGetReader is a Reader for the StorageProxyMetricsCasReadContentionGet structure.
type StorageProxyMetricsCasReadContentionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasReadContentionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsCasReadContentionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsCasReadContentionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsCasReadContentionGetOK creates a StorageProxyMetricsCasReadContentionGetOK with default headers values
func NewStorageProxyMetricsCasReadContentionGetOK() *StorageProxyMetricsCasReadContentionGetOK {
	return &StorageProxyMetricsCasReadContentionGetOK{}
}

/*StorageProxyMetricsCasReadContentionGetOK handles this case with default header values.

StorageProxyMetricsCasReadContentionGetOK storage proxy metrics cas read contention get o k
*/
type StorageProxyMetricsCasReadContentionGetOK struct {
	Payload interface{}
}

func (o *StorageProxyMetricsCasReadContentionGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadContentionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsCasReadContentionGetDefault creates a StorageProxyMetricsCasReadContentionGetDefault with default headers values
func NewStorageProxyMetricsCasReadContentionGetDefault(code int) *StorageProxyMetricsCasReadContentionGetDefault {
	return &StorageProxyMetricsCasReadContentionGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsCasReadContentionGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsCasReadContentionGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics cas read contention get default response
func (o *StorageProxyMetricsCasReadContentionGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsCasReadContentionGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsCasReadContentionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsCasReadContentionGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
