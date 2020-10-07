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

// ColumnFamilyMetricsPendingFlushesByNameGetReader is a Reader for the ColumnFamilyMetricsPendingFlushesByNameGet structure.
type ColumnFamilyMetricsPendingFlushesByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsPendingFlushesByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsPendingFlushesByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsPendingFlushesByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsPendingFlushesByNameGetOK creates a ColumnFamilyMetricsPendingFlushesByNameGetOK with default headers values
func NewColumnFamilyMetricsPendingFlushesByNameGetOK() *ColumnFamilyMetricsPendingFlushesByNameGetOK {
	return &ColumnFamilyMetricsPendingFlushesByNameGetOK{}
}

/*ColumnFamilyMetricsPendingFlushesByNameGetOK handles this case with default header values.

ColumnFamilyMetricsPendingFlushesByNameGetOK column family metrics pending flushes by name get o k
*/
type ColumnFamilyMetricsPendingFlushesByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsPendingFlushesByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsPendingFlushesByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsPendingFlushesByNameGetDefault creates a ColumnFamilyMetricsPendingFlushesByNameGetDefault with default headers values
func NewColumnFamilyMetricsPendingFlushesByNameGetDefault(code int) *ColumnFamilyMetricsPendingFlushesByNameGetDefault {
	return &ColumnFamilyMetricsPendingFlushesByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsPendingFlushesByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsPendingFlushesByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics pending flushes by name get default response
func (o *ColumnFamilyMetricsPendingFlushesByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsPendingFlushesByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsPendingFlushesByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsPendingFlushesByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
