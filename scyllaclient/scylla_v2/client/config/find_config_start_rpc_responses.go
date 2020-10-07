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

// FindConfigStartRPCReader is a Reader for the FindConfigStartRPC structure.
type FindConfigStartRPCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigStartRPCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigStartRPCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigStartRPCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigStartRPCOK creates a FindConfigStartRPCOK with default headers values
func NewFindConfigStartRPCOK() *FindConfigStartRPCOK {
	return &FindConfigStartRPCOK{}
}

/*FindConfigStartRPCOK handles this case with default header values.

Config value
*/
type FindConfigStartRPCOK struct {
	Payload bool
}

func (o *FindConfigStartRPCOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigStartRPCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigStartRPCDefault creates a FindConfigStartRPCDefault with default headers values
func NewFindConfigStartRPCDefault(code int) *FindConfigStartRPCDefault {
	return &FindConfigStartRPCDefault{
		_statusCode: code,
	}
}

/*FindConfigStartRPCDefault handles this case with default header values.

unexpected error
*/
type FindConfigStartRPCDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config start rpc default response
func (o *FindConfigStartRPCDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigStartRPCDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigStartRPCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigStartRPCDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
