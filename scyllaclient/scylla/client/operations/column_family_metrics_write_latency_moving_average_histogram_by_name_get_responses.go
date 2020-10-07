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

// ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGet structure.
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK creates a ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK() *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK {
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK column family metrics write latency moving average histogram by name get o k
*/
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault creates a ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault(code int) *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault {
	return &ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency moving average histogram by name get default response
func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyMovingAverageHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
