// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/scyllaclient/scylla_v2/models"
)

// FindConfigFdMaxIntervalMsReader is a Reader for the FindConfigFdMaxIntervalMs structure.
type FindConfigFdMaxIntervalMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigFdMaxIntervalMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigFdMaxIntervalMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigFdMaxIntervalMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigFdMaxIntervalMsOK creates a FindConfigFdMaxIntervalMsOK with default headers values
func NewFindConfigFdMaxIntervalMsOK() *FindConfigFdMaxIntervalMsOK {
	return &FindConfigFdMaxIntervalMsOK{}
}

/*FindConfigFdMaxIntervalMsOK handles this case with default header values.

Config value
*/
type FindConfigFdMaxIntervalMsOK struct {
	Payload int64
}

func (o *FindConfigFdMaxIntervalMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigFdMaxIntervalMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigFdMaxIntervalMsDefault creates a FindConfigFdMaxIntervalMsDefault with default headers values
func NewFindConfigFdMaxIntervalMsDefault(code int) *FindConfigFdMaxIntervalMsDefault {
	return &FindConfigFdMaxIntervalMsDefault{
		_statusCode: code,
	}
}

/*FindConfigFdMaxIntervalMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigFdMaxIntervalMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config fd max interval ms default response
func (o *FindConfigFdMaxIntervalMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigFdMaxIntervalMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigFdMaxIntervalMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigFdMaxIntervalMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
