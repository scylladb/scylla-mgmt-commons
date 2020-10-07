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

// ColumnFamilyMetricsCasPrepareByNameGetReader is a Reader for the ColumnFamilyMetricsCasPrepareByNameGet structure.
type ColumnFamilyMetricsCasPrepareByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCasPrepareByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsCasPrepareByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsCasPrepareByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsCasPrepareByNameGetOK creates a ColumnFamilyMetricsCasPrepareByNameGetOK with default headers values
func NewColumnFamilyMetricsCasPrepareByNameGetOK() *ColumnFamilyMetricsCasPrepareByNameGetOK {
	return &ColumnFamilyMetricsCasPrepareByNameGetOK{}
}

/*ColumnFamilyMetricsCasPrepareByNameGetOK handles this case with default header values.

ColumnFamilyMetricsCasPrepareByNameGetOK column family metrics cas prepare by name get o k
*/
type ColumnFamilyMetricsCasPrepareByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsCasPrepareByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsCasPrepareByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsCasPrepareByNameGetDefault creates a ColumnFamilyMetricsCasPrepareByNameGetDefault with default headers values
func NewColumnFamilyMetricsCasPrepareByNameGetDefault(code int) *ColumnFamilyMetricsCasPrepareByNameGetDefault {
	return &ColumnFamilyMetricsCasPrepareByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsCasPrepareByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsCasPrepareByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics cas prepare by name get default response
func (o *ColumnFamilyMetricsCasPrepareByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsCasPrepareByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsCasPrepareByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsCasPrepareByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
