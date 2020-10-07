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

// FindConfigListenInterfaceReader is a Reader for the FindConfigListenInterface structure.
type FindConfigListenInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigListenInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigListenInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigListenInterfaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigListenInterfaceOK creates a FindConfigListenInterfaceOK with default headers values
func NewFindConfigListenInterfaceOK() *FindConfigListenInterfaceOK {
	return &FindConfigListenInterfaceOK{}
}

/*FindConfigListenInterfaceOK handles this case with default header values.

Config value
*/
type FindConfigListenInterfaceOK struct {
	Payload string
}

func (o *FindConfigListenInterfaceOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigListenInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigListenInterfaceDefault creates a FindConfigListenInterfaceDefault with default headers values
func NewFindConfigListenInterfaceDefault(code int) *FindConfigListenInterfaceDefault {
	return &FindConfigListenInterfaceDefault{
		_statusCode: code,
	}
}

/*FindConfigListenInterfaceDefault handles this case with default header values.

unexpected error
*/
type FindConfigListenInterfaceDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config listen interface default response
func (o *FindConfigListenInterfaceDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigListenInterfaceDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigListenInterfaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigListenInterfaceDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
