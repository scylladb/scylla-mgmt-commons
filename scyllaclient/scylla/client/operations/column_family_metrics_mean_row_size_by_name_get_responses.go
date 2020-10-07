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

// ColumnFamilyMetricsMeanRowSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMeanRowSizeByNameGet structure.
type ColumnFamilyMetricsMeanRowSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMeanRowSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMeanRowSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMeanRowSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMeanRowSizeByNameGetOK creates a ColumnFamilyMetricsMeanRowSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMeanRowSizeByNameGetOK() *ColumnFamilyMetricsMeanRowSizeByNameGetOK {
	return &ColumnFamilyMetricsMeanRowSizeByNameGetOK{}
}

/*ColumnFamilyMetricsMeanRowSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMeanRowSizeByNameGetOK column family metrics mean row size by name get o k
*/
type ColumnFamilyMetricsMeanRowSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMeanRowSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMeanRowSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMeanRowSizeByNameGetDefault creates a ColumnFamilyMetricsMeanRowSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsMeanRowSizeByNameGetDefault(code int) *ColumnFamilyMetricsMeanRowSizeByNameGetDefault {
	return &ColumnFamilyMetricsMeanRowSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMeanRowSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMeanRowSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics mean row size by name get default response
func (o *ColumnFamilyMetricsMeanRowSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMeanRowSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMeanRowSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMeanRowSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
