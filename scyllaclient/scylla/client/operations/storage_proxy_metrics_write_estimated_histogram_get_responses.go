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

// StorageProxyMetricsWriteEstimatedHistogramGetReader is a Reader for the StorageProxyMetricsWriteEstimatedHistogramGet structure.
type StorageProxyMetricsWriteEstimatedHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsWriteEstimatedHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsWriteEstimatedHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsWriteEstimatedHistogramGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsWriteEstimatedHistogramGetOK creates a StorageProxyMetricsWriteEstimatedHistogramGetOK with default headers values
func NewStorageProxyMetricsWriteEstimatedHistogramGetOK() *StorageProxyMetricsWriteEstimatedHistogramGetOK {
	return &StorageProxyMetricsWriteEstimatedHistogramGetOK{}
}

/*StorageProxyMetricsWriteEstimatedHistogramGetOK handles this case with default header values.

StorageProxyMetricsWriteEstimatedHistogramGetOK storage proxy metrics write estimated histogram get o k
*/
type StorageProxyMetricsWriteEstimatedHistogramGetOK struct {
}

func (o *StorageProxyMetricsWriteEstimatedHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyMetricsWriteEstimatedHistogramGetDefault creates a StorageProxyMetricsWriteEstimatedHistogramGetDefault with default headers values
func NewStorageProxyMetricsWriteEstimatedHistogramGetDefault(code int) *StorageProxyMetricsWriteEstimatedHistogramGetDefault {
	return &StorageProxyMetricsWriteEstimatedHistogramGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsWriteEstimatedHistogramGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsWriteEstimatedHistogramGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics write estimated histogram get default response
func (o *StorageProxyMetricsWriteEstimatedHistogramGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsWriteEstimatedHistogramGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsWriteEstimatedHistogramGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsWriteEstimatedHistogramGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
