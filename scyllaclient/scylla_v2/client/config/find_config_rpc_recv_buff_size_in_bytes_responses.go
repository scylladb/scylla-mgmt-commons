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

// FindConfigRPCRecvBuffSizeInBytesReader is a Reader for the FindConfigRPCRecvBuffSizeInBytes structure.
type FindConfigRPCRecvBuffSizeInBytesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRPCRecvBuffSizeInBytesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRPCRecvBuffSizeInBytesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRPCRecvBuffSizeInBytesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRPCRecvBuffSizeInBytesOK creates a FindConfigRPCRecvBuffSizeInBytesOK with default headers values
func NewFindConfigRPCRecvBuffSizeInBytesOK() *FindConfigRPCRecvBuffSizeInBytesOK {
	return &FindConfigRPCRecvBuffSizeInBytesOK{}
}

/*FindConfigRPCRecvBuffSizeInBytesOK handles this case with default header values.

Config value
*/
type FindConfigRPCRecvBuffSizeInBytesOK struct {
	Payload int64
}

func (o *FindConfigRPCRecvBuffSizeInBytesOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigRPCRecvBuffSizeInBytesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRPCRecvBuffSizeInBytesDefault creates a FindConfigRPCRecvBuffSizeInBytesDefault with default headers values
func NewFindConfigRPCRecvBuffSizeInBytesDefault(code int) *FindConfigRPCRecvBuffSizeInBytesDefault {
	return &FindConfigRPCRecvBuffSizeInBytesDefault{
		_statusCode: code,
	}
}

/*FindConfigRPCRecvBuffSizeInBytesDefault handles this case with default header values.

unexpected error
*/
type FindConfigRPCRecvBuffSizeInBytesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config rpc recv buff size in bytes default response
func (o *FindConfigRPCRecvBuffSizeInBytesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRPCRecvBuffSizeInBytesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRPCRecvBuffSizeInBytesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRPCRecvBuffSizeInBytesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
