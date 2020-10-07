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

// StorageProxyMetricsWriteMovingAverageHistogramGetReader is a Reader for the StorageProxyMetricsWriteMovingAverageHistogramGet structure.
type StorageProxyMetricsWriteMovingAverageHistogramGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsWriteMovingAverageHistogramGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsWriteMovingAverageHistogramGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsWriteMovingAverageHistogramGetOK creates a StorageProxyMetricsWriteMovingAverageHistogramGetOK with default headers values
func NewStorageProxyMetricsWriteMovingAverageHistogramGetOK() *StorageProxyMetricsWriteMovingAverageHistogramGetOK {
	return &StorageProxyMetricsWriteMovingAverageHistogramGetOK{}
}

/*StorageProxyMetricsWriteMovingAverageHistogramGetOK handles this case with default header values.

StorageProxyMetricsWriteMovingAverageHistogramGetOK storage proxy metrics write moving average histogram get o k
*/
type StorageProxyMetricsWriteMovingAverageHistogramGetOK struct {
}

func (o *StorageProxyMetricsWriteMovingAverageHistogramGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageProxyMetricsWriteMovingAverageHistogramGetDefault creates a StorageProxyMetricsWriteMovingAverageHistogramGetDefault with default headers values
func NewStorageProxyMetricsWriteMovingAverageHistogramGetDefault(code int) *StorageProxyMetricsWriteMovingAverageHistogramGetDefault {
	return &StorageProxyMetricsWriteMovingAverageHistogramGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsWriteMovingAverageHistogramGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsWriteMovingAverageHistogramGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics write moving average histogram get default response
func (o *StorageProxyMetricsWriteMovingAverageHistogramGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsWriteMovingAverageHistogramGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsWriteMovingAverageHistogramGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsWriteMovingAverageHistogramGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
