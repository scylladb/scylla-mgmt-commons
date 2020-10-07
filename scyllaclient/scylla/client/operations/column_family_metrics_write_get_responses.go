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

// ColumnFamilyMetricsWriteGetReader is a Reader for the ColumnFamilyMetricsWriteGet structure.
type ColumnFamilyMetricsWriteGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteGetOK creates a ColumnFamilyMetricsWriteGetOK with default headers values
func NewColumnFamilyMetricsWriteGetOK() *ColumnFamilyMetricsWriteGetOK {
	return &ColumnFamilyMetricsWriteGetOK{}
}

/*ColumnFamilyMetricsWriteGetOK handles this case with default header values.

ColumnFamilyMetricsWriteGetOK column family metrics write get o k
*/
type ColumnFamilyMetricsWriteGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsWriteGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsWriteGetDefault creates a ColumnFamilyMetricsWriteGetDefault with default headers values
func NewColumnFamilyMetricsWriteGetDefault(code int) *ColumnFamilyMetricsWriteGetDefault {
	return &ColumnFamilyMetricsWriteGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write get default response
func (o *ColumnFamilyMetricsWriteGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
