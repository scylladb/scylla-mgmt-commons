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

// FindConfigRequestSchedulerReader is a Reader for the FindConfigRequestScheduler structure.
type FindConfigRequestSchedulerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRequestSchedulerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRequestSchedulerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRequestSchedulerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRequestSchedulerOK creates a FindConfigRequestSchedulerOK with default headers values
func NewFindConfigRequestSchedulerOK() *FindConfigRequestSchedulerOK {
	return &FindConfigRequestSchedulerOK{}
}

/*FindConfigRequestSchedulerOK handles this case with default header values.

Config value
*/
type FindConfigRequestSchedulerOK struct {
	Payload string
}

func (o *FindConfigRequestSchedulerOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigRequestSchedulerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRequestSchedulerDefault creates a FindConfigRequestSchedulerDefault with default headers values
func NewFindConfigRequestSchedulerDefault(code int) *FindConfigRequestSchedulerDefault {
	return &FindConfigRequestSchedulerDefault{
		_statusCode: code,
	}
}

/*FindConfigRequestSchedulerDefault handles this case with default header values.

unexpected error
*/
type FindConfigRequestSchedulerDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config request scheduler default response
func (o *FindConfigRequestSchedulerDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRequestSchedulerDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRequestSchedulerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRequestSchedulerDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
